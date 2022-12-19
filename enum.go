package explorer

const (
	// value types enum
	globalType = iota + 1
	globalArrayType
	jsonObjectType
	jsonArrayType
)

const (
	// interface types that we use
	mapInterfaceType   = "map[string]interface {}"
	interfaceArrayType = "[]interface {}"
)

const (
	// generate string enums
	baseObject = iota + 1
	innerObject
	arrayObject
)
