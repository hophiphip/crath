package main

import (
	"fmt"
	"math/big"

	"math.io/crath/diopheq"
	"math.io/crath/gcd"
	"math.io/crath/modular"
	"math.io/crath/primroot"
)

func main() {
	var a, b, c, d = big.NewInt(100), big.NewInt(101), big.NewInt(5), big.NewInt(3)
	fmt.Println("aaaaaa", gcd.Gcd(big.NewInt(0), big.NewInt(0)), a, a.Mod(a, b), a.Mod(a, b).Cmp(a) == 0, a.Or(a, b), a)
	fmt.Println(diopheq.Simple(c, d))
	fmt.Println(modular.GetSolution(modular.Modularfract, big.NewInt(7), big.NewInt(10), big.NewInt(19)))
	fmt.Println(primroot.Primrootp(big.NewInt(3)))
	fmt.Println(primroot.Primrootpn(big.NewInt(3), big.NewInt(2)))
	fmt.Println(primroot.Primroot2pn(big.NewInt(3), big.NewInt(2)))
	fmt.Println(modular.BinaryExponention(big.NewInt(3), big.NewInt(10)))
	fmt.Println(modular.BinaryModulo(big.NewInt(3), big.NewInt(16), big.NewInt(31)))
}
