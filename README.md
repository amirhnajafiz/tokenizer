<h1 align="center">
  Explorer
</h1>

<br />

<p align="center">
    <img src="https://img.shields.io/badge/Go-1.19-00ADD8?style=for-the-badge&logo=go" alt="go version" />
    <img src="https://img.shields.io/badge/Version-0.1.0-001122?style=for-the-badge&logo=github" alt="version" />
    <img src="https://img.shields.io/github/workflow/status/amirhnajafiz/explorer/test.github-ci?label=pipeline&logo=git&style=for-the-badge" alt="pipeline" /><br />
    Parse mysterious json objects with Explorer
</p>

<br />

Parse every **JSON** object in **Golang** with **Explorer**. You can use
this tool to explore any mysterious json object. All you need to do is passing
the array of bytes to explorer, after that you can parse your json object and
find every information that you need about your mysterious json object.

## Get package

Use ```go get github.com/amirhnajafiz/explorer``` to install this package 
in your Go project.

## Examples

### Simple example

```go
package main

import (
	"encoding/json"
	"fmt"

	"github.com/amirhnajafiz/explorer"
)

type Object struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

func main() {
	obj := &Object{
		Name:  "amir",
		Value: 20,
	}

	bytes, _ := json.Marshal(obj)

	objMap, err := explorer.ParseJsonObject(bytes)
	if err != nil {
		panic(err)
	}

	fmt.Println(objMap)
}
```

### Complex Object Example

```go
package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/amirhnajafiz/explorer"
)

type Object struct {
	Name     string    `json:"name"`
	Value    int       `json:"value"`
	Wallet   Inner     `json:"wallet"`
	Wallets  []Inner   `json:"wallets"`
	Defaults []int     `json:"defaults"`
	Created  time.Time `json:"created"`
}

type Inner struct {
	Id   int32      `json:"id"`
	Bank SuperInner `json:"bank"`
}

type SuperInner struct {
	Name string `json:"name"`
}

func main() {
	obj := &Object{
		Name:  "amir",
		Value: 20,
		Wallet: Inner{
			Id: 90,
		},
		Wallets: []Inner{
			{
				Id: 20,
				Bank: SuperInner{
					"asp",
				},
			},
			{
				Id: 22,
				Bank: SuperInner{
					"asp",
				},
			},
		},
		Defaults: []int{1, 2, 3},
		Created:  time.Now(),
	}

	bytes, _ := json.Marshal(obj)

	objMap, err := explorer.ParseJsonObject(bytes)
	if err != nil {
		panic(err)
	}

	fmt.Println(objMap.Pretty(4))
	fmt.Println()
	fmt.Println(objMap.Schema())
}
```
