package main

import "fmt"

func maxProfit(prices []int) int {
	length := len(prices)
	index := 0
	Profit := 0
	buyed := false
	buyPrice := 0
	for index<length-1 {
		if prices[index]<prices[index+1] && !buyed {//有利
			//buy
			buyed=true
			buyPrice=prices[index]
		}
		if prices[index]>=prices[index+1] && buyed{//有利
			buyed=false
			Profit=Profit+prices[index]-buyPrice
		}
		if prices[index]<prices[index+1] && !buyed {//有利
			//buy
			buyed=true
			buyPrice=prices[index]
		}

		index++
	}
	if buyed {
		Profit=Profit+prices[length-1]-buyPrice
	}
	return Profit
}

func main() {
	//nums := []int{1,2,3,4,5}
	nums := []int{7,6,4,3,1}
	fmt.Println(maxProfit(nums))
}