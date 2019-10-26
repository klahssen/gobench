package proto

import "testing"

func BenchmarkMarhshalHello(b *testing.B) {
	for n := 0; n < b.N; n++ {
		MarshalHello()
	}
}

func BenchmarkUnmarshalHello(b *testing.B) {
	for n := 0; n < b.N; n++ {
		UnmarshalHello()
	}
}

func BenchmarkMarhshalProfile(b *testing.B) {
	for n := 0; n < b.N; n++ {
		MarshalProfile()
	}
}

func BenchmarkUnmarshalProfile(b *testing.B) {
	for n := 0; n < b.N; n++ {
		UnmarshalProfile()
	}
}
