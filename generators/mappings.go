package generators

const arrayType = "Array"

var protoToClickhouse = map[string]string{
	"double":   "Float64",
	"float":    "Float32",
	"int32":    "Int32",
	"int64":    "Int64",
	"uint32":   "UInt32",
	"uint64":   "UInt64",
	"sint32":   "Int32",
	"sint64":   "Int64",
	"fixed32":  "Int32",
	"fixed64":  "Int64",
	"sfixed32": "Int32",
	"sfixed64": "Int64",
	"string":   "String",
	"bool":     "UInt8",
	"bytes":    "String",
	"enum":     "UInt8",
}
