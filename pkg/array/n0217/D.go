package main

import "fmt"

func containsDuplicate(nums []int) bool {
	dict := make(map[int]int,len(nums))
	for _,num := range nums{
		if _, ok := dict[num]; ok {
			return true
		}
		dict[num]=0
	}
	return false
}

func main() {
	nums := []int{1,2,3,4}
	//nums := []int{1,1,1,3,3,4,3,2,4,2}
	fmt.Println(containsDuplicate(nums))
}