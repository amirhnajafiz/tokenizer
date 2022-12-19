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
func ParseArray(bytes []byte) (JsonArray, error) {
	var jArr []map[string]interface{}

	// unmarshalling json into variables
	if err := json.Unmarshal(bytes, &jArr); err != nil {
		return nil, ErrJsonObject
	}

	return parseArr(jArr)
}

func parseObj(obj map[string]interface{}) (JsonObject, error) {
	jObj := newJsonObject(nil)

	for key := range obj {
		if reflect.TypeOf(obj[key]).String() == "map[string]interface {}" {
			tmp, err := parseObj(obj[key].(map[string]interface{}))
			if err != nil {
				return JsonObject{}, err
			}

			jObj.items[key] = tmp
		} else if reflect.TypeOf(obj[key]).String() == "[]map[string]interface {}" {
			_, err := parseArr(obj[key].([]map[string]interface{}))
			if err != nil {
				return JsonObject{}, err
			}

			jObj.items[key] = JsonObject{}
		} else {
			jObj.items[key] = newJsonObject(obj[key])
		}
	}

	return jObj, nil
}

func parseArr(obj []map[string]interface{}) (JsonArray, error) {
	var items JsonArray

	for _, item := range obj {
		tmp, err := parseObj(item)
		if err != nil {
			return nil, err
		}

		items = append(items, tmp)
	}

	return items, nil
}
