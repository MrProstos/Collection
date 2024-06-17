package collection

func (collect *Array[Key, Item, BaseCollection]) Each(closure func(key Key, item Item)) Collection[Key, Item, BaseCollection, *Array[Key, Item, BaseCollection]] {
	for key, item := range *collect {
		closure(Key(key), item)
	}
	return collect
}

func (collect *Array[Key, Item, BaseCollection]) Map(closure func(key Key, item Item) (Item, bool)) Collection[Key, Item, BaseCollection, *Array[Key, Item, BaseCollection]] {
	var newCollect Collection[Key, Item, BaseCollection, *Array[Key, Item, BaseCollection]] = &Array[Key, Item, BaseCollection]{}
	collect.Each(func(key Key, item Item) {
		if value, ok := closure(key, item); ok {
			newCollect.Base().Push(value)
		}
	})
	return newCollect
}

func (collect *Array[Key, Item, BaseCollection]) Filter(closure func(key Key, item Item) bool) Collection[Key, Item, BaseCollection, *Array[Key, Item, BaseCollection]] {
	return collect.Map(func(key Key, item Item) (Item, bool) {
		return item, closure(key, item)
	})
}
