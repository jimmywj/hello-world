package main

import "fmt"
import "prj1/module1"

var a int = 10
var s string = "jimmy"

func main() {

	numbers := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(numbers)
	var k []int = numbers[1:2]
	fmt.Println(k,a)

	module1.PublicFunc(a)
	module1.PublicFunc1(s)
	fmt.Println(a)
	fmt.Println(s)
}

func add(a int, b int) int {
	return a + b
}
