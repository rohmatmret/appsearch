## AppSearch Client for ElasticAppsearch

### Require
- Apikey
- Host


[![Go Report Card](https://goreportcard.com/badge/github.com/rohmatmret/appsearch)](https://goreportcard.com/report/github.com/rohmatmret/appsearch)


## Installation
 ```sh
 go get -u https://github.com/rohmatmret/appsearch
```

## Init configurations

```go
func main() {

	// axample init engine
	appsearch.ApiKey = "your-private-key"
	appsearch.Url = "your engine url"
	appsearch.EngineName = "enginename"
	appsearch.Connect()

	// or you can init engine with your config
	appsearch.NewAppSearch("your-private-key", "your engine url", "enginename")
}

```

> For more Information Detail, please see Official Documentation [here..](https://www.elastic.co/guide/en/app-search/current/api-reference.html)

- [AngularJS] - HTML enhanced for web apps!
- 
## Feature

- Create Engine
- Create Schema
- IndexCatalog
- Search
- Analitycs
- Suggestion


## Example 

### Create New Engine

```go

// Example for Create Engine
func main() {
	appsearch.ApiKey = "your-private-key"
	appsearch.Url = "your engine url"
	appsearch.EngineName = "enginename"
	
	app := appsearch.Connect()
	resp := app.CreateEngine("yourname_engine")
}

```

### Delete Engine
```go
	func main() {
		appsearch.ApiKey = "your-private-key"
		appsearch.Url = "your engine url"
		appsearch.EngineName = "enginename"
		
		app := appsearch.Connect()
		resp := app.DeleteEngine("yourname_engine")
	}
```