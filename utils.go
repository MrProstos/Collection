package collection

import (
	"reflect"
)

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
