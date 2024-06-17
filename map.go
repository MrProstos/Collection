package collection

var _ Collection[string, any, map[string]any, *Map[string, any, map[string]any]] = &Map[string, any, map[string]any]{}

type Map[Key string, Item any, BaseCollection ~map[Key]Item] map[Key]Item

func NewMap[Item any]() *Map[string, Item, map[string]Item] {
	return &Map[string, Item, map[string]Item]{}
}

func NewMapFrom[Item any](data map[string]Item) *Map[string, Item, map[string]Item] {
	var collect Map[string, Item, map[string]Item] = data
	return &collect
}

func (collect *Map[Key, Item, BaseCollection]) Base() *Map[Key, Item, BaseCollection] {
	return collect
}

func (collect *Map[Key, Item, BaseCollection]) All() BaseCollection {
	return BaseCollection(*collect)
}

func (collect *Map[Key, Item, BaseCollection]) Len() int {
	return len(*collect)
}

func (collect *Map[Key, Item, BaseCollection]) Clone() Collection[Key, Item, BaseCollection, *Map[Key, Item, BaseCollection]] {
	cloneCollect := Map[Key, Item, BaseCollection]{}
	collect.Each(func(key Key, item Item) {
		cloneCollect[key] = item
	})
	return &cloneCollect
}

func (collect *Map[Key, Item, BaseCollection]) Has(key Key) bool {
	_, ok := (*collect)[key]
	return ok
}

func (collect *Map[Key, Item, BaseCollection]) Get(key Key) (Item, bool) {
	value, ok := (*collect)[key]
	return value, ok
}

func (collect *Map[Key, Item, BaseCollection]) Set(key Key, value Item) {
	(*collect)[key] = value
}

func (collect *Map[Key, Item, BaseCollection]) Delete(key Key) {
	delete(*collect, key)
}

func (collect *Map[Key, Item, BaseCollection]) Each(closure func(key Key, item Item)) Collection[Key, Item, BaseCollection, *Map[Key, Item, BaseCollection]] {
	for key, value := range *collect {
		closure(key, value)
	}
	return collect
}

func (collect *Map[Key, Item, BaseCollection]) Map(closure func(key Key, item Item) (Item, bool)) Collection[Key, Item, BaseCollection, *Map[Key, Item, BaseCollection]] {
	var newCollect Collection[Key, Item, BaseCollection, *Map[Key, Item, BaseCollection]] = &Map[Key, Item, BaseCollection]{}
	collect.Each(func(key Key, item Item) {
		if value, ok := closure(key, item); ok {
			newCollect.Set(key, value)
		}
	})
	return newCollect
}

func (collect *Map[Key, Item, BaseCollection]) Filter(closure func(key Key, item Item) bool) Collection[Key, Item, BaseCollection, *Map[Key, Item, BaseCollection]] {
	return collect.Map(func(key Key, item Item) (Item, bool) {
		return item, closure(key, item)
	})
}
