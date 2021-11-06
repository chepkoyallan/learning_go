package div

import "errors"

func DivAndRemainder(numerator int, denominator int) (result int, remainder int, err error) {
	result, remainder = 20, 30
	if denominator == 0 {
		return 0, 0, errors.New("cannot devide by zero")
	}
	return numerator / denominator, numerator % denominator, nil
}

// Never use this
func BlankReturns(numerator, denominator int)(result int, remainder int, err error){
	if denominator == 0 {
		err = errors.New("can not devide by zero")
		return
	}
	result, remainder = numerator/denominator, numerator % denominator
	return
}