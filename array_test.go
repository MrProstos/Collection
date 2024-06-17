package collection

import (
	"reflect"
	"testing"
)

type TestStruct struct {
	Id     uint
	Name   string
	Amount int
}

var (
	arrayTestInt = []int{1, 2, 3, 4, 5}
)

func TestArray_Base(t *testing.T) {
	expectedArray := NewArrayFrom(arrayTestInt)
	actualArrayInt := NewArrayFrom(arrayTestInt).Base()
	if !reflect.DeepEqual(expectedArray, actualArrayInt) {
		t.Fatal("ArrayBase fail. Expected:", expectedArray, ", actual: ", actualArrayInt)
	}
}

func TestArray_All(t *testing.T) {
	expected := []string{"[]int", "[]collection.TestStruct"}
	actual := []string{
		reflect.ValueOf(NewArrayFrom[int](arrayTestInt).All()).Type().String(),
		reflect.ValueOf(NewArrayFrom[TestStruct]([]TestStruct{{}, {}}).All()).Type().String(),
	}

	for index := range expected {
		if expected[index] != actual[index] {
			t.Fatal("ArrayAll fail. Expected ", expected[index], ", actual: ", actual[index])
		}
	}
}

func TestArray_Len(t *testing.T) {
	expected := 5
	actual := NewArrayFrom([]int{1, 2, 3, 4, 5}).Len()
	if actual != expected {
		t.Fatal("ArrayLen fail. Expected: ", expected, ", actual: ", actual)
	}
}

func TestArray_Has(t *testing.T) {
	actual := NewArrayFrom([]int{1, 2, 3, 4, 5}).Has(3)
	if !actual {
		t.Fatal("ArrayHas fail. Expected: ", true, ", actual: ", false)
	}

	actual = NewArrayFrom([]int{1, 2, 3, 4, 5}).Has(10)
	if actual {
		t.Fatal("ArrayHas fail. Expected: ", false, ", actual: ", true)
	}
}

func TestArray_Get(t *testing.T) {
	expectedInt := 1
	actualInt, _ := NewArrayFrom([]int{1, 2, 3, 4, 5}).Get(0)
	if expectedInt != actualInt {
		t.Fatal("ArrayGet fail. Expected: ", expectedInt, ", actual: ", actualInt)
	}

	expectedInt = 0
	actualInt, _ = NewArrayFrom([]int{1, 2, 3, 4, 5}).Get(10)
	if expectedInt != actualInt {
		t.Fatal("ArrayGet fail. Expected: ", expectedInt, ", actual: ", actualInt)
	}
}

func TestArray_Set(t *testing.T) {
	expected := NewArrayFrom([]int{1, 2, 3, 4, 5})
	actual := NewArrayFrom([]int{1, 2, 3, 4, 0})
	actual.Set(4, 5)
	if !reflect.DeepEqual(expected, actual) {
		t.Fatal("ArraySet fail. Expected: ", expected, ", actual: ", actual)
	}
}

func TestArray_Delete(t *testing.T) {
	expected := NewArrayFrom([]int{1, 2, 3, 4, 5})
	actual := NewArrayFrom([]int{1, 2, 3, 4, 5, 6})
	actual.Delete(5)
	if !reflect.DeepEqual(expected, actual) {
		t.Fatal("ArrayDelete fail. Expected: ", expected, ", actual: ", actual)
	}
}

func TestArray_Push(t *testing.T) {
	expected := NewArrayFrom([]int{1, 2, 3, 4, 5})
	actual := NewArrayFrom([]int{1, 2, 3})
	actual.Push(4, 5)
	if !reflect.DeepEqual(expected, actual) {
		t.Fatal("ArrayPush fail. Expected: ", expected, ", actual: ", actual)
	}
}

func TestArray_Prepend(t *testing.T) {
	expected := NewArrayFrom([]int{1, 2, 3, 4, 5})
	actual := NewArrayFrom([]int{3, 4, 5})
	actual.Prepend(1, 2)
	if !reflect.DeepEqual(expected, actual) {
		t.Fatal("ArrayPrepend fail. Expected: ", expected, ", actual: ", actual)
	}
}

func TestArray_Pop(t *testing.T) {
	expected := NewArrayFrom([]int{1, 2, 3, 4, 5})
	actual := NewArrayFrom([]int{1, 2, 3, 4, 5, 6})
	last := actual.Pop()
	if !reflect.DeepEqual(expected, actual) || last != 6 {
		t.Fatal("ArrayPop fail. Expected: ", expected, ", actual: ", actual, ", last: ", last)
	}
}

func TestArray_Shift(t *testing.T) {
	expected := NewArrayFrom([]int{2, 3, 4, 5})
	actual := NewArrayFrom([]int{1, 2, 3, 4, 5})
	first := actual.Shift()
	if !reflect.DeepEqual(expected, actual) || first != 1 {
		t.Fatal("ArrayPop fail. Expected: ", expected, ", actual: ", actual, ", first: ", first)
	}
}

func TestArray_Each(t *testing.T) {
	expected := NewArrayFrom([]int{1, 2, 3, 4, 5})
	actual := NewArrayFrom([]int{1, 2, 3, 4, 5})
	actual.Each(func(index int, _ int) {
		if (*expected)[index] != (*actual)[index] {
			t.Fatal("ArrayEach fail. Expected: ", (*expected)[index], ", actual: ", (*actual)[index])
		}
	})
}

func TestArray_Map(t *testing.T) {
	expected := NewArrayFrom([]int{1, 1, 1, 1, 1})
	actualOrigin := NewArrayFrom([]int{1, 2, 3, 4, 5})
	actualChanged := actualOrigin.Map(func(key int, item int) (int, bool) {
		return 1, true
	})

	if reflect.DeepEqual(expected, actualOrigin) {
		t.Fatal("ArrayMap fail. Expected: ", expected, ", actualOrigin: ", actualOrigin)
	}

	if !reflect.DeepEqual(expected, actualChanged) {
		t.Fatal("ArrayMap fail. Expected: ", expected, ", actualOrigin: ", actualChanged)
	}
}

func TestArray_Filter(t *testing.T) {
	expected := NewArrayFrom([]int{1, 2})
	actualOrigin := NewArrayFrom([]int{1, 2, 3, 4, 5, 9})
	actualChanged := actualOrigin.Filter(func(_ int, item int) bool { return item < 3 })

	if reflect.DeepEqual(expected, actualOrigin) {
		t.Fatal("ArrayFilter fail. Expected: ", expected, ", actualOrigin: ", actualOrigin)
	}

	if !reflect.DeepEqual(expected, actualChanged) {
		t.Fatal("ArrayFilter fail. Expected: ", expected, ", actualOrigin: ", actualChanged)
	}
}

func TestArray_First(t *testing.T) {
	arr := NewArrayFrom([]int{1, 2})

	expected := 1
	actual, _ := arr.First()
	if expected != actual {
		t.Fatal("ArrayFirst fail. Expected: ", expected, ", actual: ", actual)
	}

	arr = NewArrayFrom([]int{})
	_, exists := arr.First()
	if exists {
		t.Fatal("ArrayFirst fail. Expected: ", false, ", actual: ", true)
	}
}

func TestArray_Last(t *testing.T) {
	arr := NewArrayFrom([]int{1, 2})

	expected := 2
	actual, _ := arr.Last()
	if expected != actual {
		t.Fatal("ArrayLast fail. Expected: ", expected, ", actual: ", actual)
	}

	arr = NewArrayFrom([]int{})
	_, exists := arr.Last()
	if exists {
		t.Fatal("ArrayLast fail. Expected: ", false, ", actual: ", true)
	}
}

func TestArray_Sort(t *testing.T) {
	expectedArrs := [][]int{
		{1, 2, 3, 4, 5},
		{5, 4, 3, 2, 1},
	}
	actualArrs := [][]int{
		{5, 4, 3, 2, 1},
		{1, 2, 3, 4, 5},
	}
	sortTypes := []SortType{Ascending, Descending}

	for index := range actualArrs {
		expected := NewArrayFrom(expectedArrs[index])
		actualOrigin := NewArrayFrom(actualArrs[index])
		actualChanged := actualOrigin.Sort(sortTypes[index])

		if reflect.DeepEqual(expected, actualOrigin) {
			t.Fatal("ArrayFilter fail. Expected: ", expected, ", actualOrigin: ", actualOrigin)
		}

		if !reflect.DeepEqual(expected, actualChanged) {
			t.Fatal("ArrayFilter fail. Expected: ", expected, ", actualOrigin: ", actualChanged)
		}
	}
}

func TestArray_SortWithKey(t *testing.T) {
	expectedArrs := [][]TestStruct{
		{{Id: 1}, {Id: 2}, {Id: 3}},
		{{Id: 3}, {Id: 2}, {Id: 1}},
	}
	actualArrs := [][]TestStruct{
		{{Id: 3}, {Id: 2}, {Id: 1}},
		{{Id: 1}, {Id: 2}, {Id: 3}},
	}
	sortTypes := []SortType{Ascending, Descending}

	for index := range actualArrs {
		expected := NewArrayFrom(expectedArrs[index])
		actualOrigin := NewArrayFrom(actualArrs[index])
		actualChanged := actualOrigin.SortWithKey("Id", sortTypes[index])

		if reflect.DeepEqual(expected, actualOrigin) {
			t.Fatal("ArrayFilter fail. Expected: ", expected, ", actualOrigin: ", actualOrigin)
		}

		if !reflect.DeepEqual(expected, actualChanged) {
			t.Fatal("ArrayFilter fail. Expected: ", expected, ", actualOrigin: ", actualChanged)
		}
	}
}
