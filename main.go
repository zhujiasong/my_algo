package main

import (
	"fmt"
	"sort"
)

func main() {

	keys := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	index := sort.Search(len(keys), func(i int) bool {
		return keys[i] >= 4
	})

	fmt.Println(index)
}
