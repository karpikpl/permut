package fastPermutation

import (
	"math/big"
)

func FindPermut(n int, i *big.Int) []int {
	fact := make([]*big.Int, n)
	perm := make([]int, n)

	fact[0] = big.NewInt(1)
	for k := int64(1); k < int64(n); k++ {
		fact[k] = new(big.Int).Mul(fact[k-1], big.NewInt(k))
	}

	for k := 0; k < n; k++ {
		perm[k] = int(new(big.Int).Div(i, fact[n-1-k]).Int64())
		i = new(big.Int).Mod(i, fact[n-1-k])
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
