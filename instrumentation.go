package instrumentation

import "time"

func Measure(f func()) time.Duration {
	t0 := time.Now()
	f()
	return time.Since(t0)
}
