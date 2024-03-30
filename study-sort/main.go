package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []int{1, 7, 5}
	sort.Ints(s)
	fmt.Println(s)

	nums := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	for i := 0; i < len(nums)-1; i++ {
		for j := 0; j < len(nums)-1-i; j++ {
			if nums[j] > nums[j+1] {
				temp := nums[j]
				nums[j] = nums[j+1]
				nums[j+1] = temp
			}
		}
	}
	fmt.Println(nums)
}
