package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "welcome to user input boss\n"
	// fmt
	fmt.Println(welcome)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating of our pizza:\n")

	//comma ok syntax || err ok syntax
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating ", input)

}
