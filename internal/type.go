package internal

const (
	// object types enum
	singleTypeObject = iota + 1
	nonSingleTypeObject

	// value types enum
	globalType
	jsonObjectType
	jsonArrayType
)

const (
	// interface types that we use
	mapInterfaceType   = "map[string]interface {}"
	interfaceArrayType = "[]interface {}"
)
