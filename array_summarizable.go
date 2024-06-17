package collection

import (
	"reflect"
)

func (collect *Array[Key, Item, BaseCollection]) Avg() float64 {
	return collect.avg("")
}

func (collect *Array[Key, Item, BaseCollection]) AvgWithKey(key string) float64 {
	mustKey(key)
	return collect.avg(key)
}

func (collect *Array[Key, Item, BaseCollection]) avg(key string) float64 {
	if collect.Len() > 0 {
		return collect.sum(key) / float64(collect.Len())
	} else {
		return 0
	}
}

func (collect *Array[Key, Item, BaseCollection]) Sum() float64 {
	return collect.sum("")
}

func (collect *Array[Key, Item, BaseCollection]) SumWithKey(key string) float64 {
	mustKey(key)
	return collect.sum(key)
}

func (collect *Array[Key, Item, BaseCollection]) sum(key string) float64 {
	var sum float64
	for _, el := range *collect {
		sum += getSummarizableValueWithKey(key, reflect.ValueOf(el))
	}
	return sum
}

func (collect *Array[Key, Item, BaseCollection]) Max() float64 {
	return collect.max("")
}

func (collect *Array[Key, Item, BaseCollection]) MaxWithKey(key string) float64 {
	mustKey(key)
	return collect.max(key)
}

func (collect *Array[Key, Item, BaseCollection]) max(key string) float64 {
	var maxValue float64
	collect.Each(func(_ Key, item Item) {
		value := getSummarizableValueWithKey(key, reflect.ValueOf(item))
		if value > maxValue {
			maxValue = value
		}
	})
	return maxValue
}

func (collect *Array[Key, Item, BaseCollection]) Min() float64 {
	return collect.min("")
}

func (collect *Array[Key, Item, BaseCollection]) MinWithKey(key string) float64 {
	mustKey(key)
	return collect.min(key)
}

func (collect *Array[Key, Item, BaseCollection]) min(key string) float64 {
	var (
		minValue  float64
		firstIter = true
	)
	collect.Each(func(_ Key, item Item) {
		if minValue == 0 && firstIter {
			minValue = getSummarizableValueWithKey(key, reflect.ValueOf(item))
			firstIter = false
		}

		value := getSummarizableValueWithKey(key, reflect.ValueOf(item))
		if value < minValue {
			minValue = value
		}
	})
	return minValue
}

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
