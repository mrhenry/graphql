package graphql

import (
	"fmt"
	"math"
	"strconv"
	"time"

	"github.com/graphql-go/graphql/language/ast"
)

// As per the GraphQL Spec, Integers are only treated as valid when a valid
// 32-bit signed integer, providing the broadest support across platforms.
//
// n.b. JavaScript's integers are safe between -(2^53 - 1) and 2^53 - 1 because
// they are internally represented as IEEE 754 doubles.
func coerceInt(value interface{}) interface{} {
	var (
		i int64
		r float64
	)

	switch value := value.(type) {

	case bool:
		if value == true {
			i = 1
		} else {
			i = 0
		}
	case *bool:
		if *value == true {
			i = 1
		} else {
			i = 0
		}

	case int:
		i = int64(value)
	case *int:
		i = int64(*value)

	case int8:
		i = int64(value)
	case *int8:
		i = int64(*value)

	case int16:
		i = int64(value)
	case *int16:
		i = int64(*value)

	case int32:
		i = int64(value)
	case *int32:
		i = int64(*value)

	case int64:
		i = int64(value)
	case *int64:
		i = int64(*value)

	case uint:
		r = float64(value)
		i = int64(value)
	case *uint:
		r = float64(*value)
		i = int64(*value)

	case uint8:
		i = int64(value)
	case *uint8:
		i = int64(*value)

	case uint16:
		i = int64(value)
	case *uint16:
		i = int64(*value)

	case uint32:
		i = int64(value)
	case *uint32:
		i = int64(*value)

	case uint64:
		r = float64(value)
		i = int64(value)
	case *uint64:
		r = float64(*value)
		i = int64(*value)

	case float32:
		r = float64(value)
		i = int64(value)
	case *float32:
		r = float64(*value)
		i = int64(*value)

	case float64:
		r = float64(value)
		i = int64(value)
	case *float64:
		r = float64(*value)
		i = int64(*value)

	case string:
		val, err := strconv.ParseFloat(value, 0)
		if err != nil {
			return nil
		}
		r = float64(val)
		i = int64(val)
	case *string:
		val, err := strconv.ParseFloat(*value, 0)
		if err != nil {
			return nil
		}
		r = float64(val)
		i = int64(val)

	default:
		// If the value cannot be transformed into an int, return nil instead of '0'
		// to denote 'no integer found'
		return nil
	}

	// If the value is out of range, return nil instead of '0'
	if r < math.MinInt32 || r > math.MaxInt32 {
		return nil
	}
	if i < math.MinInt32 || i > math.MaxInt32 {
		return nil
	}

	return int(i)
}

// Int is the GraphQL Integer type definition.
var Int = NewScalar(ScalarConfig{
	Name: "Int",
	Description: "The `Int` scalar type represents non-fractional signed whole numeric " +
		"values. Int can represent values between -(2^31) and 2^31 - 1. ",
	Serialize:  coerceInt,
	ParseValue: coerceInt,
	ParseLiteral: func(valueAST ast.Value) interface{} {
		switch valueAST := valueAST.(type) {
		case *ast.IntValue:
			if intValue, err := strconv.Atoi(valueAST.Value); err == nil {
				return intValue
			}
		}
		return nil
	},
})

func coerceFloat(value interface{}) interface{} {
	var f float64

	switch value := value.(type) {

	case bool:
		if value == true {
			f = 1.0
		} else {
			f = 0.0
		}
	case *bool:
		if *value == true {
			f = 1.0
		} else {
			f = 0.0
		}

	case int:
		f = float64(value)
	case *int:
		f = float64(*value)

	case int8:
		f = float64(value)
	case *int8:
		f = float64(*value)

	case int16:
		f = float64(value)
	case *int16:
		f = float64(*value)

	case int32:
		f = float64(value)
	case *int32:
		f = float64(*value)

	case int64:
		f = float64(value)
	case *int64:
		f = float64(*value)

	case uint:
		f = float64(value)
	case *uint:
		f = float64(*value)

	case uint8:
		f = float64(value)
	case *uint8:
		f = float64(*value)

	case uint16:
		f = float64(value)
	case *uint16:
		f = float64(*value)

	case uint32:
		f = float64(value)
	case *uint32:
		f = float64(*value)

	case uint64:
		f = float64(value)
	case *uint64:
		f = float64(*value)

		// retain precision
	case float32:
		return value
	case *float32:
		return *value

	case float64:
		f = float64(value)
	case *float64:
		f = float64(*value)

	case string:
		val, err := strconv.ParseFloat(value, 0)
		if err != nil {
			return nil
		} else {
			f = float64(val)
		}
	case *string:
		val, err := strconv.ParseFloat(*value, 0)
		if err != nil {
			return nil
		} else {
			f = float64(val)
		}

	}

	return f
}

// Float is the GraphQL float type definition.
var Float = NewScalar(ScalarConfig{
	Name: "Float",
	Description: "The `Float` scalar type represents signed double-precision fractional " +
		"values as specified by " +
		"[IEEE 754](http://en.wikipedia.org/wiki/IEEE_floating_point). ",
	Serialize:  coerceFloat,
	ParseValue: coerceFloat,
	ParseLiteral: func(valueAST ast.Value) interface{} {
		switch valueAST := valueAST.(type) {
		case *ast.FloatValue:
			if floatValue, err := strconv.ParseFloat(valueAST.Value, 32); err == nil {
				return floatValue
			}
		case *ast.IntValue:
			if floatValue, err := strconv.ParseFloat(valueAST.Value, 32); err == nil {
				return floatValue
			}
		}
		return nil
	},
})

func coerceString(value interface{}) interface{} {
	switch value := value.(type) {
	case string:
		return value
	case *string:
		return *value
	}

	if v, ok := value.(fmt.Stringer); ok {
		return v.String()
	}

	return fmt.Sprintf("%v", value)
}

// String is the GraphQL string type definition
var String = NewScalar(ScalarConfig{
	Name: "String",
	Description: "The `String` scalar type represents textual data, represented as UTF-8 " +
		"character sequences. The String type is most often used by GraphQL to " +
		"represent free-form human-readable text.",
	Serialize:  coerceString,
	ParseValue: coerceString,
	ParseLiteral: func(valueAST ast.Value) interface{} {
		switch valueAST := valueAST.(type) {
		case *ast.StringValue:
			return valueAST.Value
		}
		return nil
	},
})

func coerceBool(value interface{}) interface{} {
	switch value := value.(type) {
	case bool:
		return value
	case *bool:
		return *value
	case string:
		switch value {
		case "", "false":
			return false
		}
		return true
	case *string:
		return coerceBool(*value)
	case float64:
		if value != 0 {
			return true
		}
		return false
	case *float64:
		return coerceBool(*value)
	case float32:
		if value != 0 {
			return true
		}
		return false
	case *float32:
		return coerceBool(*value)
	case int:
		if value != 0 {
			return true
		}
		return false
	case *int:
		return coerceBool(*value)
	}
	return false
}

// Boolean is the GraphQL boolean type definition
var Boolean = NewScalar(ScalarConfig{
	Name:        "Boolean",
	Description: "The `Boolean` scalar type represents `true` or `false`.",
	Serialize:   coerceBool,
	ParseValue:  coerceBool,
	ParseLiteral: func(valueAST ast.Value) interface{} {
		switch valueAST := valueAST.(type) {
		case *ast.BooleanValue:
			return valueAST.Value
		}
		return nil
	},
})

// ID is the GraphQL id type definition
var ID = NewScalar(ScalarConfig{
	Name: "ID",
	Description: "The `ID` scalar type represents a unique identifier, often used to " +
		"refetch an object or as key for a cache. The ID type appears in a JSON " +
		"response as a String; however, it is not intended to be human-readable. " +
		"When expected as an input type, any string (such as `\"4\"`) or integer " +
		"(such as `4`) input value will be accepted as an ID.",
	Serialize:  coerceString,
	ParseValue: coerceString,
	ParseLiteral: func(valueAST ast.Value) interface{} {
		switch valueAST := valueAST.(type) {
		case *ast.IntValue:
			return valueAST.Value
		case *ast.StringValue:
			return valueAST.Value
		}
		return nil
	},
})

func serializeDateTime(value interface{}) interface{} {
	switch value := value.(type) {
	case time.Time:
		buff, err := value.MarshalText()
		if err != nil {
			return nil
		}

		return string(buff)
	case *time.Time:
		return serializeDateTime(*value)
	default:
		return nil
	}
}

func unserializeDateTime(value interface{}) interface{} {
	switch value := value.(type) {
	case []byte:
		t := time.Time{}
		err := t.UnmarshalText(value)
		if err != nil {
			return nil
		}

		return t
	case string:
		return unserializeDateTime([]byte(value))
	case *string:
		return unserializeDateTime([]byte(*value))
	default:
		return nil
	}
}

var DateTime = NewScalar(ScalarConfig{
	Name: "DateTime",
	Description: "The `DateTime` scalar type represents a DateTime." +
		" The DateTime is serialized as an RFC 3339 quoted string",
	Serialize:  serializeDateTime,
	ParseValue: unserializeDateTime,
	ParseLiteral: func(valueAST ast.Value) interface{} {
		switch valueAST := valueAST.(type) {
		case *ast.StringValue:
			return valueAST.Value
		}
		return nil
	},
})
