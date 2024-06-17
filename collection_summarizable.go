package collection

type Summarizable[Key int | string, Item any, BaseCollection any] interface {
	Avg() float64
	AvgWithKey(key string) float64
	Sum() float64
	SumWithKey(key string) float64
	Max() float64
	MaxWithKey(key string) float64
	Min() float64
	MinWithKey(key string) float64
}
