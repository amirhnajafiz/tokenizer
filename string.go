package explorer

import (
	"fmt"
	"reflect"
)

// generateTabsString puts tabs next to string.
func generateTabsString(number int) string {
	tmp := ""

	for i := 0; i < number; i++ {
		tmp = tmp + " "
	}

	return tmp
}

// renderSingleItem render a single key value set.
func renderSingleItem(key string, value interface{}, numberOfTabs int, flag ...bool) string {
	// set the value
	tmp := fmt.Sprintf("\"%v\"", value)

	// check the flags
	if len(flag) > 0 {
		if flag[0] {
			tmp = fmt.Sprintf("\"%v\"", reflect.TypeOf(value))
		}
	}

	// check for key existence
	if key != "" {
		tmp = fmt.Sprintf("\"%s\": %s", key, tmp)
	}

	return fmt.Sprintf("%s%s", generateTabsString(numberOfTabs), tmp)
}

// renderObject converts a json object to string.
func renderObject(object JsonObject, numberOfTabs, tabUnit, control int, schema bool) string {
	var (
		// tmp stores the result of printing
		tmp string
		// get tabs
		tabsNow  = generateTabsString(numberOfTabs)
		tabsNext = generateTabsString(numberOfTabs + tabUnit)
		// get then number of items
		numberOfItems = len(object.items)
	)

	// iterate over items
	index := 1
	for key := range object.items {
		obj := object.items[key]

		if obj.valueType == globalType {
			// get the value based on schema or not
			var value interface{}
			if schema {
				value = obj.Type()
			} else {
				value = obj.Value()
			}

			tmp = fmt.Sprintf("%s\n%s", tmp, renderSingleItem(key, value, numberOfTabs+tabUnit))
		} else if obj.valueType == jsonObjectType {
			tmp = fmt.Sprintf(
				"%s\n%s\"%s\": %s",
				tmp,
				tabsNext,
				key,
				renderObject(obj, numberOfTabs+tabUnit, tabUnit, innerObject, schema),
			)
		} else if obj.valueType == jsonArrayType {
			temp := ""
			numb := len(obj.value.([]JsonObject))

			in := 1
			for _, innerObj := range obj.value.([]JsonObject) {
				temp = fmt.Sprintf(
					"%s\n%s",
					temp,
					renderObject(innerObj, numberOfTabs+2*tabUnit, tabUnit, arrayObject, schema),
				)

				// add ',' to the end
				if in != numb {
					temp = temp + ","
				}

				in++
			}

			tmp = fmt.Sprintf("%s\n%s\"%s\": [%s%s\n%s]",
				tmp,
				tabsNext,
				key,
				tabsNext,
				temp,
				tabsNext,
			)
		} else if obj.valueType == globalArrayType {
			temp := ""
			numb := len(obj.value.([]interface{}))

			in := 1
			for _, innerObj := range obj.value.([]interface{}) {
				temp = fmt.Sprintf(
					"%s\n%s",
					temp,
					renderSingleItem("", innerObj, numberOfTabs+2*tabUnit, schema),
				)

				// add ',' to the end
				if in != numb {
					temp = temp + ","
				}

				in++
			}

			tmp = fmt.Sprintf("%s\n%s\"%s\": [%s%s\n%s]",
				tmp,
				tabsNext,
				key,
				tabsNext,
				temp,
				tabsNext,
			)
		}

		// add ',' to the end
		if index != numberOfItems {
			tmp = tmp + ","
		}

		// increase index
		index++
	}

	// things are different for first object and inner ones
	switch control {
	case baseObject:
		return fmt.Sprintf("{%s\n}", tmp)
	case innerObject:
		return fmt.Sprintf("{%s\n%s}", tmp, tabsNow)
	case arrayObject:
		return fmt.Sprintf("%s{%s\n%s}", tabsNow, tmp, tabsNow)
	default:
		return ""
	}
}

// Pretty returns a pretty json string.
func (j JsonObject) Pretty(space int) string {
	return renderObject(j, space, space, baseObject, false)
}

// Schema returns json object keys and types.
func (j JsonObject) Schema() string {
	return renderObject(j, 2, 2, baseObject, true)
}
