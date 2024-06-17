package collection

import (
	"reflect"
	"sort"
)

func (collect *Array[Key, Item, BaseCollection]) Sort(sortType SortType) Collection[Key, Item, BaseCollection, *Array[Key, Item, BaseCollection]] {
	var newCollect = *collect.Clone().Base()
	sort.Slice(newCollect, func(i, j int) bool { return compare(newCollect[i], newCollect[j], "", sortType) })
	return &newCollect
}

func (collect *Array[Key, Item, BaseCollection]) SortWithKey(key string, sortType SortType) Collection[Key, Item, BaseCollection, *Array[Key, Item, BaseCollection]] {
	mustKey(key)
	var newCollect = collect.Clone().Base()
	sort.Slice(*newCollect, func(i, j int) bool { return compare((*newCollect)[i], (*newCollect)[j], key, sortType) })
	return newCollect
}

type SortType string

const (
	Ascending  SortType = "asc"
	Descending SortType = "desc"
)

type comparableCollection interface {
	less(a, b reflect.Value) bool
	greater(a, b reflect.Value) bool
}

type numericComparableCollection struct {
}

func (s numericComparableCollection) less(a, b reflect.Value) bool {
	return toFloat(a) < toFloat(b)
}

func (s numericComparableCollection) greater(a, b reflect.Value) bool {
	return toFloat(a) > toFloat(b)
}

type structComparableCollection struct {
	sortType SortType
	key      string
}

func (s structComparableCollection) less(a, b reflect.Value) bool {
	return toFloat(a.FieldByName(s.key)) < toFloat(b.FieldByName(s.key))
}

func (s structComparableCollection) greater(a, b reflect.Value) bool {
	return toFloat(a.FieldByName(s.key)) > toFloat(b.FieldByName(s.key))
}

func compare(a, b any, key string, sortType SortType) bool {
	aReflect := reflect.ValueOf(a)
	bReflect := reflect.ValueOf(b)

	if aReflect.Kind() == reflect.Ptr {
		aReflect = reflect.Indirect(aReflect)
	}
	if bReflect.Kind() == reflect.Ptr {
		bReflect = reflect.Indirect(bReflect)
	}

	var comparableCollection comparableCollection
	if key == "" {
		comparableCollection = numericComparableCollection{}
	} else {
		comparableCollection = structComparableCollection{key: key}
	}

	switch sortType {
	case Ascending:
		return comparableCollection.less(aReflect, bReflect)
	case Descending:
		return comparableCollection.greater(aReflect, bReflect)
	default:
		panic("collection: unknown sort type")
	}
}
