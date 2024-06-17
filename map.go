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
