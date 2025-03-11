package main

import "fmt"

func main() {
	// for i := range 3 {
	// 	defer fmt.Println(i)
	// }
	defer fmt.Println("World")
	defer fmt.Println("1")
	defer fmt.Println("2")
	fmt.Println("Hello")
	mydef()
	// last in first out method here

}
func mydef() {
	for i := range 5 {
		defer fmt.Println(i)
	}
}
