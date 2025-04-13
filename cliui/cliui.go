package cliui

import (
	"fmt"
	"netmongo/netfunc"
)

func RunDownloadTest() {
	fmt.Println("Starting download test...")
	result := netfunc.DownloadSpeedTest()
	printResult("Download", result)
}

func RunUploadTest() {
	fmt.Println("Starting upload test...")
	result := netfunc.UploadSpeedTest()
	printResult("Upload", result)
}

func printResult(label string, result netfunc.SpeedResult) {
	fmt.Println("===================================")
	fmt.Printf("Total %s: %.2f MB\n", label, float64(result.TotalBytes)/(1024*1024))
	fmt.Printf("Time elapsed: %.2f sec\n", result.Duration.Seconds())
	fmt.Printf("%s speed: %.2f Mbyte/s\n", label, result.MBps)
	fmt.Println("===================================")
}
