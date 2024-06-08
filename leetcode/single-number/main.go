package main

import "fmt"

func main() {
	fmt.Println(singleNumber([]int{4, 1, 2, 1, 2}))
}

// func singleNumber(nums []int) int {
// 	var store map[int]int = make(map[int]int)
// 	for _, v := range nums {
// 		if _, ok := store[v]; ok {
// 			store[v] += 1
// 		} else {
// 			store[v] = 1
// 		}
// 	}
// 	answer := 0
// 	for i := range store {
// 		if store[i] == 1 {
// 			answer = i
// 		}
// 	}
// 	return answer
// }

func singleNumber(nums []int) int {
	result := nums[0]
	for i := 1; i < len(nums); i++ {
		result ^= nums[i]
		continue
	}
	return result
}
