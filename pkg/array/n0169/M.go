package main

import (
	"sort"
	"fmt"
)


//44 / 44 test cases passed.
//Status: Accepted
//Runtime: 28 ms
//https://leetcode.com/submissions/detail/161918929/

func majorityElement(nums []int) int {
	sort.Ints(nums)
	length := len(nums)
	index:=0
	for index<length/2 {
		//fmt.Println(nums[index],nums[index+length/2])
		if nums[index]==nums[index+length/2] {
			return nums[index]
		}
		index++
	}
	return nums[index]
}

func majorityElement2(nums []int) int {
	count := 0
	var data int
	for i := 0; i < len(nums); i++ {
		if count == 0 {
			data = nums[i]
		}

		if nums[i] == data {
			count += 1
		} else {
			count += -1
		}
	}
	return data
}

func main() {
	nums := []int {2,2,1,1,1,2,2}
	fmt.Println(majorityElement(nums))
}