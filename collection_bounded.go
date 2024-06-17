package collection

type Bounded[Key int | string, Item any, BaseCollection any] interface {
	First() (Item, bool)
	Last() (Item, bool)
}
