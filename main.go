package main

import (
	"fmt"

	"github.com/diegovaltrick/pkg/ds"
)

type Foo struct {
	F int
}

func main() {
	arr := ds.NewArray[Foo]()
	arr.Push(Foo{1})
	arr.Push(Foo{13})
	arr.Push(Foo{7})
	arr.Push(Foo{9})
	arr.Push(Foo{25})
	fmt.Println(arr.At(3))

	arr.ForEach(func(item Foo, idx int) {
		fmt.Printf("%d: %d\n", idx, item.F)
	})

	newArr := arr.Map(func(item Foo, idx int) bool {
		return item.F > 10
	})

	fmt.Println(newArr)

	fmt.Println(arr.Pop())
	fmt.Println(arr)
	fmt.Println(arr.Shift())
	fmt.Println(arr)
}