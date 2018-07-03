package main

import "fmt"

func isOneBitCharacter(bits []int) bool {
	index:=0
	length:=len(bits)
	for index<length {
		if index==length-1 {
			return true
		}
		currBit:=bits[index]
		if index==length-2  {
			if currBit==1{
				return false
			}else {
				return true
			}
		}
		if currBit==0 {
			index++
			continue
		}
		if currBit==1 {
			index+=2
		}
	}
	return true
}

func main() {
	//nums := []int {1, 1, 1, 0}
	nums := []int {1, 0, 0}
	fmt.Println(isOneBitCharacter(nums))
}