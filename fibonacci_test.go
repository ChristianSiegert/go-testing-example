package fibonacci

import (
	"testing"
)

var expectedResults = map[uint64]uint64{
	0:  0,
	1:  1,
	2:  1,
	3:  2,
	4:  3,
	5:  5,
	6:  8,
	7:  13,
	8:  21,
	9:  34,
	10: 55,
	11: 89,
	12: 144,
	13: 233,
	14: 377,
	15: 610,
	16: 987,
	17: 1597,
	18: 2584,
	19: 4181,
	20: 6765,
}

func resetLookupTable() {
	lookupTable = map[uint64]uint64{}
}

func TestFibonacci(t *testing.T) {
	var i uint64
	for itemCount := uint64(len(expectedResults)); i < itemCount; i++ {
		if result := Fibonacci(i); result != expectedResults[i] {
			t.Errorf("Fibonacci(%d) returned %d, expected %d", i, result, expectedResults[i])
		}
	}
}

func TestFibonacciFast(t *testing.T) {
	resetLookupTable()

	var i uint64
	for itemCount := uint64(len(expectedResults)); i < itemCount; i++ {
		if result := FibonacciFast(i); result != expectedResults[i] {
			t.Errorf("FibonacciFast(%d) returned %d, expected %d", i, result, expectedResults[i])
		}
	}
}

func TestFibonacciFastest(t *testing.T) {
	resetLookupTable()

	var i uint64
	for itemCount := uint64(len(expectedResults)); i < itemCount; i++ {
		if result := FibonacciFastest(i); result != expectedResults[i] {
			t.Errorf("FibonacciFastest(%d) returned %d, expected %d", i, result, expectedResults[i])
		}
	}
}

func BenchmarkFibonacci(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fibonacci(40)
	}
}

func BenchmarkFibonacciFast(b *testing.B) {
	resetLookupTable()

	for i := 0; i < b.N; i++ {
		FibonacciFast(40)
	}
}

func BenchmarkFibonacciFastest(b *testing.B) {
	resetLookupTable()

	for i := 0; i < b.N; i++ {
		FibonacciFastest(40)
	}
}
