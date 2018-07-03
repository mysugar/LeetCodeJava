package main

import (
	"sort"
)

func arrayPairSum(nums []int) int {
	sort.Ints(nums)
	ret := 0
	for index,num := range(nums){
		if index%2==0 {
			ret+=num
		}
	}
	return ret
}
