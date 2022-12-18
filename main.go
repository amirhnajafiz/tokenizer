package main

import (
	"encoding/json"
	"log"
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

	var objmap map[string]interface{}

	if err := json.Unmarshal(bytes, &objmap); err != nil {
		log.Fatal(err)
	}

	log.Println(objmap)
	log.Println(objmap["wallet"].(map[string]interface{})["id"])
}
