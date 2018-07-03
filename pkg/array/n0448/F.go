package main

import "fmt"

//34 / 34 test cases passed.
//Status: Accepted
//Runtime: 552 ms
//https://leetcode.com/submissions/detail/161892388/

func findDisappearedNumbers(nums []int) []int {
	var ret []int
	index :=0
	for index<len(nums) {
		val := nums[index]
		if val!=index+1 {//当前位置放置的值不对
			num :=nums[index]
			tmp :=nums[num-1]
			if num==tmp {//要替换位置的值正确，则说明要替换的位置的值为重复值，当前位置即为缺少值
				//fmt.Println("@@@",index,nums)
				//ret = append(ret,index+1)
				index++
				continue
			}
			nums[num-1]=num
			nums[index]=tmp
		}else {
			index++
		}

	}
	fmt.Println(nums)
	for index,num :=range nums {
		if num!=index+1 {
			ret=append(ret,index+1)
		}
	}
	////repeat := 0
	////遍历出现重复的数据则一定就是缺数据了
	//for index,num :=range nums {
	//	fmt.Println(index,num-1)
	//	fmt.Println(nums)
	//	if index!=num-1 {
	//		//swap
	//		//把当前位置的值num放到指定位置num-1
	//		tmp :=nums[num-1]
	//		//if num==tmp {//要替换位置的值正确，则说明要替换的位置的值为重复值,不进行替换了
	//		//	//fmt.Println(num-1,tmp,num)
	//		//	//ret=append(ret,index+1)
	//		//	continue
	//		//}
	//		nums[num-1]=num
	//		//判断是否需要写入
	//		nums[index]=tmp
	//	}
	//	//该点放置的值为自己
	//	fmt.Println(nums)
	//}
	//for index,num :=range nums {
	//	if num!=index+1 {
	//		ret=append(ret,index+1)
	//	}
	//}

	return ret
}

func main() {
	//nums:=[]int{3,1,2,3,5}
	nums:=[]int{5,5,5,5,5}
	//nums:=[]int{4,3,2,7,8,2,3,1}
	fmt.Println(nums)
	fmt.Println(findDisappearedNumbers(nums))
}