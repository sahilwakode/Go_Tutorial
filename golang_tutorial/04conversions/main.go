package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to pizza app give rating from 1 to 5")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Println(input)
	numrating, err := strconv.ParseFloat(strings.TrimSpace(input), 64) // use same syntax or \n remains in sentence
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("New rating", numrating+1)
	}
}
