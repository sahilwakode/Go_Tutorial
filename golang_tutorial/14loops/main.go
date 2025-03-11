package main

import "fmt"

func main() {
	days := []string{"Sunday", "Tuesday", "Wednesday", "Thursday"}
	fmt.Println(days)
	for d := 0; d < len(days); d++ {
		fmt.Println(days[d])
	}
	for i := range days {
		fmt.Println(days[i])
	}
	rougeval := 1
	for rougeval < 10 {
		if rougeval == 8 {
			goto lco
		}
		fmt.Println(rougeval)
		rougeval++
	}
	lco:
		fmt.Println("hello")

		
	fmt.Println("bahar")
}
