package concurrency

import (
	"testing"
	"time"
)

func BenchmarkSpawnGoSleep10x5ms(b *testing.B) {
	for n := 0; n < b.N; n++ {
		spawnGoSleep(10, time.Millisecond*5)
	}
}
func BenchmarkSpawnGoSleep100x5ms(b *testing.B) {
	for n := 0; n < b.N; n++ {
		spawnGoSleep(100, time.Millisecond*5)
	}
}
func BenchmarkSpawnGoSleepWG10x5ms(b *testing.B) {
	for n := 0; n < b.N; n++ {
		spawnGoSleepWG(10, time.Millisecond*5)
	}
}
func BenchmarkSpawnGoSleepWG100x5ms(b *testing.B) {
	for n := 0; n < b.N; n++ {
		spawnGoSleepWG(100, time.Millisecond*5)
	}
}
