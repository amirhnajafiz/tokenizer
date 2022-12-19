package internal

import (
	"fmt"
)

// JsonObject is a single json structure.
type JsonObject struct {
	// object value key
	key string
	// value type
	valueType int
	// value which is generic
	value interface{}
	// object items which is a map
	items map[string]JsonObject
}

// newJsonObject generates a new base json object structure.
func newJsonObject(key string, valueType int, value interface{}) JsonObject {
	return JsonObject{
		key:       key,
		valueType: valueType,
		value:     value,
		items:     make(map[string]JsonObject),
	}
}

// Get returns an interface.
func (j JsonObject) Get(key string) JsonObject {
	// check to see if the value exists
	if item, ok := j.items[key]; !ok {
		panic(fmt.Errorf("get value failed:\n\terror:%v\n\tkey:%s", ErrKeyNotFound, key))
	} else {
		return item
	}
}

// Value returns json object value.
func (j JsonObject) Value() interface{} {
	if j.valueType == jsonArrayType || j.valueType == globalType {
		return j.value
	}

	return j.items
}
