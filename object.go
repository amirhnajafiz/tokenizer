package explorer

import (
	"fmt"
)

// JsonObject is a single json structure.
type JsonObject struct {
	// object value key
	key string
	// value type
	valueType int
	// type of the object value
	typeString string
	// value which is generic
	value interface{}
	// object items which is a map
	items map[string]JsonObject
}

// newJsonObject generates a new base json object structure.
func newJsonObject(key string, valueType int, value interface{}, typeString ...string) JsonObject {
	ts := ""

	if len(typeString) > 0 {
		ts = typeString[0]
	}

	return JsonObject{
		key:        key,
		valueType:  valueType,
		typeString: ts,
		value:      value,
		items:      make(map[string]JsonObject),
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

// Type returns object value type.
func (j JsonObject) Type() string {
	return j.typeString
}

// Value returns json object value.
func (j JsonObject) Value() interface{} {
	if j.valueType == jsonArrayType || j.valueType == globalType {
		return j.value
	}

	return j.items
}
