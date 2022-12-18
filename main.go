package main

import (
	"encoding/json"
	"log"
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

	var objmap map[string]interface{}

	if err := json.Unmarshal(bytes, &objmap); err != nil {
		log.Fatal(err)
	}

	log.Println(objmap)
}
