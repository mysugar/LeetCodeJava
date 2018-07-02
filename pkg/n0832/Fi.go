package main

import (
	"fmt"
)

func flipAndInvertImage(A [][]int) [][]int {
	for index,array := range(A){
		A[index]=revers(array)
	}
	return A
}

func revers(B []int)  []int{
	lenth := len(B)
	for i:=0;i<(lenth+1)/2;i++ {
		tmp:=B[i] ^ 1
		B[i]=B[lenth-i-1] ^ 1
		B[lenth-i-1]=tmp
	}
	return B
}


func main() {
	nums := []int{1,1,0,0}
	//fmt.Println()
	fmt.Println(revers(nums))
}