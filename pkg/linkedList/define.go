package linkedList

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func NewListNode(arrs []int) *ListNode {
	var head *ListNode
	var now *ListNode
	for _,num := range(arrs){
		if nil==head {
			head=&ListNode{num,nil}
			now=head
		}else {
			curr:=&ListNode{num,nil}
			now.Next=curr
			now=curr
		}
		//fmt.Println(head)
	}
	return head
}

func List2Array(head *ListNode)  {
	curr := head
	for curr!=nil {
		fmt.Print(curr.Val,"\t")
		curr=curr.Next
	}
	fmt.Println()
}
