package main

import "fmt"

func main() {
	languages := make(map[string]string)
	languages["JS"] = "Javascript"
	languages["ReactJS"] = "React Javascript"
	languages["PY"] = "Python"
	fmt.Println(languages)
	fmt.Println(languages["PY"])
	delete(languages, "PY")
	fmt.Println(languages)

	// loops over map
	for key, values := range languages {
		fmt.Println(key, values)
	}
}
