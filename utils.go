package collection

import (
	"reflect"
)

func getSummarizableValueWithKey(key string, value reflect.Value) float64 {
	if key == "" && value.Comparable() {
		return toFloat(value)
	}

	if key != "" && value.Kind() == reflect.Interface || value.Kind() == reflect.Struct {
		if structField := value.FieldByName(key); value.Comparable() {
			return toFloat(structField)
		}
	}

	return 0
}

func toFloat(value reflect.Value) float64 {
	switch true {
	case value.CanInt():
		return float64(value.Int())
	case value.CanUint():
		return float64(value.Uint())
	case value.CanFloat():
		return value.Float()
	default:
		panic("collection: value must be int, uint or float")
	}
}

func mustKey(key string) {
	if key == "" {
		panic("collection: must key")
	}
}
