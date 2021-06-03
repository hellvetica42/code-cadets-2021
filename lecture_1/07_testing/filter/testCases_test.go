package filter_test

type testCase struct {
	inputStart   int
	inputEnd     int
	inputDivisor int

	expectedOutput []int
	expectingError bool
}

func getTestCases() []testCase {
	return []testCase{
		{
			inputStart:   1,
			inputEnd:     10,
			inputDivisor: 1,

			expectedOutput: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			expectingError: false,
		},
		{
			inputStart:   5,
			inputEnd:     8,
			inputDivisor: 1,

			expectedOutput: []int{5, 6, 7, 8},
			expectingError: false,
		},
		{
			inputStart:   5,
			inputEnd:     12,
			inputDivisor: 3,

			expectedOutput: []int{6, 9, 12},
			expectingError: false,
		},
		{
			inputStart:   5,
			inputEnd:     12,
			inputDivisor: 30,

			expectedOutput: nil,
			expectingError: false,
		},
		{
			inputStart:   5,
			inputEnd:     2,
			inputDivisor: 1,

			expectingError: true,
		},
		{
			inputStart:   5,
			inputEnd:     10,
			inputDivisor: 0,

			expectingError: true,
		},
		{
			inputStart:   5,
			inputEnd:     10,
			inputDivisor: -2,

			expectingError: true,
		},
	}
}
