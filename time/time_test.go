package time

import (
	"testing"
	"time"
)

func BenchmarkClock100msGetTime(b *testing.B) {
	ck := newClock(time.Millisecond * 100)
	defer ck.Stop()
	for n := 0; n < b.N; n++ {
		ck.GetTime()
	}
}

func BenchmarkTimeNow(b *testing.B) {
	for n := 0; n < b.N; n++ {
		time.Now()
	}
}

func BenchmarkGetTimex1M(b *testing.B) {
	for n := 0; n < b.N; n++ {
		getTimeFromClock(1000000)
	}
}

func BenchmarkTimeNowx1M(b *testing.B) {
	for n := 0; n < b.N; n++ {
		getTimeNow(1000000)
	}
}
