//Given a sorted array of distinct integers and a target value, return the index if the target is found. 
//If not, return the index where it would be if it were inserted in order.
//You must write an algorithm with O(log n) runtime complexity.

package main

import "fmt"

func searchInsert(nums []int, target int) int {
	if nums[0] > target {
		return 0
	}
	if nums[len(nums)-1] < target {
		return len(nums)
	}
	return search(nums, target, 0, len(nums))
}

func search(nums []int, target int, start int, end int) int {
	if start == end {
		return start
	}
	middleIndex := (start + end) / 2
	if nums[middleIndex] == target {
		return middleIndex
	} else if nums[middleIndex] < target {
		return search(nums, target, middleIndex+1, end)
	} else if nums[middleIndex] > target {
		return search(nums, target, start, middleIndex)
	}
	return 0
}
