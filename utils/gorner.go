package utils

import (
	"fmt"
	"math/big"
)

func Pow(base, pow, mod *big.Int) {
	fmt.Printf("%v = ", pow)
	for i := pow.BitLen() - 1; i >= 0; i-- {
		fmt.Print(pow.Bit(i))
	}
	fmt.Print("\n\n")

	res := big.NewInt(1)

	for i := pow.BitLen() - 1; i >= 0; i-- {
		if pow.Bit(i) == 1 {
			fmt.Printf("%v * %v = ", res, base)
			res.Mul(res, base)
			res.Mod(res, mod)
			fmt.Printf("%v\n", res)
		}

		if i != 0 {
			fmt.Printf("%v^2 = ", res)
			res.Mul(res, res)
			res.Mod(res, mod)
			fmt.Printf("%v\n", res)
		}
	}

	fmt.Print("\n")

	fmt.Printf("%v^%v mod %v = %v\n", base, pow, mod, res)
}
