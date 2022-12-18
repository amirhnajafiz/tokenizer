package main

import (
	"encoding/json"
)

func Parse(bytes []byte) (interface{}, error) {
	var (
		jObj JsonObject
		jArr JsonArray
	)

	if err := json.Unmarshal(bytes, &jObj); err != nil {
		return nil, err
	} else if err = json.Unmarshal(bytes, &jArr); err != nil {
		return nil, err
	} else {
		return nil, ErrJsonObject
	}
}
