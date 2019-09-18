# errorbench [![Build Status](https://travis-ci.com/schjan/errorbench.svg?token=a7XuXvf6EcPq4Uz5ch1B&branch=master)](https://travis-ci.com/schjan/errorbench)
Small benchmark between different errorhandling packages in Go.

```
12.07.2019 on Travis CI Go 1.12.7

BenchmarkXerrors-2      	50000000	        38.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkPkgErrors-2    	30000000	        47.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkMicroerror-2   	100000000	        20.1 ns/op	       0 B/op	       0 allocs/op
```

```
04.09.2019 on Travis CI Go 1.13
BenchmarkErrors-2           	28432845	        40.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkXerrors-2          	32176359	        37.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkPkgErrors-2        	23320749	        50.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkMicroerror-2       	56690780	        21.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkWrapError-2        	 3857890	       311 ns/op	      80 B/op	       2 allocs/op
BenchmarkWrapMicroerror-2   	  583891	      1999 ns/op	     528 B/op	       6 allocs/op
BenchmarkWrapPkgErrors-2    	 1000000	      1158 ns/op	     352 B/op	       4 allocs/op
```
