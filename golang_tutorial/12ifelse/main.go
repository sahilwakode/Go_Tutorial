package main

import "fmt"

func main() {
	loginCount := 23
	var result string
	if loginCount < 10 {
		result = "regular user"
	} else if loginCount > 10 {
		result = "watch out"
	} else {
		result = "exact 1"
	}
	fmt.Println(result)

	if num := 3; num < 10 {
		fmt.Println("num less than 10")
	} else {
		fmt.Println("num greater than 10")
	}
	// if err!=nil{

	// }
}
