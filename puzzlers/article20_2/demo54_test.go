package main

import "testing"

func BenchmarkGetPrimes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetPrimes(100000)
	}
}
