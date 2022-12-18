package internal

import (
	"encoding/json"
	"reflect"
)

// Parse method gets an array of bytes and returns
// a json object that is one n-d array.
func Parse(bytes []byte) (interface{}, error) {
	var (
		jObj map[string]interface{}
		jArr []map[string]interface{}
	)

	// unmarshalling json into variables
	if err := json.Unmarshal(bytes, &jObj); err != nil {
		if err = json.Unmarshal(bytes, &jArr); err != nil {
			return nil, ErrJsonObject
		}
	}

	// check the variables
	if jObj != nil {
		return parseObj(jObj)
	} else if jArr != nil {
		return parseArray(jArr)
	} else {
		return nil, ErrJsonObject
	}
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
			tmp, err := parseArray(obj[key].([]map[string]interface{}))
			if err != nil {
				return nil, err
			}

			obj[key] = tmp
		}
	}

	return obj, nil
}

func parseArray(obj []map[string]interface{}) ([]map[string]interface{}, error) {
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
