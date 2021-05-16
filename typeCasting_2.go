package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i int = 42 // integer declaration
	fmt.Printf("%v, %T \n", i, i)

	var j string
	j = string(i) // string conversion (rune)
	fmt.Printf("%v, %T \n", j, j)

	j = strconv.Itoa(i) // string conversion (proper)  fmt.Sprint does the same thing with more operands read docs for more info
	fmt.Printf("%v, %T \n", j, j)

}
