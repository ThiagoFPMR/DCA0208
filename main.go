package main

import (
	"dca0208/list"
	"fmt"
)

func main() {

	var arraylist list.ArrayList
	arraylist.Init()
	limit := 15
	for i := 0; i < limit; i++ {
		arraylist.AddToBack(i)
	}
	for i := 0; i < limit; i++ {
		fmt.Println(arraylist.Get(i))
	}
	for i := 0; i < limit; i++ {
		if i%2 == 0 {
			arraylist.Update(i, i*2)
		}
	}
	fmt.Println("######################")
	for i := 0; i < limit; i++ {
		fmt.Println(arraylist.Get(i))
	}
}
