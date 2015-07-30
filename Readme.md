# About
task from https://open.kattis.com/problems/namethatpermutation

## Input
Input consists of up to 1000 test cases, one per line. Each test case has two integers `1≤n≤50` and `0≤k≤n!−1`. Input ends at the end of file.

That means large number need to be handled - up to 3.0414093e+64

### How to Test
http://blog.stretchr.com/2014/03/05/test-driven-development-specifically-in-golang/

```
$ go test ./... -test.v=true
$ go test ./permutation -test.coverprofile=coverageReport.out -test.v=true
```

### Benchmarking

```
$ go test ./... -test.bench=".*" -test.v=true
```

### How to build

```
$ go build ./...
```

### How to format

```
$ go fmt ./...
```

### How to run

```
$ go run main\main.go < sample.in > answer.ans
```
