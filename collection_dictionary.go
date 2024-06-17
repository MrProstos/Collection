package collection

type Dictionary[Key int | string, Item any, BaseCollection any, BaseType any] interface {
	Has(key Key) bool
	Get(key Key) (Item, bool)
	Set(key Key, value Item)
	Delete(key Key)
}
