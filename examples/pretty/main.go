package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/amirhnajafiz/explorer"
)

type Object struct {
	Name    string    `json:"name"`
	Value   int       `json:"value"`
	Wallet  Inner     `json:"wallet"`
	Created time.Time `json:"created"`
}

type Inner struct {
	Id int32 `json:"id"`
}

func main() {
	obj := &Object{
		Name:  "amir",
		Value: 20,
		Wallet: Inner{
			Id: 90,
		},
		Created: time.Now(),
	}

	bytes, _ := json.Marshal(obj)

	objMap, err := explorer.ParseJsonObject(bytes)
	if err != nil {
		panic(err)
	}

	fmt.Println(objMap.Pretty(4))
}
