package progressiveTax_test

import (
	"code-cadets-2021/homework_1/zadatak2/progressiveTax"
	"fmt"
	"testing"
)

func TestGetProgressiveTax(t *testing.T) {
	for _, tc := range getTestCases() {
		actualOutput, actualErr := progressiveTax.GetProgressiveTax(tc.amount, tc.taxBrackets)

		if tc.expectingError {
			if actualErr == nil {
				t.Errorf("expected an error but not `nil` error")
			}
		} else {
			if actualErr != nil {
				t.Errorf("expected no error but got non-nil error %v", actualErr)
			}
		}

		if actualErr == nil {
			if actualOutput != tc.expectedTax {
				fmt.Println(tc)
				t.Errorf("expected output: %.2f. Got %.2f", tc.expectedTax, actualOutput)
			}
		}
	}

}