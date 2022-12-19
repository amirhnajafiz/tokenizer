package internal

import (
	"fmt"
)

// JsonObject is a single json structure.
type JsonObject struct {
	key    string
	jType  string
	items  map[string]JsonObject
	value  interface{}
	values []JsonObject
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
	if j.jType == "array" {
		return j.values
	}

	if len(j.items) == 0 {
		return j.value
	}

	return j.items
}

// Pretty returns one pretty json string.
func (j JsonObject) Pretty() string {
	return fmt.Sprintf("{\t%s\n}", j.buildPretty(1, ""))
}

// buildPretty builds json string.
func (j JsonObject) buildPretty(index int, mainKey string) string {
	if j.jType == "array" {
		limit := len(j.values)
		if limit == 0 {
			return ""
		}

		tmp := ""

		i := 0
		for _, value := range j.values {
			resp := value.Pretty()

			if i != 0 {
				tmp = fmt.Sprintf("%s,\n%s", tmp, resp)
			} else {
				tmp = resp
			}

			i++
		}

		return fmt.Sprintf("\"%s\":\n[\n\t%s\n]", mainKey, tmp)
	}

	limit := len(j.items)

	if limit == 0 {
		return fmt.Sprintf("\"%s\": \"%v\"", j.key, j.value)
	}

	tmp := ""
	i := 0
	for key := range j.items {
		res := j.items[key].buildPretty(index+1, key)

		for tab := 0; tab < index; tab++ {
			res = fmt.Sprintf("\t%s", res)
		}

		if i+1 != limit {
			res = res + ","
		}

		tmp = fmt.Sprintf("%s\n%s", tmp, res)
		i++
	}

	if mainKey != "" {
		tabs := ""
		for tab := 0; tab < index-1; tab++ {
			tabs = fmt.Sprintf("\t%s", tabs)
		}

		return fmt.Sprintf("\"%s\": {\t%s\n%s}", mainKey, tmp, tabs)
	}

	return tmp
}
