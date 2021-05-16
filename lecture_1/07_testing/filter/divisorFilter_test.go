package filter_test

import (
	"testing"

	"code-cadets-2021/lecture_1/07_testing/filter"
)

func areSlicesEqual(first []int, second []int) bool {
	if len(first) != len(second) {
		return false
	}

	for idx, x := range first {
		if x != second[idx] {
			return false
		}
	}

	return true
}

func TestGetDivisibleFromRange(t *testing.T) {
	for _, tc := range getTestCases() {
		actualOutput, actualErr := filter.GetDivisibleFromRange(tc.inputStart, tc.inputEnd, tc.inputDivisor)

		if tc.expectingError {
			if actualErr == nil {
				t.Errorf("Expected an error but got `nil` error")
			}
		} else {
			if actualErr != nil {
				t.Errorf("Expected no error but got non-nil error: %v", actualErr)
			}

			if !areSlicesEqual(actualOutput, tc.expectedOutput) {
				t.Errorf(
					"Actual and expected output is not equal - actual: %v, expected: %v",
					actualOutput,
					tc.expectedOutput)
			}
		}
	}
}
