package main

import (
	"fmt"
	//"sort"
)
//图遍历 类似fillblood

func maxAreaOfIsland(grid [][]int) int {
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
				count := getIlandCount(i,j,rowN,colN,grid)
				if count>ret {
					ret=count
				}
			}
		}
	}
	return ret
}
func getIlandCount(i int, j int,rowN int, colN int, grid [][]int) int {
	if (i==rowN || j==colN || i<0 || j<0){
		return 0
	}
	if grid[i][j]==0 {
		return 0
	}else {
		grid[i][j]=0
		return 1+getIlandCount(i+1,j,rowN,colN,grid)+getIlandCount(i,j+1,rowN,colN,grid)+getIlandCount(i-1,j,rowN,colN,grid)+getIlandCount(i,j-1,rowN,colN,grid)
	}
}


func maxAreaOfIsland2(grid [][]int) int {
	//节点做key，iland号做value，
	//修改为map[string][]int
	//ilands := make([]int,250)
	ilands := make(map[int][]string)
	nodeMap := make(map[string]int)
	curid := 1
	for i,row := range grid{
		for j,col:=range row{
			if col==1 {
				//遍历已探索的四周，获取当前iland_id,如果不存在则使用当前id加入map，将当前id+1,当前iland+1
				hasIland,id :=getExploredID(i,j,nodeMap,ilands)
				if hasIland {
					point := fmt.Sprintf("%d,%d",i,j)
					nodeMap[point]=id
					ilands[id]=append(ilands[id], point)
				}else {
					point := fmt.Sprintf("%d,%d",i,j)
					nodeMap[point]=curid
					ilands[curid]=append(ilands[curid], point)
					//ilands[curid]++
					curid++
				}
			}
		}
	}
	ret:=0
	for _,val := range ilands {
		if len(val)>ret {
			ret=len(val)
		}
	}
	//fmt.Println(ilands)
	//fmt.Println(nodeMap)
	//sort.Sort(sort.Reverse(sort.IntSlice(ilands)))
	return ret
}
func getExploredID(i int, j int, nodeMap map[string]int,iland map[int][]string) (bool,int) {
	if i == 0 && j == 0 {
		return false, 0
	} else {

		leftFind := false
		leftId := 0
		//get left
		if j > 0 {
			left := fmt.Sprintf("%d,%d", i, j-1)
			leftId, leftFind = nodeMap[left]
		}
		upFind := false
		upId := 0
		//get up
		if i > 0 {
			up := fmt.Sprintf("%d,%d", i-1, j)
			upId, upFind = nodeMap[up]
		}
		//fmt.Println(i,j,upFind,leftFind)
		if upFind && leftFind {
			if upId != leftId {
				//merge and update
				upZone := iland[upId]
				//fmt.Println(i,j,upZone)
				//fmt.Println(iland[leftId])
				iland[leftId] = append(iland[leftId], upZone...)
				for _, point := range upZone {
					nodeMap[point] = leftId
				}
				delete(iland, upId)
			}
			return leftFind,leftId
		} else {
			//fmt.Println(upFind || leftFind, upId | leftId)
			return upFind || leftFind, upId | leftId
		}
		//return false,0
	}
}

func main() {
	nums := [][]int{
		{0,0,1,0,0,0,0,1,0,0,0,0,0},
		{0,0,0,0,0,0,0,1,1,1,0,0,0},
		{0,1,1,0,1,0,0,0,0,0,0,0,0},
		{0,1,0,0,1,1,0,0,1,0,1,0,0},
		{0,1,0,0,1,1,0,0,1,1,1,0,0},
		{0,0,0,0,0,0,0,0,0,0,1,0,0},
		{0,0,0,0,0,0,0,1,1,1,0,0,0},
		{0,0,0,0,0,0,0,1,1,0,0,0,0},
	}
	fmt.Println(maxAreaOfIsland(nums))

}