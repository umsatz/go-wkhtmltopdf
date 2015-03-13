# libwkhtmltox playground /w golang

`go run c.go main.go`

I've been playing around with converting documents in parallel using goroutines and
channels but the conversion either hangs indefinitly or libwkhtmltox crashes with
QT error messages (execution not in main thread).

Might need more hacks to setup a parallel pdf generation service with this...