package atomic

import (
	"testing"
)

func BenchmarkAtomic(b *testing.B) {
	c := &atomicCounter{}
	for n := 0; n < b.N; n++ {
		c.increment(uint64(n))
	}
}

func BenchmarkMutex(b *testing.B) {
	c := &muxCounter{}
	for n := 0; n < b.N; n++ {
		c.increment(uint64(n))
	}
}
