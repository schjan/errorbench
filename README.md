# errorbench [![Build Status](https://travis-ci.com/schjan/errorbench.svg?token=a7XuXvf6EcPq4Uz5ch1B&branch=master)](https://travis-ci.com/schjan/errorbench)
Small benchmark between different errorhandling packages in Go.

```
12.07.2019 on Travis CI

BenchmarkXerrors-2      	50000000	        38.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkPkgErrors-2    	30000000	        47.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkMicroerror-2   	100000000	        20.1 ns/op	       0 B/op	       0 allocs/op
```