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

