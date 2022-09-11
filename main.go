package main

import (
	"fmt"

	"github.com/diegovaltrick/pkg/ds"
)

func main() {
	arr := ds.NewArray[int]()
	arr.Push(1)
	arr.Push(2)
	arr.Push(3)
	arr.Push(36)
	arr.Push(42)
	fmt.Println(arr.At(3))

	arr.ForEach(func(item, idx int) {
		fmt.Printf("%d: %d\n", idx, item)
	})

	newArr := arr.Map(func(item, idx int) bool {
		return item > 10
	})

	fmt.Println(newArr)

	fmt.Println(arr.Pop())
	fmt.Println(arr)
	fmt.Println(arr.Shift())
	fmt.Println(arr)
	arr.Unshift(7)
	fmt.Println(arr)
}