package constants

const (
	CONFIG_FILE_NAME        = "type-utils.yaml"
	DEFAULT_SCHEMA_DIR_NAME = "schemas"
)

const (
	Q = "$(if $(filter 1,$(V)),,@)"
	M = "\033[36;1m>\033[0m"
)

const (
	// Basic types
	TypeString  = "string"
	TypeInt     = "int"
	TypeInt32   = "int32"
	TypeInt64   = "int64"
	TypeFloat   = "float"
	TypeDouble  = "double"
	TypeBoolean = "boolean"

	// Date and time types
	TypeDate      = "date"
	TypeTime      = "time"
	TypeDateTime  = "datetime"
	TypeTimestamp = "timestamp"
)
