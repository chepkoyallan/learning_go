package params

import "errors"

// multiple return
func divAndRemainder(numerator int, denominator int) (int, int, error) {
	if denominator == 0 {
		return 0, 0, errors.New("Cannot devide by zero")
	}
	return numerator / denominator, numerator % denominator, nil
}
