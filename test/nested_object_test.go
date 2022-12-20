package test

import (
	"encoding/json"
	"errors"
	"github.com/amirhnajafiz/explorer"
	"testing"
)

// creating a fake struct
type obj2 struct {
	Name  string `json:"name"`
	Value inner  `json:"value"`
}

// the inner struct
type inner struct {
	Value int `json:"value"`
}

func TestNestedObjectParsing(t *testing.T) {
	// creating fake bytes
	bytes, _ := json.Marshal(&obj2{
		Name: "amir",
		Value: inner{
			Value: 20,
		},
	})

	// parse
	objMap, err := explorer.ParseJsonObject(bytes)
	if err != nil {
		t.Error(err)
	}

	// check validation
	if int(objMap.Get("value").Get("value").Value().(float64)) != 20 {
		t.Error(errors.New("parsing nested object failed"))
	}
}
