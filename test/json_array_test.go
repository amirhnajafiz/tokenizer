package test

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/amirhnajafiz/explorer"
)

// TestArrayParsing.
func TestArrayParsing(t *testing.T) {
	collection := []obj{
		{
			Name:  "peter",
			Value: 45,
		},
		{
			Name:  "amirhossein",
			Value: 20,
		},
		{
			Name:  "daniel",
			Value: 100,
		},
	}

	bytes, _ := json.Marshal(collection)

	// check parsing
	objMap, err := explorer.ParseJsonArray(bytes)
	if err != nil {
		t.Error(err)
	}

	// check array
	if len(objMap) != 3 {
		t.Error(errors.New("data loss while parsing"))
	}
}
