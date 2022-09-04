package instrumentation

import (
	"testing"
	"time"
)

func TestMeasure(t *testing.T) {
	const d = 100 * time.Millisecond
	m := Measure(func() {
		time.Sleep(d)
	})
	if m < d {
		t.Error("faild")
	}
}
