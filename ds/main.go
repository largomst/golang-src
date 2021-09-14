package main

import (
	"ds/linkedlist"
	"fmt"
)

func main() {
	obj := linkedlist.Constructor()
	for i := 0; i < 100; i++ {
		// obj.AddAtHead(i)
		// obj.AddAtTail(i)
		obj.AddAtIndex(0, i)
	}

	for {
		val := obj.Get(0)
		if val != -1 {
			fmt.Println(val)
			obj.DeleteAtIndex(0)
		} else {
			break
		}
	}

	// for i := 99; i >= 0; i-- {
	// 	val := obj.Get(i)
	// 	fmt.Println(val)
	// 	obj.DeleteAtIndex(i)
	// }
}
