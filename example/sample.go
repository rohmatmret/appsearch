package main

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
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
	appsearch.ApiKey = os.Getenv("API_KEY")
	appsearch.Url = os.Getenv("URL")
	appsearch.EngineName = "catalog"
	app := appsearch.NewAppSearch(appsearch.ApiKey, appsearch.Url, appsearch.EngineName)

	payload := strings.NewReader(`{ 
		query :"nova",
		page : {current :1, size :1}
	}`)

	resp := app.Search(payload)

	var data Result
	err := json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(data)
}

func FilterDocument() {
	appsearch.ApiKey = os.Getenv("API_KEY")
	appsearch.Url = os.Getenv("URL")
	appsearch.EngineName = "catalog"
	app := appsearch.NewAppSearch(appsearch.ApiKey, appsearch.Url, appsearch.EngineName)

	payload := strings.NewReader(`{
    "query":"bobo",
    "filters": {
        "any":[
            {"all":[{"type":"book"}]}
        ]        
    },
    "page": {
        "current":1,
        "size":1
    }
	}`)

	resp := app.FilterDocument(payload)

	var data Result
	err := json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(reflect.TypeOf(data))
}

func Suggestion() {
	appsearch.ApiKey = os.Getenv("API_KEY")
	appsearch.Url = os.Getenv("URL")
	appsearch.EngineName = "catalog"
	app := appsearch.NewAppSearch(appsearch.ApiKey, appsearch.Url, appsearch.EngineName)
	resp := app.Suggestions("bobo")

	var data Result
	err := json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(reflect.TypeOf(data))
}
func FindAll() {
	appsearch.ApiKey = os.Getenv("API_KEY")
	appsearch.Url = os.Getenv("URL")
	appsearch.EngineName = "catalog"
	app := appsearch.NewAppSearch(appsearch.ApiKey, appsearch.Url, appsearch.EngineName)

	resp := app.ListDocument(nil)
	var data Result
	err := json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		panic(err)
	}
	fmt.Println(reflect.TypeOf(data))
}

func FinByID() {
	appsearch.ApiKey = os.Getenv("API_KEY")
	appsearch.Url = os.Getenv("URL")
	appsearch.EngineName = "catalog"
	app := appsearch.NewAppSearch(appsearch.ApiKey, appsearch.Url, appsearch.EngineName)
	resp := app.FindIds("239933")

	var data Result
	err := json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(reflect.TypeOf(data))
}

func main() {
	Search()         // Search with pagging
	FilterDocument() // Filter Document
	FinByID()        // Find by ID
}
