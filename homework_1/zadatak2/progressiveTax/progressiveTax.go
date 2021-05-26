package progressiveTax

import (
	"errors"
)
type TaxBracket struct {
	Threshold float32
	Tax float32 //in percentage form ie 10%
}

func validateBrackets(taxBrackets []TaxBracket) error {
	//check if first threshold is 0
	if taxBrackets[0].Threshold != 0 {
		return errors.New("first threshold must be 0")
	}
	for _, bracket := range taxBrackets {

		//check if thresholds are positive
		if bracket.Threshold < 0 {
			return errors.New("some thresholds not positive")
		}

		//check if tax values are positive
		if bracket.Tax < 0 {
			return errors.New("some tax values not positive")
		}
	}

	//check if thresholds are in ascending order
	//also assures there are no duplicates
	for i := 1; i < len(taxBrackets); i++ {
		if taxBrackets[i-1].Threshold >= taxBrackets[i].Threshold {
			return errors.New("tax brackets not in ascending order")
		}
	}
	return nil
}

func GetProgressiveTax(amount float32, taxBrackets []TaxBracket) (float32, error){
	if amount < 0 {
		return 0.0, errors.New("Invalid amount")
	}
	err := validateBrackets(taxBrackets)

	if err != nil {
		return 0.0, err
	}

	var taxAmount float32 = 0.0

	for i := 1; i < len(taxBrackets); i++ {

		if amount > taxBrackets[i].Threshold {  //amount gets full tax from bracket
			taxAmount += (taxBrackets[i].Threshold - taxBrackets[i-1].Threshold) * (taxBrackets[i-1].Tax / 100.0)

			if i == len(taxBrackets)-1 { //last element is open ended interval. Amount gets taxed the overhead
				taxAmount += (amount - taxBrackets[i].Threshold) * (taxBrackets[i].Tax / 100.0)
			}

		} else if amount > taxBrackets[i-1].Threshold { //amount gets partial tax from bracket
			taxAmount += (amount - taxBrackets[i-1].Threshold) * (taxBrackets[i-1].Tax / 100.0)
		}
	}

	return taxAmount, nil
}
