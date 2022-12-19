<h1 align="center">
  Explorer
</h1>

<br />

Parse every **JSON** object in **Golang** with **Explorer**.

## Example

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
