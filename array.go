package collection

var _ Collection[int, any, []any, *Array[int, any, []any]] = &Array[int, any, []any]{}

type Array[Key int, Item any, BaseCollection ~[]Item] []Item

func NewArray[Item any]() *Array[int, Item, []Item] {
	return &Array[int, Item, []Item]{}
}

func NewArrayFrom[Item any](data []Item) *Array[int, Item, []Item] {
	return (*Array[int, Item, []Item])(&data)
}

func (collect *Array[Key, Item, BaseCollection]) Base() *Array[Key, Item, BaseCollection] {
	return collect
}

func (collect *Array[Key, Item, BaseCollection]) All() BaseCollection {
	return BaseCollection(*collect)
}

func (collect *Array[Key, Item, BaseCollection]) Len() int {
	return len(*collect)
}

func (collect *Array[Key, Item, BaseCollection]) Clone() Collection[Key, Item, BaseCollection, *Array[Key, Item, BaseCollection]] {
	cloneCollect := &Array[Key, Item, BaseCollection]{}
	cloneCollect.Push(*collect...)
	return cloneCollect
}

func (collect *Array[Key, Item, BaseCollection]) Push(items ...Item) {
	*collect = append(*collect, items...)
}

func (collect *Array[Key, Item, BaseCollection]) Prepend(items ...Item) {
	*collect = append(items, *collect...)
}

func (collect *Array[Key, Item, BaseCollection]) Pop() Item {
	var result Item
	lastIndex := collect.Len() - 1
	if collect.Has(Key(lastIndex)) {
		result = (*collect)[lastIndex]
		collect.Delete(Key(lastIndex))
	}
	return result
}

func (collect *Array[Key, Item, BaseCollection]) Shift() Item {
	var result Item
	firstIndex := Key(0)
	if collect.Has(firstIndex) {
		result = (*collect)[firstIndex]
		collect.Delete(firstIndex)
	}
	return result
}
