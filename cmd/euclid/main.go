package main

import (
	"applied-algebra-helpers/utils"
	"flag"
	"math/big"
)

var a int64
var b int64

func init() {
	flag.Int64Var(&a, "a", -1, "")
	flag.Int64Var(&b, "b", -1, "")

	flag.Parse()

	if b <= 1 || a <= 1 {
		panic("required values a > 1, b > 1")
	}
}

func main() {
	utils.Euclid(big.NewInt(a), big.NewInt(b))
}
