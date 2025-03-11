package main

import "fmt"

func main() {
	// no iheritance in golang
	// no super no parent
	sahil := User{"sahil", "shqidj@gmail.com", true, 21}
	fmt.Println(sahil.Age)
	// fmt.Println("%+v\n",sahil)
	fmt.Println(sahil)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
	// all are capital because we want to access outside
}
