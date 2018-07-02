package n1

import (
	"sort"
	"fmt"
)

//func twoSum(nums []int, target int) []int {
//	if nums==nil ||len(nums)<2 {
//		return nil
//	}
//	dict := make(map[int]int)
//	ret := make([]int,2)
//	for index,num := range nums {
//		i2,ok:=dict[target-num];
//		if ok{
//			ret[0]=i2
//			ret[1]=index
//			return ret
//		}
//		dict[num]=index
//	}
//	return nil
//}

//func twoSum(nums []int, target int) []int {
//	if nums==nil ||len(nums)<2 {
//		return nil
//	}
//	//fmt.Println(nums,target)
//	dict := make(map[int]int)
//	ret := make([]int,2)
//	for index,num := range nums {
//		i2,ok:=dict[target-num];
//		if ok{
//			if nums[i2]>= nums[index]{
//				ret[1]=nums[i2]
//				ret[0]=nums[index]
//			}else {
//				ret[0]=nums[i2]
//				ret[1]=nums[index]
//			}
//			return ret
//		}
//		dict[num]=index
//	}
//	return nil
//}
//
//func threeSum(nums []int) [][]int {
//	var result [][]int
//	for index,num := range nums{
//		//fmt.Println(len(nums[:index]),len(nums[index+1:]))
//		//fmt.Println(nums)
//		//fmt.Println("@@",num)
//		lth := len(nums)
//		tmp := nums[:lth:lth]
//		ret2 := twoSum(joinSlce(tmp[:index], tmp[index+1:]) ,0-num)
//		if ret2!=nil {
//			ret := appendTo(ret2,num)
//			result=append(result,ret)
//			//fmt.Println(ret2,num,ret)
//		}
//	}
//	return result
//}
//
//func joinSlce(x,y []int) []int {
//	ret := make([]int,len(x)+len(y))
//	if len(x)>0 {
//		copy(ret[:len(x)-1],x[:len(x)-1])
//	}
//	copy(ret[len(x):],y[:])
//	return ret
//}
//
//func appendTo(nums []int,num int) []int {//排序插入
//	fmt.Println("@",nums,num)
//	for index,value := range nums{
//		if value>num {
//			//ret := make([]int,len(nums)+1)
//			if index==0 {
//				//ret[0]=num
//				return append([]int{num}, nums...)
//			}else {
//				nums3 := make([]int,len(nums)+1)
//				copy(nums3[:index],nums[:index])
//				nums3[index]=num
//				//fmt.Println(index,nums[:index],nums[index:])
//				copy(nums3[index+1:],nums[index:])
//				return nums3
//			}
//			//else {
//			//	fmt.Println("@",ret,nums,append(append(nums[:index], num),nums[index:]...))
//			//	copy(ret[0:index-1],nums[0:index-1])
//			//	//ret[0:index-1]=nums[0:index-1]
//			//	ret[index]=num
//			//	return append(ret,nums[index:]...)
//			//}
//			//return append(nums[:index],num,nums[index+1:])
//		}
//	}
//	return append(nums,num)
//}

func threeSum(nums []int) [][]int {
	var result [][]int
	sort.Ints(nums)
	for index,num := range nums{
		if num>=len(nums)-2 {
			break
		}
		if index > 0 && nums[index-1] == nums[index] {
			continue
		}
		j, k := index+1, len(nums)-1
		for j < k {
			if j > index+1 && nums[j] == nums[j-1] {
				j++
				continue
			}//去重
			if nums[index]+nums[j]+nums[k] == 0 {
				result = append(result, []int{nums[index], nums[j], nums[k]})
				j++
				k--
			} else if nums[index]+nums[j]+nums[k] < 0 {
				j++
			} else {
				k--
			}
		}
	}
	return result
}

func fourSum(nums []int, target int) [][]int {
	var result [][]int
	sort.Ints(nums)
	fmt.Println(nums)
	for index,num := range nums {
		if index >= len(nums)-3 {
			break
		}
		if index > 0 && nums[index-1] == nums[index] {
			continue
		}
		ret := threeSumHelp(nums[index+1:],target-num)
		if len(ret)!=0 {
			for _,v := range ret {
				r:= append([]int{num},v...)
				result = append(result,r)
			}
		}
	}
	return result
}

func threeSumHelp(nums []int,target int) [][]int {
	var result [][]int
	for index,_ := range nums{
		if index>=len(nums)-2 {
			break
		}
		if index > 0 && nums[index-1] == nums[index] {
			continue
		}
		j, k := index+1, len(nums)-1
		for j < k {
			if j > index+1 && nums[j] == nums[j-1] {
				j++
				continue
			}//去重
			if nums[index]+nums[j]+nums[k] == target {
				result = append(result, []int{nums[index], nums[j], nums[k]})
				j++
				k--
			} else if nums[index]+nums[j]+nums[k] < target {
				j++
			} else {
				k--
			}
		}
	}
	return result
}