package internal

import (
	"encoding/json"
	"reflect"
)

// ParseJsonObject parses a json object into golang map.
func ParseJsonObject(bytes []byte) (JsonObject, error) {
	// creating a new map interface
	var mapInterface map[string]interface{}

	// unmarshalling json into variables
	if err := json.Unmarshal(bytes, &mapInterface); err != nil {
		return JsonObject{}, ErrJsonObject
	}

	return parseObject(mapInterface)
}

// ParseJsonArray parses a collection of json objects into golang map.
func ParseJsonArray(bytes []byte) ([]JsonObject, error) {
	// creating a new map interface collection
	var mapInterfaceArray []interface{}

	// unmarshalling json into variables
	if err := json.Unmarshal(bytes, &mapInterfaceArray); err != nil {
		return nil, ErrJsonObject
	}

	return parseArray(mapInterfaceArray)
}

// parseObject gets a map interface and converts it to json object.
func parseObject(obj map[string]interface{}) (JsonObject, error) {
	jObj := newJsonObject("", nil)

	for key := range obj {
		if reflect.TypeOf(obj[key]).String() == "map[string]interface {}" {
			tmp, err := parseObject(obj[key].(map[string]interface{}))
			if err != nil {
				return JsonObject{}, err
			}

			jObj.items[key] = tmp
		} else if reflect.TypeOf(obj[key]).String() == "[]interface {}" {
			tmp, err := parseArray(obj[key].([]interface{}))
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

// parseArray gets a collection of interfaces and converts it to
// a collection of json objects.
func parseArray(obj []interface{}) ([]JsonObject, error) {
	var items []JsonObject

	for _, item := range obj {
		if reflect.TypeOf(item).String() == "map[string]interface {}" {
			tmp, err := parseObject(item.(map[string]interface{}))
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
