package main

import (
	"encoding/json"
	"fmt"

	"github.com/amirhnajafiz/explorer/internal"
)

type Object struct {
	Name   string `json:"name"`
	Value  int    `json:"value"`
	Wallet Inner  `json:"wallet"`
	// Wallets []Inner `json:"wallets"`
	//Defaults []int   `json:"defaults"`
}

type Inner struct {
	Id int `json:"id"`
	//Bank SuperInner `json:"bank"`
}

//
//type SuperInner struct {
//	Name string `json:"name"`
//}

func main() {
	obj := &Object{
		Name:  "amir",
		Value: 20,
		Wallet: Inner{
			Id: 90,
		},
		//Wallets: []Inner{
		//	{
		//		Id: 20,
		//		//Bank: SuperInner{
		//		//	"asp",
		//		//},
		//	},
		//	{
		//		Id: 22,
		//		//Bank: SuperInner{
		//		//	"asp",
		//		//},
		//	},
		//},
		//Defaults: []int{1, 2, 3},
	}

	bytes, _ := json.Marshal(obj)

	objMap, err := internal.ParseJsonObject(bytes)
	if err != nil {
		panic(err)
	}

	fmt.Println(objMap.Pretty(2))
}
