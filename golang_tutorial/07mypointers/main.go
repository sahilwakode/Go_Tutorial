package main

import "fmt"

func main() {
	mynumber := 23
	var ptr = &mynumber
	fmt.Println("Value of pointer", ptr)
	fmt.Println("value stored is ", *ptr)
	*ptr = *ptr * 2
	fmt.Println(mynumber)
}
