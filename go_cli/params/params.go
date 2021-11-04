package params

import (
	"fmt"
	"os"
)

// Example 5-1 Using a struct to simulate named parameters
type MyFuncOpts struct {
	FirstName string
	LastName  string
	Age       int
}

func MyFunc(opts MyFuncOpts) error {
	//do domething here
}

func main() {
	// simulation named and optional paramters
	MyFunc(MyFuncOpts{
		FirstName: "Allan",
		LastName:  "Chepkoy",
		Age:       29,
	})
	MyFunc(MyFuncOpts{
		FirstName: "Mana",
		LastName:  "Hatori",
	})

	// Variadic input parameters and slices
	fmt.Println(addTo(3))
	fmt.Println(addTo(3, 2))
	fmt.Println(addTo(3, 2, 4, 6, 8))
	a := []int{4, 3}
	fmt.Println(addTo(3, a...))
	fmt.Println(addTo(3, []int{1, 2, 3, 4, 5}...))

	//Multiple return values

	result, remainder, err := divAndRemainder(5, 2)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(result, remainder)

	// ignoring returned values
	result_2, _, err := divAndRemainder(5, 2)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(result_2)

	// named returned values
	x, y, z := namedDivAndRemainder(5, 2)
	fmt.Println(x, y, z)

	

}
