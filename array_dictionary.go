package collection

func (collect *Array[Key, Item, BaseCollection]) Has(key Key) bool {
	return key >= 0 && key < Key(collect.Len())
}

func (collect *Array[Key, Item, BaseCollection]) Get(key Key) (Item, bool) {
	var (
		result Item
		exists bool
	)
	if collect.Has(key) {
		result = (*collect)[key]
		exists = true
	}
	return result, exists
}

func (collect *Array[Key, Item, BaseCollection]) Set(key Key, value Item) {
	(*collect)[key] = value
}

func (collect *Array[Key, Item, BaseCollection]) Delete(key Key) {
	if collect.Has(key) {
		*collect = append((*collect)[:key], (*collect)[key+1:]...)
	}
}
