package main

import (
	"code-cadets-2021/homework_1/zadatak2/progressiveTax"
	"fmt"
	"log"
)

func main() {

	//intervals have to be in ascending order, not intermittent and not overlapping
	//first min amount and last max amount have to be 0
	var taxBrackets  = []progressiveTax.TaxBracket {
		{
			MinAmount: 0,
			MaxAmount: 1000,
			Tax: 0,
		},
		{
			MinAmount: 1000,
			MaxAmount: 5000,
			Tax: 10,
		},
		{
			MinAmount: 5000,
			MaxAmount: 10000,
			Tax: 20,
		},
		{
			MinAmount: 10000,
			MaxAmount: 0, //open ended interval
			Tax: 30,
		},
	}

	var amount float32 = 7000
	tax, err := progressiveTax.GetProgressiveTax(amount, taxBrackets)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Progressive tax for amount %.2f is %.2f", amount, tax)

}
