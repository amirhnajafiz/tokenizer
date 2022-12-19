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
func parseObject(object map[string]interface{}) (JsonObject, error) {
	// create a new json object
	tempJsonObject := newJsonObject("", jsonObjectType, nonSingleTypeObject, nil)

	// get the object keys
	var keys []string
	for key := range object {
		keys = append(keys, key)
	}

	// iterate over objects
	for _, key := range keys {
		// check the item type
		switch reflect.TypeOf(object[key]).String() {
		case mapInterfaceType:
			// another json object that we need to parse
			tmp, err := parseObject(object[key].(map[string]interface{}))
			if err != nil {
				return JsonObject{}, ErrJsonObject
			}
			tempJsonObject.items[key] = tmp
		case interfaceArrayType:
			// an array of interfaces
			tmp, err := parseArray(object[key].([]interface{}))
			if err != nil {
				return JsonObject{}, ErrArrayStructure
			}
			tempJsonObject.items[key] = newJsonObject("", jsonArrayType, nonSingleTypeObject, tmp)
		default:
			// any global value which we don't care about
			tempJsonObject.items[key] = newJsonObject(key, globalType, singleTypeObject, object[key])
		}
	}

	return tempJsonObject, nil
}

// parseArray gets a collection of interfaces and converts it to
// a collection of json objects.
func parseArray(obj []interface{}) ([]JsonObject, error) {
	// create a new collection of json objects
	var jsonObjectsArray []JsonObject

	// iterate over objects
	for _, item := range obj {
		// check to see if we need to parse a json object or not
		if reflect.TypeOf(item).String() == mapInterfaceType {
			tmp, err := parseObject(item.(map[string]interface{}))
			if err != nil {
				return nil, err
			}
			jsonObjectsArray = append(jsonObjectsArray, tmp)
		} else {
			tmp := newJsonObject("", globalType, nonSingleTypeObject, item)
			jsonObjectsArray = append(jsonObjectsArray, tmp)
		}
	}

	return jsonObjectsArray, nil
}
