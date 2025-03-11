package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("hello files in go:\n")
	content := "This needs to go in file- linkedin.com"
	file, err := os.Create("./mylcogofile.txt")
	if err != nil {
		panic(err)
	}
	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}
	fmt.Println("length", length)
	defer file.Close()
	readFile("./mylcogofile.txt")
	// io.WriteString()
}
func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fmt.Println("data we have here is:", string(databyte))
}
