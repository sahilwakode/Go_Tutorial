package main

import "fmt"

func main() {
	fmt.Println("hi open")
	greet()
	greet2()
	proRes := proAdder(2, 5, 4, 8)
	fmt.Println(proRes)
}
func greet2() {
	result := adder(2, 3)
	fmt.Println("another method", result)
}
func adder(val1 int, val2 int) int {
	return val1 + val2
}
func greet() {
	fmt.Println("namastey bros")
}
func proAdder(values ...int) int {
	total := 0
	for _, val := range values {
		total += val
	}
	return total
}
