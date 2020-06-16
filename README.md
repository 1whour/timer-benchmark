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
Benchmark_antlabs_Timer_AddTimer/N-1m-16        	14327593	        77.1 ns/op
Benchmark_antlabs_Timer_AddTimer/N-5m-16        	16078015	        80.0 ns/op
Benchmark_antlabs_Timer_AddTimer/N-10m-16       	16101303	        86.5 ns/op
Benchmark_RussellLuo_Timingwheel_AddTimer/N-1m-16         	 5994146	       195 ns/op
Benchmark_RussellLuo_Timingwheel_AddTimer/N-5m-16         	 6636303	       190 ns/op
Benchmark_RussellLuo_Timingwheel_AddTimer/N-10m-16        	 6803047	       198 ns/op
Benchmark_ouqiang_Timewheel/N-1m-16                       	  861498	      1670 ns/op
Benchmark_ouqiang_Timewheel/N-5m-16                       	  685520	      1724 ns/op
Benchmark_ouqiang_Timewheel/N-10m-16                      	  695408	      1685 ns/op
Benchmark_Stdlib_AddTimer/N-1m-16                         	 6622006	       185 ns/op
Benchmark_Stdlib_AddTimer/N-5m-16                         	 7021538	       187 ns/op
Benchmark_Stdlib_AddTimer/N-10m-16                        	 7219875	       170 ns/op
PASS
ok  	benchmark	104.387s

```