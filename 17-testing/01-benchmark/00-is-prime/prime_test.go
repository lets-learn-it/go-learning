package isprime

import (
	"testing"
)

// Benchmarking
func BenchmarkGeneratePrimes(b *testing.B) {
	for b.Loop() {
		GeneratePrimes(2, 100)
	}
}
