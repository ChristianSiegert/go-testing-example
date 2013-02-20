package fibonacci

// lookupTable stores computed results from FibonacciFast or FibonacciFastest.
var lookupTable = map[uint64]uint64{}

// Fibonacci computes the Fibonacci number for n.
func Fibonacci(n uint64) uint64 {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}

	return Fibonacci(n-1) + Fibonacci(n-2)
}

// FibonacciFast is similar to Fibonacci, with the only difference that computed
// results are stored in a lookup table. This makes FibonacciFast several
// million times faster than Fibonacci.
func FibonacciFast(n uint64) uint64 {
	if _, resultExists := lookupTable[n]; resultExists {
		return lookupTable[n]
	}

	var result uint64

	if n == 0 {
		result = 0
	} else if n == 1 {
		result = 1
	} else {
		result = FibonacciFast(n-1) + FibonacciFast(n-2)
	}

	lookupTable[n] = result
	return result
}

// FibonacciFastest is similar to FibonacciFast, with the only difference that
// the lookup table is accessed only once. This makes FibonacciFastest twice as
// fast as FibonacciFast.
func FibonacciFastest(n uint64) uint64 {
	if result, resultExists := lookupTable[n]; resultExists {
		return result
	}

	var result uint64

	if n == 0 {
		result = 0
	} else if n == 1 {
		result = 1
	} else {
		result = FibonacciFastest(n-1) + FibonacciFastest(n-2)
	}

	lookupTable[n] = result
	return result
}
