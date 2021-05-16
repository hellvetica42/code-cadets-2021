package fizzbuzz

import (
	"github.com/pkg/errors"
	"strconv"
)

func FizzBuzz(start, end int) ([]string, error) {

	if start > end {
		return nil, errors.New("Start is greater than end")
	}

	if start <= 0 {
		return nil, errors.New("Start is <= 0")
	}

	if end <= 0 {
		return nil, errors.New("End is <= 0")
	}

	var fizzbuzzOutput []string
	for i := start; i <= end; i++ {
		tmpString := ""

		if i % 3 == 0 {
			tmpString += "Fizz"
		}

		if i % 5 == 0 {
			tmpString += "Buzz"
		}

		if len(tmpString) == 0{
			tmpString = strconv.Itoa(i)
		}

		fizzbuzzOutput = append(fizzbuzzOutput, tmpString)

	}

	return fizzbuzzOutput, nil
}
