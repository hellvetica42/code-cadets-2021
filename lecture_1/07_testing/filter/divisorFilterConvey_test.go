package filter_test

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"code-cadets-2021/lecture_1/07_testing/filter"
)

// NOTE - Convey infix in the function name is here just to prevent a name
// clash with the method in `divisorFilter_test.go`
func TestConveyGetDivisibleFromRange(t *testing.T) {

	for idx, tc := range getTestCases() {
		Convey(fmt.Sprintf("Given test case #%v: %+v", idx, tc), t, func() {

			actualOutput, actualErr := filter.GetDivisibleFromRange(tc.inputStart, tc.inputEnd, tc.inputDivisor)

			if tc.expectingError {
				So(actualErr, ShouldNotBeNil)
			} else {
				So(actualErr, ShouldBeNil)
				So(actualOutput, ShouldResemble, tc.expectedOutput)
			}

		})
	}
}
