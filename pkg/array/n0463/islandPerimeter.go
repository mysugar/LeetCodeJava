package main

import "fmt"

func islandPerimeter(grid [][]int) int {
	ret :=0
	rowN := len(grid)
	var colN int
	if rowN>0 {
		colN=len(grid[0])
	}else {
		return ret
	}
	for i,row := range grid{
		for j,col:=range row{
			if col==1 {
				ret += getPointP(i,j,rowN,colN,grid)
			}
		}
	}
	return ret
}
func getPointP(i int, j int, rowN int, colN int, grid [][]int) int {
	count:=0
	if (i==rowN || j==colN || i<0 || j<0){
		return count
	}
	if grid[i][j]==0 {
		return count
	}else {
		//判断当前节点
		tmp := getPoint(i , j , rowN , colN , grid)
		return tmp
	}
}
func getPoint(i int, j int, rowN int, colN int, grid [][]int) int {
	n:=4
	//任一方向有点，则-1
	if i+1<rowN && grid[i+1][j]!=0 {
		n--
	}
	if i-1>=0 && grid[i-1][j]!=0 {
		n--
	}
	if j+1<colN && grid[i][j+1]!=0 {
		n--
	}
	if j-1>=0 && grid[i][j-1]!=0 {
		n--
	}
	return n
}

func main() {
	nums := [][]int{
		{0,1,0,0},
		{1,1,1,0},
		{0,1,0,0},
		{1,1,0,0},
	}
	fmt.Println(islandPerimeter(nums))
}