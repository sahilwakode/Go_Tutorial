package main

import "fmt"

func main() {
	var fruits [4]string
	fruits[0] = "Apple"
	fruits[3] = "Bans"
	fruits[2] = "Mango"
	fmt.Println("fruits: ", fruits)
	fmt.Println("fruits: ", len(fruits))

	var veg = [3]string{"siajs", "wjdi3", "ishdi"}
	fmt.Println(veg)
}
