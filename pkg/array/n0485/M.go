package main

import "fmt"

func findMaxConsecutiveOnes(nums []int) int {
	count := 0
	ret := 0
	for _,num := range(nums) {
		if num==1 {
			count+=1
		}else {
			if ret < count{
				ret=count
			}
			count=0
		}
	}
	if ret < count{
		ret=count
	}
	return ret
}

func main() {
	nums := []int{1,1,1,1,0,1,1,1,0,1,1,1,1,1}
	fmt.Println(findMaxConsecutiveOnes(nums))
}