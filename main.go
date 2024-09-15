// Write a Golang program to sum all the numbers provided in a given range.
package main

import "fmt"

func main() {

	fmt.Println("Enter range")
	var size int
	fmt.Scan(&size)

	sum := 0

	for i := 0; i <= size; i++ {
		sum += i
	}
	fmt.Println("Sum is", sum)

}
