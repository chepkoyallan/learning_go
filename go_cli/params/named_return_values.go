package params

import "errors"

func namedDivAndRemainder(numerator int, denominator int) (result int, remainder int, err error) {
	if denominator == 0 {
		err = errors.New("Cannot devide by zero")
		return result, remainder, err
	}
	result, remainder = numerator / denominator, numerator % denominator
	return result, remainder, err
}
