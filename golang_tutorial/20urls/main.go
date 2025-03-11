package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://google.com:2000/learn?coursename:2982904u"

func main() {
	fmt.Println(myurl)
	//parsing
	result, _ := url.Parse(myurl)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery)
	fmt.Println(result.Port())
	qparams := result.Query()
	fmt.Printf("The type of query params is %T\n", qparams)
	for _, val := range qparams {
		fmt.Println(val)
	}

	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=sahil",
	}
	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)
}
