package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/rohmatmret/appsearch"
)

type Result interface{}

// Example for Create Engine
func CreateEngine() {
	app := appsearch.Connect()
	resp := app.CreateEngine("test")
	fmt.Println(resp)
}

// example for delete engine
func DeleteEngine() {
	app := appsearch.Connect()
	resp := app.DeleteEngine("test")
	fmt.Println(resp)
}

// example for get list Document
func ListDocument() {
	app := appsearch.Connect()
	resp := app.ListDocument(nil)

	var data Result
	err := json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}

// example for search with pagging
func Search() {
	app := appsearch.Connect()
	page := strings.NewReader(`{page :{current:1,size:10}}`)
	resp := app.Search("keyword", page)

	var data Result
	err := json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}

func main() {

	// axample init engine
	appsearch.ApiKey = "your-private-key"
	appsearch.Url = "your engine url"
	appsearch.EngineName = "enginename"
	appsearch.Connect()

	// or you can init engine with your config
	appsearch.NewAppSearch("your-private-key", "your engine url", "enginename")
}
