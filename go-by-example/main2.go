package main

import (
	"fmt"
	"sort"
)

func main() {
	// 排序
	sorts()

	// 自定义排序
	bylength()
}

func sorts() {
	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println("Strings", strs)

	ints := []int{7, 3, 4, 1}
	sort.Ints(ints)
	fmt.Println("Ints", ints)

	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted:", s)
}

type byLength []string

func bylength() {
	fruits := []string{"peach", "banana", "kiwi"}
	sort.Sort(byLength(fruits))
	fmt.Println(fruits)
}
func (s byLength) Len() int {
	return len(s)
}
func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}
