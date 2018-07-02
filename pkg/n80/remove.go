package main

import "fmt"

func removeDuplicates(nums []int) int {
	lenth := len(nums)
	if lenth==0 {
		return 0
	}
	dup := 0
	index:=2
	for i:=2;i<lenth;i++  {
		if index>=lenth {
			break
		}
		index++
		if nums[i]==nums[i-1] && nums[i]==nums[i-2] {
			dup+=1
			if i+1<lenth {
				copy(nums[i:], nums[i+1:])
			}
			i-=1
		}

	}
	return lenth-dup
}


func removeDuplicates2(nums []int) int {
	var i = 0
	for _, num := range(nums) {
		if i < 2 || num > nums[i - 2] {
			nums[i] = num
			i++
		}
	}

	return i
}

func main()  {
	//nums := []int{0,0,1,1,1,1,2,3,3}
	//nums := []int{1,1,1}
	nums := []int{1,1,1,2,2,3}
	fmt.Println(nums)
	fmt.Println(removeDuplicates2(nums))
	fmt.Println(nums)
}