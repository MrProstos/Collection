package collection

type Collection[Key int | string, Item any, BaseCollection any, BaseType any] interface {
	Base() BaseType
	All() BaseCollection
	Len() int
	Clone() Collection[Key, Item, BaseCollection, BaseType]

	Dictionary[Key, Item, BaseCollection, BaseType]
	Iterable[Key, Item, BaseCollection, BaseType]
}
