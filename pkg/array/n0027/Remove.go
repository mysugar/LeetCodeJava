package main

import "fmt"

func removeElement(nums []int, val int) int {
	rmCount := 0
	length := len(nums)
	for index,num := range nums {
		if num==val {
			rmCount++
		}else {
			nums[index-rmCount]=num
		}

	}
	return length-rmCount
}

func main() {
	nums:=[]int{0,1,2,2,3,0,4,2}
	fmt.Println(removeElement(nums,2))
}
