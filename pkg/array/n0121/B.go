package main

import "fmt"


//200 / 200 test cases passed.
//Status: Accepted
//Runtime: 8 ms
//https://leetcode.com/submissions/detail/161924317/

func maxProfit2(prices []int) int {//O(n^2)
	length := len(prices)
	Profit := 0
	for i,num := range prices{
		j:=i+1
		if j==length {
			break
		}
		for j<length{
			curP:=prices[j]-num
			Profit = max(Profit,curP)
			j++
		}
	}
	return Profit
}

func maxProfit(prices []int) int {
	length := len(prices)
	var begin, end, max int
	for end<length  {
		if prices[end]-prices[begin]>max {
			max=prices[end]-prices[begin]
		}
		if( prices[end]<=prices[begin]) {//当前价值比原来低，则更适合买入
			begin=end
		}
		end++
	}
	return max
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	//nums := []int{7,1,5,3,6,4}
	nums := []int{7,6,4,3,1}
	fmt.Println(maxProfit(nums))
}