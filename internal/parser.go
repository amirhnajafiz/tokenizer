package internal

import (
	"encoding/json"
	"reflect"
)

// ParseObject parses a json object into golang map.
func ParseObject(bytes []byte) (JsonObject, error) {
	var jObj map[string]interface{}

	// unmarshalling json into variables
	if err := json.Unmarshal(bytes, &jObj); err != nil {
		return JsonObject{}, ErrJsonObject
	}

	return parseObj(jObj)
}

// ParseArray parses a collection of json objects into golang map.
func ParseArray(bytes []byte) ([]JsonObject, error) {
	var jArr []interface{}

	// unmarshalling json into variables
	if err := json.Unmarshal(bytes, &jArr); err != nil {
		return nil, ErrJsonObject
	}

	return parseArr(jArr)
}

func parseObj(obj map[string]interface{}) (JsonObject, error) {
	jObj := newJsonObject("", nil)

	for key := range obj {
		if reflect.TypeOf(obj[key]).String() == "map[string]interface {}" {
			tmp, err := parseObj(obj[key].(map[string]interface{}))
			if err != nil {
				return JsonObject{}, err
			}

			jObj.items[key] = tmp
		} else if reflect.TypeOf(obj[key]).String() == "[]interface {}" {
			tmp, err := parseArr(obj[key].([]interface{}))
			if err != nil {
				return JsonObject{}, err
			}

			jTmp := newJsonObject("", nil)
			jTmp.jType = "array"
			jTmp.values = tmp

			jObj.items[key] = jTmp
		} else {
			jObj.items[key] = newJsonObject(key, obj[key])
		}
	}

	return jObj, nil
}

func parseArr(obj []interface{}) ([]JsonObject, error) {
	var items []JsonObject

	for _, item := range obj {
		if reflect.TypeOf(item).String() == "map[string]interface {}" {
			tmp, err := parseObj(item.(map[string]interface{}))
			if err != nil {
				return nil, err
			}
			items = append(items, tmp)
		} else {
			items = append(items, newJsonObject("", item))
		}
	}

	return items, nil
}
