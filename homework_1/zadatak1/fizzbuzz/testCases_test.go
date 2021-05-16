package fizzbuzz_test

type testCase struct {
	inputStart int
	inputEnd int

	expectingError bool
}

func getTestCases() []testCase {
	return []testCase {
		{
			inputStart: 1,
			inputEnd: 10,

			expectingError: false,
		},
		{
			inputStart: 20,
			inputEnd: 100,

			expectingError: false,
		},
		{
			inputStart: 0,
			inputEnd: 10,

			expectingError: true,
		},
		{
			inputStart: 10,
			inputEnd: 5,

			expectingError: true,
		},
		{
			inputStart: -5,
			inputEnd: 0,

			expectingError: true,
		},
		{
			inputStart: -10,
			inputEnd: 0,

			expectingError: true,
		},
	}
}