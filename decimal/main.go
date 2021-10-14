package main

import (
	"fmt"
	"math/big"

	"github.com/shopspring/decimal"
)

/*
浮点数
https://github.com/shopspring/decimal
*/
func main() {
	
	num := decimal.NewFromFloat(0.0)
	fmt.Println(num)
	for i := 0 ; i < 8; i++ {
		num = num.Add(decimal.NewFromFloat(0.1))
		fmt.Println(num)
	}
	fmt.Println(num)

	fmt.Println("------------")


	n := 0.0
	fmt.Println(n)
	for i := 0; i < 8; i++ {
		n += 0.1
		fmt.Println(n)
	}
	fmt.Println(n)


	fmt.Println("------------")
	f,_ := new(big.Float).SetString("0")
	v,_ := new(big.Float).SetString("0.01")
	fmt.Println(f, v)
	for i := 0; i < 8; i++ {
		f.Add(f, v)
		fmt.Println(f)
	}
	fmt.Println(f.String())

}