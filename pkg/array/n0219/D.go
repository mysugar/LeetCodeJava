package main

import "fmt"


//23 / 23 test cases passed.
//Status: Accepted
//Runtime: 20 ms
//https://leetcode.com/submissions/detail/161927161/

func containsNearbyDuplicate(nums []int, k int) bool {
	dict := make(map[int]int,len(nums))
	for i,num := range nums{
		if val, ok := dict[num]; ok {
			if i-val<=k  {
				return true
			}
		}
		dict[num]=i//由于i>val 越小越容易取得true
	}
	return false
}

func main() {
	nums := []int{1,0,1,1}
	//nums := []int{1,1,1,3,3,4,3,2,4,2}
	fmt.Println(containsNearbyDuplicate(nums,1))
}