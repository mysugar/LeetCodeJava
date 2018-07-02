package n1

import (
	"testing"
	"fmt"
)

func TestBackend(t *testing.T) {
	//x := []int{-1,0,1,2,-1,-4}
	//fmt.Println(threeSum(x))

	nums := []int{0,1,5,0,1,5,5,-4}
	fmt.Println(fourSum(nums,11))
}

//func TestX(t *testing.T)  {
//
//	nums:=[]int{1,2,3}
//	nums2 := make([]int,5)
//	copy(nums2,nums)
//	num := 10
//	nums3 := appendTo(nums,num)
//	appendTo(nums[:2],num)
//	fmt.Println(nums)
//	fmt.Println(nums2)
//	fmt.Println(nums3)
//	fmt.Println(append(nums2[:4],nums...))
//	fmt.Println(nums2)
//
//}