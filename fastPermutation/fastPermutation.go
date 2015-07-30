package fastPermutation

import (
	"math"
)

func FindPermut(n int, i float64) []int {
	fact := make([]float64, n)
	perm := make([]int, n)

	fact[0] = 1
	for k := 1; k < n; k++ {
		fact[k] = math.Floor(fact[k-1] * float64(k))
	}

	for k := 0; k < n; k++ {
		perm[k] = int(i / fact[n-1-k])
		i = math.Floor(math.Mod(i, fact[n-1-k]))
	}

	for k := int(n - 1); k > 0; k-- {
		for j := k - 1; j >= 0; j-- {
			if perm[j] <= perm[k] {
				perm[k]++
			}
		}
	}

	for k := 0; k < n; k++ {
		perm[k]++
	}

	return perm
}
