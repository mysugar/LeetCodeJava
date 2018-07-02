package main

import "fmt"

func removeDuplicates(nums []int) int {
	lenth := len(nums)
	if lenth==0 {
		return 0
	}
	count :=1
	for i:=1;i<lenth;i++  {
		if !(nums[i]==nums[i-1]) {
			nums[count]=nums[i]
			count+=1
		}
	}
	return count
}


func main()  {
	nums := []int{0,0,1,1,1,2,2,3,3,4}
	fmt.Println(removeDuplicates(nums))
	fmt.Println(nums)
}