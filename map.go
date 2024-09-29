package main

import (
	"fmt"
	"sort"
)

func TestMap() {

	m := map[string]int{"a": 1, "b": 2, "c": 3}
	fmt.Println(m)

	for k, v := range m {
		fmt.Println(k, v)
	}

	var sl []int

	for k := range m {
		sl = append(sl, m[k])
	}

	sort.Ints(sl)

	for _, v := range sl {
		fmt.Println(v, m)
	}
}
