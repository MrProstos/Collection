package collection

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
