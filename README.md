### 硬件
cpu:  3700x

### go1.13.5
```
goos: linux
goarch: amd64
pkg: benchmark
Benchmark_antlabs_Timer_AddTimer/N-1m-16        	 9177537	       124 ns/op
Benchmark_antlabs_Timer_AddTimer/N-5m-16        	10152950	       128 ns/op
Benchmark_antlabs_Timer_AddTimer/N-10m-16       	 9955639	       127 ns/op
Benchmark_RussellLuo_Timingwheel_AddTimer/N-1m-16         	 5316916	       222 ns/op
Benchmark_RussellLuo_Timingwheel_AddTimer/N-5m-16         	 5848843	       218 ns/op
Benchmark_RussellLuo_Timingwheel_AddTimer/N-10m-16        	 5872621	       231 ns/op
Benchmark_ouqiang_Timewheel/N-1m-16                       	  720667	      1622 ns/op
Benchmark_ouqiang_Timewheel/N-5m-16                       	  807018	      1573 ns/op
Benchmark_ouqiang_Timewheel/N-10m-16                      	  666183	      1557 ns/op
Benchmark_Stdlib_AddTimer/N-1m-16                         	 8031864	       144 ns/op
Benchmark_Stdlib_AddTimer/N-5m-16                         	 8437442	       151 ns/op
Benchmark_Stdlib_AddTimer/N-10m-16                        	 8080659	       167 ns/op

```

### go1.14.4
```
goos: linux
goarch: amd64
pkg: benchmark
Benchmark_antlabs_Timer_AddTimer/N-1m-16        	12083608	        92.5 ns/op
Benchmark_antlabs_Timer_AddTimer/N-5m-16        	13804861	        90.1 ns/op
Benchmark_antlabs_Timer_AddTimer/N-10m-16       	14031312	        92.3 ns/op
Benchmark_Stdlib_AddTimer/N-1m-16               	 7007185	       178 ns/op
Benchmark_Stdlib_AddTimer/N-5m-16               	 9036193	       166 ns/op
Benchmark_Stdlib_AddTimer/N-10m-16              	 7071096	       185 ns/op
Benchmark_RussellLuo_Timingwheel_AddTimer/N-1m-16         	 6584066	       193 ns/op
Benchmark_RussellLuo_Timingwheel_AddTimer/N-5m-16         	 6649886	       192 ns/op
Benchmark_RussellLuo_Timingwheel_AddTimer/N-10m-16        	 6103797	       181 ns/op
Benchmark_ouqiang_Timewheel/N-1m-16                       	  946098	      1638 ns/op
Benchmark_ouqiang_Timewheel/N-5m-16                       	  728776	      1516 ns/op
Benchmark_ouqiang_Timewheel/N-10m-16                      	  736267	      1539 ns/op
PASS
ok  	benchmark	108.899s

```
