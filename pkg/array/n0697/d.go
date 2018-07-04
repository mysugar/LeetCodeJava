package main

import (
	"fmt"
)


//89 / 89 test cases passed.
//Status: Accepted
//Runtime: 40 ms
//https://leetcode.com/submissions/detail/161999539/

//输入: [1,2,2,3,1,4,2]
func findShortestSubArray(nums []int) int {
	//想知道哪个数出现次数最多
	dict := make(map[int][]int)
	max:=1
	for index,num := range nums {
		if arr,ok:=dict[num];ok {
			arr=append(arr,index)
			if len(arr) > max {
				max = len(arr)
			}
			dict[num]=arr
		}else {
			dict[num]=[]int{index}
		}
	}
	shortLen := len(nums)
	for _,v := range dict{
		degree :=len(v)
		if degree==max {
			lenTmp:=v[degree-1]-v[0]
			//fmt.Println(v[degree-1],v[0])
			if lenTmp < shortLen{
				shortLen=lenTmp
			}
		}
	}
	//fmt.Println(max)
	return shortLen+1
}

func main() {
	//nums := []int{1,2,2,3,1,4,2}
	nums := []int{1, 2, 2, 3, 1}
	fmt.Println(findShortestSubArray(nums))
}