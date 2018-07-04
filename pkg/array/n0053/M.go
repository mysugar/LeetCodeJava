package main

import (
	"fmt"
	"net/http"
)

//202 / 202 test cases passed.
//Status: Accepted
//Runtime: 8 ms
//https://leetcode.com/submissions/detail/162012812/

func maxSubArray(nums []int) int {
	sum := 0
	maxSum := nums[0]
	for _, num := range nums {
		sum += num
		if sum > maxSum {
			maxSum = sum
		}
		if sum < 0 {//如果小于0则丢弃重新开始计算
			sum = 0
		}
	}
	return maxSum
}

//202 / 202 test cases passed.
//Status: Accepted
//Runtime: 8 ms
//https://leetcode.com/submissions/detail/162023075/
//分治解法
func maxSubArray2(nums []int) int {
	if len(nums)<1{
		return 0
	}
	return recursionMax(nums, 0, len(nums)-1)
}

func recursionMax(nums []int, left int, right int) int {
	if left == right {
		return nums[left]
	}

	// 从中间开始，最后拼接起来也为连续序列
	center := (left + right) / 2
	maxLeftBorderSum := nums[center]
	leftBorderSum := 0
	for i := center; i >= left; i-- {
		leftBorderSum += nums[i]
		if leftBorderSum > maxLeftBorderSum {
			maxLeftBorderSum = leftBorderSum
		}
	}
	maxRightBorderSum := nums[center+1]
	rightBorderSum := 0
	for j := center + 1; j <= right; j++ {
		rightBorderSum += nums[j]
		if rightBorderSum > maxRightBorderSum {
			maxRightBorderSum = rightBorderSum
		}
	}

	maxLeftSum := recursionMax(nums, left, center)
	maxRightSum := recursionMax(nums, center+1, right)
	tmp := max(maxLeftSum, maxRightSum)
	//fmt.Println(maxLeftSum,maxRightSum,maxLeftBorderSum,maxRightBorderSum)
	return max(tmp, maxLeftBorderSum+maxRightBorderSum)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}




func main() {
	//nums := []int{-3,-2,-1}
	nums := []int{8,-19,5,-4,20}
	//nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray2(nums))
}
