package internal

// JsonObject is a single json structure.
type JsonObject struct {
	// object key and type
	key        string
	objectType string
	// object items which is a map
	items map[string]jsonValue
}

// jsonValue stores the value of each json object.
type jsonValue struct {
	// value type
	valueType string
	// value which is generic
	value interface{}
}
