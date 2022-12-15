package main

import (
	"applied-algebra-helpers/utils"
	"flag"
	"math/big"
)

var base int64
var pow int64
var mod int64

func init() {
	flag.Int64Var(&base, "base", -1, "")
	flag.Int64Var(&pow, "pow", -1, "")
	flag.Int64Var(&mod, "mod", -1, "")

	flag.Parse()

	if base <= 0 || pow <= 0 || mod <= 1 {
		panic("required values base > 0, pow > 0, mod > 1")
	}
}

func main() {
	utils.Pow(big.NewInt(base), big.NewInt(pow), big.NewInt(mod))
}
