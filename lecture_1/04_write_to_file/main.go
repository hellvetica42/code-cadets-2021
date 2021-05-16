package main

import (
	"fmt"
	"log"
	"os"

	"github.com/pkg/errors"
)

func countToTen(f *os.File) {
	for i := 10; i >= 0; i-- {
		defer f.WriteString(fmt.Sprint(i) + "\n")
	}
}

func main() {
	f, err := os.Create("output.txt")
	if err != nil {
		log.Fatal(
			errors.WithMessage(err, "opening a file"),
		)
	}

	defer f.Close()
	countToTen(f)
	f.Sync()
}
