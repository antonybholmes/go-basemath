package basemath

import (
	"math"
)

func LnBinomial(n uint64, k uint64) float64 {
	return LnFactorial(n) - LnFactorial(k) - LnFactorial(n-k)
}

func Binomial(n uint64, k uint64) uint64 {
	return uint64(math.Round(math.Exp(LnBinomial(n, k))))
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
	k uint64,
	N uint64,
	K uint64,
	n uint64,
) float64 {
	return LnBinomial(K, k) + LnBinomial(N-K, n-k) - LnBinomial(N, n)
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
func HypGeomPMF(k uint64, N uint64, K uint64, n uint64) float64 {
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
func HypGeomCDF(k uint64, N uint64, K uint64, n uint64) float64 {
	//console.log(k, N, K, n)

	var sum float64 = 0

	for i := range k + 1 {
		//fmt.Printf("aha %d", i)
		sum += HypGeomPMF(i, N, K, n)
	}

	return sum
}
