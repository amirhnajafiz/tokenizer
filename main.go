package main

import (
	"encoding/json"
	"log"
)

type JSONObject map[string]interface{}

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

	var objmap JSONObject

	if err := json.Unmarshal(bytes, &objmap); err != nil {
		log.Fatal(err)
	}

	log.Println(objmap)
	log.Println(objmap["wallet"].(JSONObject)["id"])
}
