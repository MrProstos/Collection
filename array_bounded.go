package collection

func (collect *Array[Key, Item, BaseCollection]) First() (Item, bool) {
	return collect.Get(Key(0))
}

func (collect *Array[Key, Item, BaseCollection]) Last() (Item, bool) {
	return collect.Get(Key(collect.Len() - 1))
}
