package fastPermutation

import (
)

func FindPermut(nInt int, iInt int) []int {
	n := uint64(nInt)
	i := uint64(iInt)
	fact := make([]uint64, n)
	perm := make([]int, n)

	fact[0] = 1
	for k := uint64(1); k < n; k++ {
		fact[k] = fact[k-1] * k
	}

	for k := uint64(0); k < n; k++ {
		perm[k] = int(i / fact[n-1-k])
		i = i % fact[n-1-k]
	}

	for k := int(n - 1); k > 0; k-- {
		for j := k - 1; j >= 0; j-- {
			if perm[j] <= perm[k] {
				perm[k]++
			}
		}
	}

	for k := uint64(0); k < n; k++ {
		perm[k]++
	}

return perm;
}
