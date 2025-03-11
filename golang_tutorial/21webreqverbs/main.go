package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to web verb video- LCO")
	// PerformGetRequest()
	// PerformPostJsonRequest()
	PerformPostFormRequest()
}
func PerformGetRequest() {
	const myurl = "http://localhost:8000/get"
	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	var responseString strings.Builder
	fmt.Println("Status Code:", response.StatusCode)
	fmt.Println("Content length:", response.ContentLength)
	content, _ := ioutil.ReadAll(response.Body)
	byteCount,_:=responseString.Write(content)

	fmt.Println("ByteCount is:",byteCount)
	fmt.Println(responseString.String())
	fmt.Println(content)
	fmt.Println(string(content))
}

func PerformPostJsonRequest(){
	const myurl="http://localhost:8000/post"

	//fake json payload
	requestBody:=strings.NewReader(`
	{
		"coursename":"golang",
		"price":"0",
		"platform":"Learn online"
	}
	`)
	response,err:=http.Post(myurl,"application/json",requestBody)
	if err!=nil{
		panic(err)
	}
	defer response.Body.Close()
	content,_:=ioutil.ReadAll(response.Body)

	fmt.Println(string(content))


}

func PerformPostFormRequest(){
	const myurl="http://localhost:8000/postform"
	//formdata
	data:=url.Values{}
	data.Add("Name","Sahil")
	data.Add("place","mumbai")
	data.Add("email","Sahil@gmail.com")
	response,err:=http.PostForm(myurl,data)
	if err!=nil{
		panic(err)
	}
	defer response.Body.Close()

	content,_:=ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}