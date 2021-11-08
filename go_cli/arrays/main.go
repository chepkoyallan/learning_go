package main

import "fmt"

//creating an array of a certain type
var ArrayX [4]int
var ArrayY [2]string
var ArrayZ [1]float64

// dont forget to initialize with the equal sign
var scores = [3]int{70, 88, 90}
var names = [2]string{"Allan, Chepkoy"}

// assigning values to array positions
var scoresSoFar = [8]int{0: 70, 1: 78, 5: 76, 2, 57}

// leaving of the number of the specified arrays

var leaveOfArray = [...]int{40, 10, 0, 10}

//checking the equality of the arrays
var x = [...]int{1, 2, 3}
var y = [...]int{1, 2, 4} // I think it cheks the values not the size

// Go has single dimentional array
// simulating multidimentional arrays
var multiDimentional [5][3]int

// reading past end of an array

// writing past end of an array
// using negative indexes
// all of this will yiels a compile time error

// const myArrayOne = [2]int{1, 3}
var myArrayTwo [2]int





func main() {
	// fmt.Println(ArrayX, ArrayY, ArrayZ)
	// fmt.Println(ArrayX[3])
	// fmt.Println(scores)
	// fmt.Println(names)
	// fmt.Println(scoresSoFar)
	// fmt.Println(leaveOfArray)
	// fmt.Println(x == y)
	// fmt.Println(multiDimentional)
	// fmt.Println(myArrayOne[2])
	// myArrayTwo[0]
	fmt.Println(myArrayOne[-1])
	

}
