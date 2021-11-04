package functions

import "fmt"

// Functions building blocks
func div(numerator int, denominator int) int {
	if denominator == 0 {
		return 0
	}
	return numerator / denominator
}

// Multiple input parameters of the same type

func same_type(numerator, denominator int) int { return 0 }

// function with no parameters and no return values
func main() {
	result := div(5, 2)
	fmt.Println(result)
}

// Example 5-1 Using a struct to simulate named parameters
type MyFuncOpts struct {
	FirstName string
	LastNamr string
	Age int
}

func MyFunc(opts MyFuncOpts) error{
	//do domething here
}