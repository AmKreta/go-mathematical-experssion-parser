package util

func Factorial[T int8 | int16 | int32 | int64 | uint8 | uint16 | uint32 | uint64](n T) T {
	if n < 0 {
		panic("factorial is not defined for negative numbers")
	}
	if n < 2 {
		return 1
	}
	return n * Factorial(n-1)
}
