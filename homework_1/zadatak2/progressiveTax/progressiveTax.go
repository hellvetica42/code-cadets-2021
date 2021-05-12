package progressiveTax

import (
	"errors"
	"fmt"
)
type TaxBracket struct {
	MinAmount float32
	MaxAmount float32 //0 for open ended interval
	Tax float32 //in percentage form ie 10%
}

func validateBrackets(taxBrackets []TaxBracket) error {
	//checking lower and upper bound values
	if taxBrackets[0].MinAmount != 0 {
		return errors.New("Invalid first tax bracket. MinAmount has to be 0")
	}

	if taxBrackets[len(taxBrackets)-1].MaxAmount != 0 {
		return errors.New("Invalid last tax bracket. MaxAmount has to be 0 (open interval)")
	}

	//checking for invalid brackets
	for i, bracket := range taxBrackets {
		if bracket.MaxAmount <= bracket.MinAmount && i != len(taxBrackets)-1{
			return errors.New(fmt.Sprintf("Invalid tax bracket. %.2f <= %.2f", bracket.MaxAmount, bracket.MinAmount))
		}

		if bracket.MaxAmount < 0 || bracket.MinAmount < 0 {
			return errors.New(fmt.Sprintf("Invalid tax bracket. %.2f, %.2f", bracket.MinAmount, bracket.MaxAmount))
		}

		if bracket.Tax < 0 {
			return errors.New(fmt.Sprintf("Invalid tax amount. %.2f", bracket.Tax))
		}
	}

	//checking for intermittent and/or overlapping intervals
	for i := 0; i < len(taxBrackets)-1; i++ {
		if taxBrackets[i].MaxAmount > taxBrackets[i+1].MinAmount {
			return errors.New(fmt.Sprintf("Tax brackets %d and %d are overlapping", i+1, i+2))
		}

		if taxBrackets[i].MaxAmount < taxBrackets[i+1].MinAmount {
			return errors.New(fmt.Sprintf("Tax brackets %d and %d are not continious", i+1, i+2))
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

	for _, bracket := range taxBrackets {
		if amount < bracket.MinAmount { //amount is smaller than the bracket range
			continue
		}

		if bracket.MaxAmount > 0 && amount > bracket.MaxAmount { //amount is bigger than the bracket range
			taxAmount += (bracket.MaxAmount - bracket.MinAmount) * (bracket.Tax / 100.0)

		} else if amount >= bracket.MinAmount{ //amount is inside the bracket range (including open ended bracket)
			taxAmount += (amount - bracket.MinAmount) * (bracket.Tax / 100.0)
		}
	}

	return taxAmount, nil
}
