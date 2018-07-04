package main

import "fmt"

func twoSum(numbers []int, target int) []int {
	length := len(numbers)
	dict := make(map[int]int,length)
	for index,num := range numbers{
		if j,ok:=dict[target-num];ok {
			return []int{j+1,index+1}
		}
		dict[num]=index
	}
	return []int{1,1}
}

func main() {
	nums := []int{2, 7, 11, 15}
	fmt.Println(twoSum(nums,9))
}