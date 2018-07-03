package main

import (
	"fmt"
)

func isToeplitzMatrix(matrix [][]int) bool {
	m := len(matrix)
	if m<2 {
		return true
	}
	n := len(matrix[0])
	if n<2 {
		return true
	}

	for i:=0;i<m;i++ {
		flag := handle(i,0,m,n,matrix)
		if !flag {
			return flag
		}
	}

	for j:=0;j<n;j++ {
		flag := handle(0,j,m,n,matrix)
		if !flag {
			return flag
		}
	}

	return true
}
func handle(i int, i2 int, m int, n int, martix [][]int) bool {
	val := martix[i][i2]
	for i<m && i2<n {
		//fmt.Printf("(%d,%d)%d\t",i,i2,martix[i][i2])
		if martix[i][i2]!=val {
			return false
		}
		i+=1
		i2+=1
	}
	return true
}

func main() {
	//matrix := [][]int{
	//	{1,2,3,4},
	//	{5,1,2,3},
	//	{9,5,1,2},
	//}
	matrix :=[][]int{
		{11,74,0,93},
		{40,11,74,7},
	}
	fmt.Println(isToeplitzMatrix(matrix))
}