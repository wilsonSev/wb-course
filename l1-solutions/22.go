package main

import (
	"fmt"
	"math/big"
)

func main() {
	var sa, sb string
	fmt.Scan(&sa, &sb)
	a := new(big.Int)
	b := new(big.Int)
	a.SetString(sa, 10)
	b.SetString(sb, 10)

	res := new(big.Int)

	fmt.Println("a + b = ", res.Add(a, b))
	fmt.Println("a - b = ", res.Sub(a, b))
	fmt.Println("a * b = ", res.Mul(a, b))
	fmt.Println("a / b = ", res.Div(a, b))
}
