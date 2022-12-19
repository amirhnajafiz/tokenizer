package internal

import (
	"fmt"
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
func renderSingleItem(key string, value interface{}, numberOfTabs int) string {
	// set the value
	tmp := fmt.Sprintf("\"%v\"", value)

	// check for key existence
	if key != "" {
		tmp = fmt.Sprintf("\"%s\": %s", key, tmp)
	}

	return fmt.Sprintf("%s%s", generateTabsString(numberOfTabs), tmp)
}

// renderObject converts a json object to string.
func renderObject(object JsonObject, numberOfTabs, tabUnit, control int) string {
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
			tmp = fmt.Sprintf("%s\n%s", tmp, renderSingleItem(key, obj.Value(), numberOfTabs+tabUnit))
		} else if obj.valueType == jsonObjectType {
			tmp = fmt.Sprintf(
				"%s\n%s\"%s\": %s",
				tmp,
				tabsNext,
				key,
				renderObject(obj, numberOfTabs+tabUnit, tabUnit, innerObject),
			)
		} else if obj.valueType == jsonArrayType {
			temp := ""
			numb := len(obj.value.([]JsonObject))

			in := 1
			for _, innerObj := range obj.value.([]JsonObject) {
				temp = fmt.Sprintf(
					"%s\n%s",
					temp,
					renderObject(innerObj, numberOfTabs+2*tabUnit, tabUnit, arrayObject),
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
					renderSingleItem("", innerObj, numberOfTabs+2*tabUnit),
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

func (j JsonObject) Pretty(space int) string {
	return renderObject(j, space, space, baseObject)
}
