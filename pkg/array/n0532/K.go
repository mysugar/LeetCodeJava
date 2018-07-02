package main

import (
	"sort"
	"fmt"
)

func findPairs(nums []int, k int) int {

	length := len(nums)
	if  length<2{
		return 0
	}
	sort.Ints(nums)

	if  nums[length-1]-nums[0]<k{
		return 0
	}
	ret := 0
	for i:=0;i<length;i++ {
		if i>0 && nums[i]==nums[i-1] {
			continue
		}
		for j:=length-1;j>=0;j-- {
			if j<=i {
				break
			}
			if j<length-1 && nums[j]==nums[j+1] {
				continue
			}
			if nums[j]-nums[i]==k {
				ret+=1;
			}
			if nums[j]-nums[i]<k {
				break
			}
		}
	}
	return ret
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	k := 1
	fmt.Println(findPairs(nums,k))
}