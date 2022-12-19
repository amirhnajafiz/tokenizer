package internal

import (
	"fmt"
)

// JsonObject is a single json structure.
type JsonObject struct {
	items map[string]JsonObject
	value interface{}
}

// newJsonObject generates a new json object.
func newJsonObject(value interface{}) JsonObject {
	// create a map variable with nil value
	var items map[string]JsonObject

	// if object had no value then it must contain a map
	if value == nil {
		items = make(map[string]JsonObject)
	}

	return JsonObject{
		items: items,
	}
}

// Get returns an interface.
func (j *JsonObject) Get(key string) JsonObject {
	if item, ok := j.items[key]; !ok {
		panic(fmt.Errorf("get failed:\n\t%v\n\t%s", ErrKeyNotFound, key))
	} else {
		return item
	}
}

// Value returns json object value.
func (j *JsonObject) Value() interface{} {
	return j.value
}

// JsonArray is a collection of json objects.
type JsonArray []JsonObject
