package collection

type Sortable[Key int | string, Item any, BaseCollection any, BaseType any] interface {
	Sort() Collection[Key, Item, BaseCollection, BaseType]
	SortWithKey(key string) Collection[Key, Item, BaseCollection, BaseType]
}
