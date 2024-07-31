package basemath

import "math"

func FactorialLn(x uint) float64 {
	// for property 0! = 1 since exp(0) == 1
	if x < 1 {
		return 0
	}

	var ret float64 = 0

	for i := range x {
		ret += math.Log(float64(i))
	}

	return ret
}

func Factorial(x uint) uint {
	if x < 1 {
		return 1
	}

	return uint(math.Round(math.Exp(FactorialLn(x))))
}

func BinomialLn(n uint, k uint) float64 {
	return FactorialLn(n) - FactorialLn(k) - FactorialLn(n-k)
}

func Binomial(n uint, k uint) uint {
	return uint(math.Round(math.Exp(BinomialLn(n, k))))
}
