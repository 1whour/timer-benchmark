package benchmark

import (
	"testing"
	"time"

	"github.com/RussellLuo/timingwheel"
	"github.com/antlabs/timer"
	"github.com/ouqiang/timewheel"
)

func genD(i int) time.Duration {
	return time.Duration(i) * 10 * time.Millisecond
	//return time.Duration(i%10000) * time.Millisecond
}

func Benchmark_antlabs_Timer_AddTimer(b *testing.B) {

	tw := timer.NewTimer()
	go tw.Run()
	defer tw.Stop()
	b.ResetTimer()

	cases := []struct {
		name string
		N    int // the data size (i.e. number of existing timers)
	}{
		{"N-1m", 1000000},
		{"N-5m", 5000000},
		{"N-10m", 10000000},
	}
	for _, c := range cases {
		b.Run(c.name, func(b *testing.B) {
			base := make([]timer.TimeNoder, c.N)
			for i := 0; i < len(base); i++ {
				base[i] = tw.AfterFunc(genD(i), func() {})
			}
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				//tw.AfterFunc(time.Second, func() {})
				tw.AfterFunc(time.Second, func() {}).Stop()
			}

			b.StopTimer()
			for i := 0; i < len(base); i++ {
				base[i].Stop()
			}
		})
	}
}

func Benchmark_RussellLuo_Timingwheel_AddTimer(b *testing.B) {
	tw := timingwheel.NewTimingWheel(time.Millisecond, 20)
	tw.Start()
	defer tw.Stop()

	cases := []struct {
		name string
		N    int // the data size (i.e. number of existing timers)
	}{
		{"N-1m", 1000000},
		{"N-5m", 5000000},
		{"N-10m", 10000000},
	}
	for _, c := range cases {
		b.Run(c.name, func(b *testing.B) {
			base := make([]*timingwheel.Timer, c.N)
			for i := 0; i < len(base); i++ {
				base[i] = tw.AfterFunc(genD(i), func() {})
			}
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				tw.AfterFunc(time.Second, func() {}).Stop()
			}

			b.StopTimer()
			for i := 0; i < len(base); i++ {
				base[i].Stop()
			}
		})
	}
}

func Benchmark_ouqiang_Timewheel(b *testing.B) {
	tw := timewheel.New(1*time.Second, 3600, func(data interface{}) {
		// do something
	})

	tw.Start()
	defer tw.Stop()

	cases := []struct {
		name string
		N    int // the data size (i.e. number of existing timers)
	}{
		{"N-1m", 1000000},
		{"N-5m", 5000000},
		{"N-10m", 10000000},
	}
	for _, c := range cases {
		b.Run(c.name, func(b *testing.B) {
			for i := -c.N; i < 0; i++ {
				tw.AddTimer(genD(i), i, i)
			}
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				tw.AddTimer(time.Second, i, i)
				tw.RemoveTimer(i)
			}

			b.StopTimer()
			for i := -c.N; i < 0; i++ {
				tw.RemoveTimer(i)
			}
		})
	}
}

func Benchmark_Stdlib_AddTimer(b *testing.B) {
	cases := []struct {
		name string
		N    int // the data size (i.e. number of existing timers)
	}{
		{"N-1m", 1000000},
		{"N-5m", 5000000},
		{"N-10m", 10000000},
	}
	for _, c := range cases {
		b.Run(c.name, func(b *testing.B) {
			base := make([]*time.Timer, c.N)
			for i := 0; i < len(base); i++ {
				base[i] = time.AfterFunc(genD(i), func() {})
			}
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				time.AfterFunc(time.Second, func() {}).Stop()
			}

			b.StopTimer()
			for i := 0; i < len(base); i++ {
				base[i].Stop()
			}
		})
	}
}
