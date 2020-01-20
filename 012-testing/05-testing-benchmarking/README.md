## Running commands in the shell

How to write benchmarks in Go
This post continues a series on the testing package I started a few weeks back. You can read the previous article on writing table-driven tests here.

Introduction
The Go testing package contains a benchmarking facility that can be used to examine the performance of your Go code. This post explains how to use the testing package to write a simple benchmark.

You should also review the introductory paragraphs of Profiling Go programs, specifically the section on configuring power management on your machine. For better or worse, modern CPUs rely heavily on active thermal management which can add noise to benchmark results.

**Writing a benchmark**
We’ll reuse the Fib function from the previous article.

func Fib(n int) int {
        if n < 2 {
                return n
        }
        return Fib(n-1) + Fib(n-2)
}
Benchmarks are placed inside _test.go files and follow the rules of their Test counterparts. In this first example we’re going to benchmark the speed of computing the 10th number in the Fibonacci series.

// from fib_test.go
func BenchmarkFib10(b *testing.B) {
        // run the Fib function b.N times
        for n := 0; n < b.N; n++ {
                Fib(10)
        }
}
Writing a benchmark is very similar to writing a test as they share the infrastructure from the testing package. Some of the key differences are

**Benchmark functions start with Benchmark not Test**
Benchmark functions are run several times by the testing package. The value of b.N will increase each time until the benchmark runner is satisfied with the stability of the benchmark. This has some important ramifications which we’ll investigate later in this article.
Each benchmark must execute the code under test b.N times. The for loop in BenchmarkFib10 will be present in every benchmark function.
Running benchmarks
Now that we have a benchmark function defined in the tests for the fib package, we can invoke it with go test -bench=.

% go test -bench=.
PASS
BenchmarkFib10   5000000               509 ns/op
ok      github.com/davecheney/fib       3.084s
Breaking down the text above, we pass the -bench flag to go test supplying a regular expression matching everything. You must pass a valid regex to -bench, just passing -bench is a syntax error. You can use this property to run a subset of benchmarks.

The first line of the result, PASS, comes from the testing portion of the test driver, asking go test to run your benchmarks does not disable the tests in the package. If you want to skip the tests, you can do so by passing a regex to the -run flag that will not match anything. I usually use

go test -run=XXX -bench=.
The second line is the average run time of the function under test for the final value of b.N iterations. In this case, my laptop can execute Fib(10) in 509 nanoseconds. If there were additional Benchmark functions that matched the -bench filter, they would be listed here.

**Benchmarking various inputs**
As the original Fib function is the classic recursive implementation, we’d expect it to exhibit exponential behavior as the input grows. We can explore this by rewriting our benchmark slightly using a pattern that is very common in the Go standard library.

func benchmarkFib(i int, b *testing.B) {
        for n := 0; n < b.N; n++ {
                Fib(i)
        }
}

func BenchmarkFib1(b *testing.B)  { benchmarkFib(1, b) }
func BenchmarkFib2(b *testing.B)  { benchmarkFib(2, b) }
func BenchmarkFib3(b *testing.B)  { benchmarkFib(3, b) }
func BenchmarkFib10(b *testing.B) { benchmarkFib(10, b) }
func BenchmarkFib20(b *testing.B) { benchmarkFib(20, b) }
func BenchmarkFib40(b *testing.B) { benchmarkFib(40, b) }
Making benchmarkFib private avoids the testing driver trying to invoke it directly, which would fail as its signature does not match func(*testing.B). Running this new set of benchmarks gives these results on my machine.

BenchmarkFib1   1000000000               2.84 ns/op
BenchmarkFib2   500000000                7.92 ns/op
BenchmarkFib3   100000000               13.0 ns/op
BenchmarkFib10   5000000               447 ns/op
BenchmarkFib20     50000             55668 ns/op
BenchmarkFib40         2         942888676 ns/op
Apart from confirming the exponential behavior of our simplistic Fib function, there are some other things to observe in this benchmark run.

Each benchmark is run for a minimum of 1 second by default. If the second has not elapsed when the Benchmark function returns, the value of b.N is increased in the sequence 1, 2, 5, 10, 20, 50, … and the function run again.
The final BenchmarkFib40 only ran two times with the average was just under a second for each run. As the testing package uses a simple average (total time to run the benchmark function over b.N) this result is statistically weak. You can increase the minimum benchmark time using the -benchtime flag to produce a more accurate result.
% go test -bench=Fib40 -benchtime=20s
PASS
BenchmarkFib40        50         944501481 ns/op
Traps for young players
Above I mentioned the for loop is crucial to the operation of the benchmark driver. Here are two examples of a faulty Fib benchmark.

func BenchmarkFibWrong(b *testing.B) {
        for n := 0; n < b.N; n++ {
                Fib(n)
        }
}

func BenchmarkFibWrong2(b *testing.B) {
        Fib(b.N)
}
On my system BenchmarkFibWrong never completes. This is because the run time of the benchmark will increase as b.N grows, never converging on a stable value. BenchmarkFibWrong2 is similarly affected and never completes.

A note on compiler optimisations
Before concluding I wanted to highlight that to be completely accurate, any benchmark should be careful to avoid compiler optimisations eliminating the function under test and artificially lowering the run time of the benchmark.

var result int

func BenchmarkFibComplete(b *testing.B) {
        var r int
        for n := 0; n < b.N; n++ {
                // always record the result of Fib to prevent
                // the compiler eliminating the function call.
                r = Fib(10)
        }
        // always store the result to a package level variable
        // so the compiler cannot eliminate the Benchmark itself.
        result = r
}
Conclusion
The benchmarking facility in Go works well, and is widely accepted as a reliable standard for measuring the performance of Go code. Writing benchmarks in this manner is an excellent way of communicating a performance improvement, or a regression, in a reproducible way.