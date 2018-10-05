package main

import "testing"

// go test -bench=. -run=^$ golearing/puzzlers/article20_2

func BenchmarkGetPrime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetPrimes(100000)
	}
}
