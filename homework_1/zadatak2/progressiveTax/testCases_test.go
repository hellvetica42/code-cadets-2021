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
			MinAmount: 0,
			MaxAmount: 1000,
			Tax: 0,
		},
		{
			MinAmount: 1000,
			MaxAmount: 5000,
			Tax: 10,
		},
		{
			MinAmount: 5000,
			MaxAmount: 10000,
			Tax: 20,
		},
		{
			MinAmount: 10000,
			MaxAmount: 0, //open ended interval
			Tax: 30,
		},
	}

	var testBracket2 = []progressiveTax.TaxBracket {
		{
			MinAmount: 0,
			MaxAmount: 10000,
			Tax: 10,
		},
		{
			MinAmount: 10000,
			MaxAmount: 20000,
			Tax: 20,
		},
		{
			MinAmount: 20000,
			MaxAmount: 0, //open ended interval
			Tax: 30,
		},
	}

	var testIntermittentBracket = []progressiveTax.TaxBracket { //intermittent bracket
		{
			MinAmount: 0,
			MaxAmount: 10000,
			Tax: 10,
		},
		{
			MinAmount: 15000,
			MaxAmount: 20000,
			Tax: 20,
		},
		{
			MinAmount: 20000,
			MaxAmount: 0, //open ended interval
			Tax: 30,
		},
	}

	var testOverlappingBracket = []progressiveTax.TaxBracket { //overlapping bracket
		{
			MinAmount: 0,
			MaxAmount: 10000,
			Tax: 10,
		},
		{
			MinAmount: 10000,
			MaxAmount: 22000,
			Tax: 20,
		},
		{
			MinAmount: 20000,
			MaxAmount: 0, //open ended interval
			Tax: 30,
		},
	}

	var testInvalidRangeBracket = []progressiveTax.TaxBracket {
		{
			MinAmount: 100,
			MaxAmount: 10000,
			Tax: 10,
		},
		{
			MinAmount: 10000,
			MaxAmount: 22000,
			Tax: 20,
		},
		{
			MinAmount: 20000,
			MaxAmount: 30000, //open ended interval
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
			taxBrackets: testOverlappingBracket,
			amount: 1000,
			expectedTax: 0,
			expectingError: true,
		},
		{
			taxBrackets: testIntermittentBracket,
			amount: 1000,
			expectedTax: 0,
			expectingError: true,
		},
		{
			taxBrackets: testInvalidRangeBracket,
			amount: 1000,
			expectedTax: 0,
			expectingError: true,
		},
	}
}
