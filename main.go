package main

import (
	"encoding/json"
	"log"

	"github.com/amirhnajafiz/explorer/internal"
)

type Object struct {
	Name   string `json:"name"`
	Value  int    `json:"value"`
	Wallet Inner  `json:"wallet"`
}

type Inner struct {
	Id int `json:"id"`
}

func main() {
	obj := &Object{
		Name:  "amir",
		Value: 20,
		Wallet: Inner{
			Id: 502,
		},
	}

	bytes, _ := json.Marshal(obj)

	objMap, err := internal.ParseObject(bytes)
	if err != nil {
		panic(err)
	}

	log.Println(objMap["wallet"].(map[string]interface{})["id"])
}
