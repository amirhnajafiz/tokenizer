package internal

import (
	"fmt"
)

// JsonObject is a single json structure.
type JsonObject struct {
	key   string
	items map[string]JsonObject
	value interface{}
}

// newJsonObject generates a new json object.
func newJsonObject(key string, value interface{}) JsonObject {
	return JsonObject{
		key:   key,
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

// Pretty returns one pretty json string.
func (j JsonObject) Pretty() string {
	return j.buildPretty(0)
}

// buildPretty builds json string.
func (j JsonObject) buildPretty(index int) string {
	tmp := ""
	limit := len(j.items)

	if limit == 0 {
		return fmt.Sprintf("%s: %s", j.key, j.value)
	}

	i := 0
	for key := range j.items {
		res := j.items[key].buildPretty(index + 1)

		for tab := 0; tab < index; tab++ {
			res = fmt.Sprintf("\t%s", res)
		}

		if i+1 != limit {
			res = res + ","
		}

		tmp = fmt.Sprintf("%s\n%s", tmp, res)
		i++
	}

	return fmt.Sprintf("{\n\t%s\n}", tmp)
}

// JsonArray is a collection of json objects.
type JsonArray []JsonObject
