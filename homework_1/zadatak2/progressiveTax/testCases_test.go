package progressiveTax_test

import "code-cadets-2021/homework_1/zadatak2/progressiveTax"

type testCase struct {
	taxBrackets []progressiveTax.TaxBracket
	amount float32

	expectedTax float32
	expectingError bool
}

func getTestCases() []testCase {
	var testBracket1 = []progressiveTax.TaxBracket {
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

	var testBracket2 = []progressiveTax.TaxBracket {
		{
			Threshold: 0,
			Tax: 10,
		},
		{
			Threshold: 10000,
			Tax: 20,
		},
		{
			Threshold: 20000,
			Tax: 30,
		},
	}

	var testInvalidBracket1 = []progressiveTax.TaxBracket {
		{
			Threshold: 100,
			Tax: 10,
		},
		{
			Threshold: 1000,
			Tax: 20,
		},
		{
			Threshold: 3000,
			Tax: 30,
		},
	}

	var testInvalidBracket2 = []progressiveTax.TaxBracket {
		{
			Threshold: 0,
			Tax: 10,
		},
		{
			Threshold: -1000,
			Tax: 20,
		},
		{
			Threshold: 3000,
			Tax: 30,
		},
	}

	var testInvalidBracket3 = []progressiveTax.TaxBracket {
		{
			Threshold: 0,
			Tax: 10,
		},
		{
			Threshold: 3000,
			Tax: 30,
		},
		{
			Threshold: 1000,
			Tax: 20,
		},
	}

	var testInvalidBracket4 = []progressiveTax.TaxBracket {
		{
			Threshold: 0,
			Tax: 10,
		},
		{
			Threshold: 3000,
			Tax: 30,
		},
		{
			Threshold: 3000,
			Tax: 20,
		},
		{
			Threshold: 4000,
			Tax: 30,
		},
	}

	return []testCase {
		{
			taxBrackets: testBracket1,
			amount: 7000,
			expectedTax: 800,
			expectingError: false,
		},
		{
			taxBrackets: testBracket1,
			amount: 0,
			expectedTax: 0,
			expectingError: false,
		},
		{
			taxBrackets: testBracket1,
			amount: -5,
			expectedTax: 0,
			expectingError: true,
		},
		{
			taxBrackets: testBracket2,
			amount: 25000,
			expectedTax: 4500,
			expectingError: false,
		},
		{
			taxBrackets: testBracket2,
			amount: 0,
			expectedTax: 0,
			expectingError: false,
		},
		{
			taxBrackets: testBracket2,
			amount: -5,
			expectedTax: 0,
			expectingError: true,
		},
		{
			taxBrackets: testInvalidBracket1,
			amount: 1000,
			expectedTax: 0,
			expectingError: true,
		},
		{
			taxBrackets: testInvalidBracket2,
			amount: 1000,
			expectedTax: 0,
			expectingError: true,
		},
		{
			taxBrackets: testInvalidBracket3,
			amount: 1000,
			expectedTax: 0,
			expectingError: true,
		},
		{
			taxBrackets: testInvalidBracket4,
			amount: 1000,
			expectedTax: 0,
			expectingError: true,
		},
	}
}
