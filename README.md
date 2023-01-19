# Dean's Cool Web Crawler

A CLI tool for grabbing anchors from webpages.

## Installation

`go run cmd/main.go -uri https://www.google.com`

## Syntax

This program takes one argument: a fuly-qualified URL for the webpage you wish to scrape.

`go-crawler <URL>`

This will then output a file to your current working directory.

`Anchors printed to <wd>/anchors-<timestamp>.txt! Thank you for using my cool tool!`

## Tests & Benchmarks

Benchmark specification:
```
goos: linux
goarch: amd64
pkg: github.com/deanfoley/vidsy-engineering-test/internal
cpu: Intel(R) Core(TM) i5-4300M CPU @ 2.60GHz
```

Benchmark command:
`go test --bench=. -benchtime=10s -count=5 -run=^#`

### /internal

#### PageGrabber

| Test | Average Cycles | Average ns/op | Bytes/op | Allocs/op |
|---|---|---|
| GrabWebpage | 58,840 | 204,057 | 91,433 | 77 |

#### PageParser

| Test | Average Cycles | Average ns/op | Bytes/op | Allocs/op |
|---|---|---|
| ExtractAnchors | 134,922 | 12,168 | 6,552 | 25 |
| FormatAnchors | 626,671 | 2718 | 593 | 3 |

#### UrlParser

| Test | Average Cycles | Average ns/op | Bytes/op | Allocs/op |
|---|---|---|
| ValidURL | 725,694 | 1,834 | 144 | 1 |
| InvaldURL | 1,000,000 | 1,128 | 208 | 3 |

### pprof

This project supports pprof!

Pass in a -cpuprofile and/or -memprofile flag with a desired output to output a prof file for either.

`go run main.go -uri https://www.vortex.com -cpuprofile cpu.prof -memprofile mem.prof`