package main

import (
	"fmt"
	"sort"
)

func main() {
	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println(strs)

	nums := []int{22, 1, 4, 645, 3}
	sort.Ints(nums)
	fmt.Println(nums)

	fmt.Println(sort.IntsAreSorted(nums))
}
