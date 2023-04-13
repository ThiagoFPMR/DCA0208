package main

import (
	"dca0208/list"
	"dca0208/search"
	"fmt"
)

func main() {
	var arr list.ArrayList
	arr.Init()

	limit := 10
	for i:=0; i < limit; i++ {
		arr.AddToBack(i)
	}

	fmt.Println(search.BinarySearch(arr, 5, 0, 10))
}

// Read this: https://stackoverflow.com/questions/40823315/x-does-not-implement-y-method-has-a-pointer-receiver