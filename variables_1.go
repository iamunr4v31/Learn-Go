package variables_1

import "fmt"

var j int = 64 // only this format allowed

func main() {
	var k float32 = 26.0
	i := 42 // type guess
	fmt.Println(i)
	println(k)
	fmt.Printf("%v, %T", j, j)
}
