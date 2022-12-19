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
	return JsonObject{
		items: make(map[string]JsonObject),
		value: value,
	}
}

// Get returns an interface.
func (j JsonObject) Get(key string) JsonObject {
	if item, ok := j.items[key]; !ok {
		panic(fmt.Errorf("get failed:\n\t%v\n\t%s", ErrKeyNotFound, key))
	} else {
		return item
	}
}

// set value of json object.
func (j *JsonObject) set(value interface{}) {
	j.value = value
}

// Value returns json object value.
func (j JsonObject) Value() interface{} {
	if len(j.items) == 0 {
		return j.value
	}

	return j.items
}

// JsonArray is a collection of json objects.
type JsonArray []JsonObject
