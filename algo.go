package main

import (
	"fmt"
	"log"
)

/*
*
给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。
示例 1：
输入：nums = [1,2,3]
输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
*/
//func all(ints [3]int) [][]int {
//	res := make([][]int, len(ints))
//	track :
//	return res
//}
//
//func sunFunc() {
//
//}

func FindSubSize(ints []int) int {
	fast := 0
	slow := 0
	for i, _ := range ints {
		if ints[fast] > ints[slow] {
			slow++
			ints[slow] = ints[fast]
			log.Print(ints[i])
		}
		fast++
	}
	return slow
}

func Test() {
	var ints = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	size := FindSubSize(ints)
	fmt.Println(ints)
	fmt.Println(size)
}
