package basemath

import (
	"fmt"
	"math"
)

func FactorialLn(n uint) float64 {
	// for property 0! = 1 since exp(0) == 1
	if n == 0 {
		return 1
	}

	var ret float64 = 0

	for i := range n {
		ret += math.Log(float64(i + 1))
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

/**
 * Calculates the probabilty mass function in log space.
 *
 * @param k number of observed successes in n
 * @param N population size (total number of items)
 * @param K number of successes in population
 * @param n number of draws, i.e. how many items we select each time
 * @returns probability of arrangement occuring by chance
 */
func LnHypgeomPMF(
	k uint,
	N uint,
	K uint,
	n uint,
) float64 {
	return BinomialLn(K, k) + BinomialLn(N-K, n-k) - BinomialLn(N, n)
}

/**
 * Calculates the probabilty mass function.
 *
 * @param k number of observed successes in n
 * @param N population size (total number of items)
 * @param K number of successes in population
 * @param n number of draws, i.e. how many items we select each time
 * @returns probability of arrangement occuring by chance
 */
func HypGeomPMF(k uint, N uint, K uint, n uint) float64 {
	return math.Exp(LnHypgeomPMF(k, N, K, n))
}

/**
 * Calculates the hypergeometric cumulative distribution function.
 *
 * @param k number of observed successes in n
 * @param N population size (total number of items)
 * @param K number of successes in population
 * @param n number of draws, i.e. how many items we select each time
 * @returns probability of arrangement occuring by chance
 */
func HypGeomCDF(k uint, N uint, K uint, n uint) float64 {
	//console.log(k, N, K, n)

	var sum float64 = 0

	for i := range k {
		fmt.Printf("%d", i)
		sum += HypGeomPMF(i, N, K, n)
	}

	return sum
}
