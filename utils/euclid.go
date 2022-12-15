package utils

import (
	"fmt"
	"math/big"
)

func Euclid(a, b *big.Int) {
	if a.Cmp(b) == 1 {
		Euclid(b, a)

		return
	}

	for a.Cmp(one) != 0 && a.Cmp(zero) != 0 {
		tmp1 := big.NewInt(0)
		tmp2 := big.NewInt(0)

		tmp1.Div(b, a)
		tmp2.Mod(b, a)

		fmt.Printf("%v = %v*%v + %v\n", b, tmp1, a, tmp2)
		b, a = a, tmp2
	}
}
