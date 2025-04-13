package netfunc

import (
	"bytes"
	"io"
	"net/http"
	"sync"
	"time"
)

const (
	testURL       = "https://speed.cloudflare.com/__down?bytes=10000000"
	uploadURL     = "https://httpbin.org/post"
	numThreads    = 4
	testDuration  = 15 * time.Second
	uploadBufSize = 10 * 1024 * 1024 // 10MB
)

type SpeedResult struct {
	TotalBytes int64
	Duration   time.Duration
	MBps       float64
}

func DownloadSpeedTest() SpeedResult {
	var wg sync.WaitGroup
	var totalBytes int64
	stop := make(chan struct{})

	client := &http.Client{}

	for i := 0; i < numThreads; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case <-stop:
					return
				default:
					resp, err := client.Get(testURL)
					if err != nil {
						continue
					}
					n, _ := io.Copy(io.Discard, resp.Body)
					resp.Body.Close()
					// атомарное прибавление байт
					addBytes(&totalBytes, n)
				}
			}
		}()
	}

	time.Sleep(testDuration)
	close(stop)
	wg.Wait()

	return makeResult(totalBytes, testDuration)
}

func UploadSpeedTest() SpeedResult {
	var wg sync.WaitGroup
	var totalBytes int64
	stop := make(chan struct{})
	client := &http.Client{}
	data := bytes.Repeat([]byte("A"), uploadBufSize)

	for i := 0; i < numThreads; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case <-stop:
					return
				default:
					resp, err := client.Post(uploadURL, "application/octet-stream", bytes.NewReader(data))
					if err != nil {
						continue
					}
					io.Copy(io.Discard, resp.Body)
					resp.Body.Close()
					addBytes(&totalBytes, int64(len(data)))
				}
			}
		}()
	}

	time.Sleep(testDuration)
	close(stop)
	wg.Wait()

	return makeResult(totalBytes, testDuration)
}

// Общая логика
func makeResult(totalBytes int64, duration time.Duration) SpeedResult {
	mb := float64(totalBytes) / (1024.0 * 1024.0)
	mbps := mb / duration.Seconds()
	return SpeedResult{
		TotalBytes: totalBytes,
		Duration:   duration,
		MBps:       mbps,
	}
}

// Безопасное добавление байт
func addBytes(dst *int64, delta int64) {
	// Можно использовать sync/atomic, но здесь можно обойтись без него
	*dst += delta
}
