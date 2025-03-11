package main

import (
	"fmt"
	"sort"
)

func main() {
	var fruits = []string{"wdnkd", "wdhiej", "whedqiojd"}
	fmt.Printf("The fruits here is ", fruits)
	fruits = append(fruits, "mwskmw2", "ekwde")
	fmt.Printf("The fruits here is ", fruits)
	fmt.Println("\n")
	fruits = append(fruits[:3])
	fmt.Printf("The fruits here is ", fruits)

	// interesting part
	highscore := make([]int, 4)
	highscore[0] = 132
	highscore[1] = 532
	highscore[2] = 432
	highscore[3] = 332
	highscore = append(highscore, 322, 211, 644)
	fmt.Println(highscore)
	sort.Ints(highscore)
	fmt.Println(highscore)

	// how to remove slice from index
	var courses = []string{"reactjs", "javascript", "swift", "python", "ruby"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
