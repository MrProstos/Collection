package collection

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
