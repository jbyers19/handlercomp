# Go Handler Comparison

Handlercomp compares the handler/routing performance of several popular Go web frameworks built on `net/http`.

I specifically chose these frameworks because they are some of the most popular web frameworks for go and are all built on top of `net/http`. With the routing changes in go v1.22, I wanted to compare the performance to see if it is worth dropping 3rd party packages and just use the standard library.

**Frameworks tested:**
- [net/http](https://pkg.go.dev/net/http)
- [chi](https://github.com/go-chi/chi)
- [echo](https://github.com/labstack/echo)
- [gin](https://github.com/gin-gonic/gin)

```bash
goos: darwin
goarch: arm64
pkg: github.com/jbyers19/handlercomp
BenchmarkChi-10          3042712               376.4 ns/op           844 B/op          7 allocs/op
BenchmarkEcho-10         2962759               396.1 ns/op           845 B/op         10 allocs/op
BenchmarkGin-10          3917362               300.6 ns/op           556 B/op          5 allocs/op
BenchmarkNetHTTP-10      3208234               368.6 ns/op           563 B/op          6 allocs/op
```

## Usage

**Pre-reqs:**
- go 1.22+ installed.

**Steps:**
1. clone this repo
1. run `go mod tidy` to download the deps
1. run `go test -run '^$' -bench Bench -benchmem` to run the benchmark tests

## Notes
The configured route is the same for all the servers: `GET /user/:id`. The request path being tested is: `/user/1234?name=MyName`. This is to cover both path and query parsing by the different frameworks.

The tests re-use the request and response writer in an attempt to only test the handler performance.
Due to how Echo works, you have to explicitly set the path parameters when running the tests so that might impact the results for Echo in this test environment.

At some point I might add more tests for running a load testing tool against running servers to get a better idea of the performance comparisons.
