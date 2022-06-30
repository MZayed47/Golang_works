package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {

	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if i == j {
				continue
			} else {
				if nums[i]+nums[j] == target {
					fmt.Println()
					return []int{i, j}
				}
			}
		}
	}
	return []int{}

}
