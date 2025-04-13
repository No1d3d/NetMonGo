# netmongo

**netmongo** is a lightweight CLI-based network speed testing utility written in Go. It allows you to measure both **download** and **upload** speeds using parallel HTTP requests. The goal is to provide a minimal, self-contained alternative to tools like Speedtest.

---

## Features

- Parallel download speed test (default: 4 threads)
- Parallel upload speed test (default: 4 threads)
- Configurable duration (default: 15 seconds)
- Clean CLI output
- Optional local HTTP server for precise upload benchmarks

---

## Requirements

- Go 1.18 or higher

---

## Installation

Clone the repository and you are good to go:

```bash
git clone https://github.com/No1d3d/netmongo.git
cd netmongo
```

---

## Usage

### Run Download + Upload Test

```bash
go run ./main.go
```

The test will run both download and upload measurements, each lasting 15 seconds by default.

---

## Example Output

```
Starting download test...
===================================
Total Download: 215.42 MB
Time elapsed: 15.00 sec
Download speed: 14.36 Mbyte/s
===================================
Starting upload test...
===================================
Total Upload: 156.23 MB
Time elapsed: 15.00 sec
Upload speed: 10.41 Mbyte/s
===================================
```

---

## Planned Features

- Command-line flags (`--download`, `--upload`, `--duration`, `--threads`)
- Jitter and ping measurements
- Graph output or JSON logs
- GUI front-end

---
