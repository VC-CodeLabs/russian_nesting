includes various test cases- these can be executed manually using file redirection or with the supplied executor testSamples.go, which will also validate the results (see Test Cases section below)

use -h parameter for various command-line options

`go run testSamples.go -h`

# Test Cases
`testCases*.txt` define various sets of test cases, and specify the expected max nesting envelopes count for each

see the comments in each file for details on the various test cases

# To Thread or Not to Thread

One area this helped with (besides basic validation of various edge cases) was determining the benefits of threading.

```
go run testSamples.go
Read 15 test cases from answerKey.txt
Threaded: true
testing JeffR_Alek_Example1.txt expected 3 actual 3 in 521.6µs
testing JeffR_Alek_Example2.txt expected 1 actual 1 in 0s
testing JeffR_Alek_Example3.txt expected 2 actual 2 in 1.0863ms
testing JeffR_EndsFit.txt expected 2 actual 2 in 1.0543ms
testing JeffR_HeadWontFit.txt expected 3 actual 3 in 527.4µs
testing JeffR_Jumbled.txt expected 3 actual 3 in 0s
testing JeffR_MaxAllMatch.txt expected 1 actual 1 in 7.4529ms
testing JeffR_MaxAllUnique.txt expected 100000 actual 100000 in 101.0548ms
testing JeffR_MaxChunked.txt expected 10 actual 10 in 1.4917314s
testing JeffR_MiddleFits.txt expected 3 actual 3 in 1.0336ms
testing JeffR_Multiline.txt expected 3 actual 3 in 527µs
testing JeffR_Reverse.txt expected 5 actual 5 in 0s
testing JeffR_Singleton.txt expected 1 actual 1 in 64.9µs
testing JeffR_TailWontFit.txt expected 3 actual 3 in 512µs
testing JeffR_Whitespace.txt expected 3 actual 3 in 0s
All tests completed in 1.6055662s in 15 total runs with an average run of 107.037746ms

```
vs threading disabled:

```
go run testSamples.go -t=false
Read 15 test cases from answerKey.txt
Threaded: false
testing JeffR_Alek_Example1.txt expected 3 actual 3 in 0s
testing JeffR_Alek_Example2.txt expected 1 actual 1 in 0s
testing JeffR_Alek_Example3.txt expected 2 actual 2 in 521.4µs
testing JeffR_EndsFit.txt expected 2 actual 2 in 524.5µs
testing JeffR_HeadWontFit.txt expected 3 actual 3 in 527.2µs
testing JeffR_Jumbled.txt expected 3 actual 3 in 527.1µs
testing JeffR_MaxAllMatch.txt expected 1 actual 1 in 5.7997ms
testing JeffR_MaxAllUnique.txt expected 100000 actual 100000 in 8.8245ms
testing JeffR_MaxChunked.txt expected 10 actual 10 in 8.2047951s
testing JeffR_MiddleFits.txt expected 3 actual 3 in 0s
testing JeffR_Multiline.txt expected 3 actual 3 in 0s
testing JeffR_Reverse.txt expected 5 actual 5 in 5.08ms
testing JeffR_Singleton.txt expected 1 actual 1 in 0s
testing JeffR_TailWontFit.txt expected 3 actual 3 in 0s
testing JeffR_Whitespace.txt expected 3 actual 3 in 0s
All tests completed in 8.2265995s in 15 total runs with an average run of 548.439966ms
```

so, running all the samples we get a better overall average ~4-5x faster even tho max "all match" is slightly slower; this ratio seems to hold up when repeating each test N times:

repeat each test 3 times, threading enabled:
```
go run testSamples.go -r=3
... // output omitted
All tests completed in 15.5373184s in 45 total runs with an average run of 345.273742ms
```

repeat each test 3 times, threading disabled:
```
go run testSamples.go -r=3 -t=false
... // output omitted
All tests completed in 49.6148058s in 45 total runs with an average run of 1.10255124s
```

repeat each test 10 times, threading enabled:
```
go run testSamples.go -r=10
... // output omitted
All tests completed in 1m37.7825186s in 150 total runs with an average run of 651.883457ms
```

repeat each test 10 times, threading disabled:
```
go run testSamples.go -r=10 -t=false
... // output omitted
All tests completed in 7m38.2047615s in 150 total runs with an average run of 3.05469841s
```

NOTE these timings are for the algorithm only, doesn't include compile + launch + console output time, so Alek will no doubt get different results when using his test harness but hopefully equivalent relative performance; the benefit of adding threading support will depend entirely on the test cases used.



