package fizzbuzz_test

import (
	"code-cadets-2021/homework_1/zadatak1/fizzbuzz"
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	for _, tc := range getTestCases() {
		actualOutput, actualErr := fizzbuzz.FizzBuzz(tc.inputStart, tc.inputEnd)

		if tc.expectingError {
			if actualErr == nil {
				t.Errorf("Expected an error but not `nil` error")
			}
		} else {
			if actualErr != nil {
				t.Errorf("Expected no error but got non-nil error %v:", actualErr)
			}
		}

		if actualErr == nil {
			expectedOutputSize := tc.inputEnd - tc.inputStart + 1
			actualOutputSize := len(actualOutput)
			if actualOutputSize != expectedOutputSize{
				t.Errorf("Output size mismatch. Expected %d, got %d", expectedOutputSize, actualOutputSize)
			}

			for index, line := range actualOutput{
				number := tc.inputStart + index
				if number % 3 == 0 && number % 5 == 0 {
					if line != "FizzBuzz" {
						t.Errorf("Expected %s, got %s for input %d", "FizzBuzz", line, number)
					}
				} else {
					if number % 3 == 0 {
						if line != "Fizz" {
							t.Errorf("Expected %s, got %s for input %d", "Fizz", line, number)
						}
					}

					if number % 5 == 0 {
						if line != "Buzz" {
							t.Errorf("Expected %s, got %s for input %d", "Buzz", line, number)
						}
					}
				}
			}
		}
	}

}



