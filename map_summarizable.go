package collection

import "reflect"

func (collect *Map[Key, Item, BaseCollection]) Avg() float64 {
	return collect.avg("")
}

func (collect *Map[Key, Item, BaseCollection]) AvgWithKey(key string) float64 {
	mustKey(key)
	return collect.avg(key)
}

func (collect *Map[Key, Item, BaseCollection]) avg(key string) float64 {
	if collect.Len() > 0 {
		return collect.sum(key) / float64(collect.Len())
	} else {
		return 0
	}
}

func (collect *Map[Key, Item, BaseCollection]) Sum() float64 {
	return collect.sum("")
}

func (collect *Map[Key, Item, BaseCollection]) SumWithKey(key string) float64 {
	mustKey(key)
	return collect.sum(key)
}

func (collect *Map[Key, Item, BaseCollection]) sum(key string) float64 {
	var sum float64
	for _, el := range *collect {
		sum += getSummarizableValueWithKey(key, reflect.ValueOf(el))
	}
	return sum
}

func (collect *Map[Key, Item, BaseCollection]) Max() float64 {
	return collect.max("")
}

func (collect *Map[Key, Item, BaseCollection]) MaxWithKey(key string) float64 {
	mustKey(key)
	return collect.max(key)
}

func (collect *Map[Key, Item, BaseCollection]) max(key string) float64 {
	var maxValue float64
	collect.Each(func(_ Key, item Item) {
		value := getSummarizableValueWithKey(key, reflect.ValueOf(item))
		if value > maxValue {
			maxValue = value
		}
	})
	return maxValue
}

func (collect *Map[Key, Item, BaseCollection]) Min() float64 {
	return collect.min("")
}

func (collect *Map[Key, Item, BaseCollection]) MinWithKey(key string) float64 {
	mustKey(key)
	return collect.min(key)
}

func (collect *Map[Key, Item, BaseCollection]) min(key string) float64 {
	var (
		minValue  float64
		firstIter = true
	)
	collect.Each(func(_ Key, item Item) {
		if firstIter {
			minValue = getSummarizableValueWithKey(key, reflect.ValueOf(item))
			firstIter = false
			return
		}

		value := getSummarizableValueWithKey(key, reflect.ValueOf(item))
		if value < minValue {
			minValue = value
		}
	})
	return minValue
}
