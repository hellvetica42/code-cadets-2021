package main

import (
	"code-cadets-2021/homework_1/zadatak2/progressiveTax"
	"fmt"
	"log"
)

func main() {

	//intervals have to be in ascending order
	//first threshold amount has to be 0
	var taxBrackets  = []progressiveTax.TaxBracket {
		{
			Threshold: 0,
			Tax: 0,
		},
		{
			Threshold: 1000,
			Tax: 10,
		},
		{
			Threshold: 5000,
			Tax: 20,
		},
		{
			Threshold: 10000,
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
