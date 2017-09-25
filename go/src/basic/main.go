package main

import "fmt"
import "basic/module1"

var a int = 10
var s string = "jimmy"

func main() {

        // local function
        add(10,20);

        // module function
	module1.PublicFunc(a)
        // public function must be capatilized
	module1.PublicFunc1(s)

        // test array
        array();

        // test scope
	fmt.Println(a) // pass by value
	fmt.Println(s) // pass by value

}

func array() {
	numbers := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(numbers)

	var k []int = numbers[1:2] // slice
	fmt.Println(k,a)
}

func looparray() {
        // loop
        var i int;
        loopmax := 10
        for i=0; i < loopmax; i++ {
          fmt.Println("loop ", i);
        } 
}

func add(a int, b int) int {
	return a + b
}
