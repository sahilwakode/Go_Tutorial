package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	dicenum := rand.Intn(6) + 1
	fmt.Println(dicenum)
	switch dicenum{
	case 1:
		fmt.Println("Dice 1")
	case 2:
		fmt.Println("Dice 2")
	case 3:
		fmt.Println("Dice 3")
	case 4:
		fmt.Println("Dice 4")
	case 5:
		fmt.Println("Dice 5")
	case 6:
		fmt.Println("Dice 6")
	}
}
