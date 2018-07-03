package main

import "fmt"

func matrixReshape(nums [][]int, r int, c int) [][]int {
	m:=len(nums)
	n:=len(nums[0])
	if (m*n)!=(r*c) {
		return nums
	}

	ret := make([][]int,r)
	for row := range ret {
		ret[row] = make([]int, c)
		for col := range ret[row] {
			vlen:=(row*c+col)
			mR:=vlen/n
			mN:=vlen%n
			//fmt.Printf("%d,%d %d,%d \r\n",i,j,mR,mN)
			ret[row][col]=nums[mR][mN]
		}
	}
	return ret
}

func main() {
	matrix :=[][]int{
		{1,2},
		{3,4},
	}
	fmt.Println(matrixReshape(matrix,1,4))
}