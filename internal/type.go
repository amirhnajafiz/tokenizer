package internal

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
