package main

import "fmt"

func moveZeroes(nums []int)  {
	zeroC := 0
	length := len(nums)
	for index,num := range nums {
		if num==0 {
			zeroC++
		}
		if num!=0 {
			nums[index-zeroC]=num
		}
	}
	for i:=zeroC;i>0;i-- {
		nums[length-i]=0
	}
}

func main() {
	nums:=[]int{0,1,0,3,12}
	moveZeroes(nums)
	fmt.Println(nums)
}
