# go-testing-example

Example of how to use Go's "testing" package for testing and benchmarking.

## Running the tests and benchmarks

Change to the project's directory:

	$ cd ./go-testing-example

To run the tests:

	$ go test

To run the benchmarks:

	$ go test -test.bench .

## Reading the benchmark output

If you run the benchmarks, you will see something like this:

	1 $ go test -test.bench .
	2 PASS
	3 BenchmarkFibonacci	       	1			1045438000 ns/op
	4 BenchmarkFibonacciFast		20000000	       136 ns/op
	5 BenchmarkFibonacciFastest		50000000	      71.8 ns/op
	6 ok  	go-testing-example		7.636s

`PASS` in line 2 indicates that all tests passed. Line 3, 4 and 5 contain the results of our three benchmarks. The first column contains the benchmark's function name which is typically "Benchmark" + [name of the function being benchmarked].

The second column contains the number of runs, i.e. how often the function that was benched was called to get a reliable average. The third column shows how much time the function required on average per call (in nanoseconds).

In our case, we can see that `Fibonacci` only ran once and took 1,045,438,000 ns (or about 1 s), `FibonacciFast` ran 20,000,000 times and an operation took about 136 ns, and `FibonacciFastest` ran 50,000,000 times and an operation took on average 72 ns.

The number of runs differs between our benchmarks because Go adjusts that number automatically. Here is how `BenchmarkFibonacciFastest` looks like:

	func BenchmarkFibonacciFastest(b *testing.B) {
		for i := 0; i < b.N; i++ {
			FibonacciFastest(40)
		}
	}

In the `for` loop, we can see `b.N`. `b` is the benchmark object provided by Go's testing package. `b.N` is Go's estimate of how many times the function must be called to get a reliable average.

When we run a benchmark, `b.N` will be `1` at first. If the function takes a long time, one run is enough to time it reliably. If the function returns too fast, go reruns the benchmark with an increased N. It repeats that process until the function can be timed reliably.

In the case of `Fibonacci`, one run was enough. In the case of `FibonacciFastest`, Go started with an N of 1, then reran the benchmark with an N of 100, then 10,000, then 1,000,000 and finally 50,000,000.

Now that you know how benchmarks work in Go, check out `fibonacci_test.go` to learn how to write tests.

## Links

* [Documentation of package "testing"](http://golang.org/pkg/testing/)
* ["How to Write Go Code"](http://golang.org/doc/code.html#Testing)
