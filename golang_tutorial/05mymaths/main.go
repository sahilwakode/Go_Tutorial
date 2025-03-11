package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//random no.s
	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Intn(3))
}
