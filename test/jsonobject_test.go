package test

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/amirhnajafiz/explorer"
)

// creating a fake struct
type obj struct {
	name  string `json:"name"`
	value int    `json:"value"`
}

// TestObjectParsing.
func TestObjectParsing(t *testing.T) {
	// creating fake bytes
	bytes, _ := json.Marshal(&obj{
		name:  "amir",
		value: 60,
	})

	// failed parsing
	_, err := explorer.ParseJsonArray(bytes)
	if err == nil {
		t.Error(err)
	}

	// successful parsing
	objMap, err := explorer.ParseJsonObject(bytes)
	if err != nil {
		t.Error(err)
	}

	// check validation
	if objMap.Get("name").Value().(string) != "amir" {
		t.Error(errors.New("data loss during parsing"))
	}
}
