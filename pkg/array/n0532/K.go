package main

import (
	"sort"
	"fmt"
)

func findPair2s(nums []int, k int) int {

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
func findPairs(nums []int, k int) int {
	if k<0{
		return 0
	}

	data := make(map[int]int)
	for _, item := range nums {
		data[item]++
	}

	flag := k > 0

	count := 0
	for key, value := range data {
		if !flag && value > 1 {
			fmt.Println(value > 1 )
			count++
		}
		_, ok := data[key + k]
		if flag && ok {
			fmt.Println(key )
			count++
		}
	}
	return count

}
func main() {
	nums := []int{3, 1, 4, 1, 5,1}
	k := 2
	fmt.Println(findPairs(nums,k))
}