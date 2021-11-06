package main

import (
	"fmt"

	div "github.com/chepkoyallan/learning_go/go_cli/returns/divmod"
)

func main() {
	// named return values
	fmt.Println(div.DivAndRemainder(5, 2))

	// blank returns
	fmt.Println(div.BlankReturns(5, 2))
}
