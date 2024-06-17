package collection

type Iterable[Key int | string, Item any, BaseCollection any, BaseType any] interface {
	Each(closure func(key Key, item Item)) Collection[Key, Item, BaseCollection, BaseType]
	Map(closure func(key Key, item Item) (Item, bool)) Collection[Key, Item, BaseCollection, BaseType]
	Filter(closure func(key Key, item Item) bool) Collection[Key, Item, BaseCollection, BaseType]
}
