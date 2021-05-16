package main

import (
	"code-cadets-2021/homework_1/zadatak1/fizzbuzz"
	"flag"
	"fmt"
	"log"
)

func main() {

	var start, end int

	flag.IntVar(&start, "start", 1, "Value (inclusive) from which to start fizzbuzz")
	flag.IntVar(&end, "end", 10, "Value (inclusive) to end fizzbuzz")

	flag.Parse()

	fizzbuzzOutput, err := fizzbuzz.FizzBuzz(start, end)

	if err != nil {
		log.Fatal(err)
	}

	for _, line := range fizzbuzzOutput{
		fmt.Println(line)
	}

}
