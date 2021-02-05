//Use panic like exceptions in runtime
package main

import "fmt"

func main() {
	var food = [4]string{"pizza", "ice cream", "hamburger", "tacos"}
	var num int

	fmt.Println("Enter element of array")
	fmt.Scan(&num)

	if num >= len(food) {
		panic("Num is out of bound")
	}

	fmt.Printf("%s %d: %s", "Element", num, food[num])
}
