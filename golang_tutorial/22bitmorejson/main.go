package main

import (
	"encoding/json"
	"fmt"
)

type courses struct {
	Name     string `json:"coursename"` // create name alias each place and use it
	Price    int    `json:"hello"`      // write json:name just aside it or giver error same happen below in omitempty case
	Platform string
	Password string   `json:"-"` //output is not shown for password you can see that
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Hello to Json\n")
	// EncodeJson()
	DecodeJson()
}
func EncodeJson() {
	lcoCourses := []courses{
		{
			"ReactJS Bootcamp", 299, "Learncodeonline.in", "abc123", []string{"web-dev", "js"},
		},
		{
			"Mern Bootcamp", 199, "Learncodeonline.in", "wkjddw", []string{"full-dev", "js"},
		},
		{
			"Angular Bootcamp", 213, "Learncodeonline.in", "wdjqod", nil,
		},
	}
	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)
}

func DecodeJson() {
	jsonDataFromWeb := []byte(
		`
		{
                "coursename": "ReactJS Bootcamp",
                "hello": 299,
                "Platform": "Learncodeonline.in",
                "tags": ["web-dev","js"]
        }
		`)
	var lcoCourse courses
	checkValid := json.Valid(jsonDataFromWeb)
	if checkValid {
		fmt.Println("Json was valid bro")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("Json was not valid")
	}
	// some cases where you just need to add data to key values
	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Println("%#v\n", myOnlineData)
	for k, v := range myOnlineData {
		fmt.Printf("Key is %v and value is %v and type is : %T\n", k, v, v)
	}
}
