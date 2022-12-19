package internal

import (
	"encoding/json"
	"reflect"
)

// ParseObject parses a json object into golang map.
func ParseObject(bytes []byte) (map[string]interface{}, error) {
	var jObj map[string]interface{}

	// unmarshalling json into variables
	if err := json.Unmarshal(bytes, &jObj); err != nil {
		return nil, ErrJsonObject
	}

	return parseObj(jObj)
}

// ParseArray parses a collection of json objects into golang map.
func ParseArray(bytes []byte) ([]map[string]interface{}, error) {
	var jArr []map[string]interface{}

	// unmarshalling json into variables
	if err := json.Unmarshal(bytes, &jArr); err != nil {
		return nil, ErrJsonObject
	}

	return parseArr(jArr)
}

func parseObj(obj map[string]interface{}) (map[string]interface{}, error) {
	for key := range obj {
		if reflect.TypeOf(obj[key]).String() == "map[string]interface {}" {
			tmp, err := parseObj(obj[key].(map[string]interface{}))
			if err != nil {
				return nil, err
			}

			obj[key] = tmp
		} else if reflect.TypeOf(obj[key]).String() == "[]map[string]interface {}" {
			tmp, err := parseArr(obj[key].([]map[string]interface{}))
			if err != nil {
				return nil, err
			}

			obj[key] = tmp
		}
	}

	return obj, nil
}

func parseArr(obj []map[string]interface{}) ([]map[string]interface{}, error) {
	items := make([]map[string]interface{}, len(obj))
	for _, item := range obj {
		tmp, err := parseObj(item)
		if err != nil {
			return nil, err
		}

		items = append(items, tmp)
	}

	return items, nil
}
