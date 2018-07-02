package main

import "fmt"


type ListNode struct {
   Val int
   Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	var pre *ListNode
	curr := head
	if curr==nil {
		return head
	}
	for curr !=nil {
		if pre!=nil && pre.Val==curr.Val{
			pre.Next=curr.Next
			curr=curr.Next
		}else {
			pre=curr
			curr=curr.Next
		}
	}
	return head
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

func main()  {
	//nums := []int{0,0,1,1,1,1,2,3,3}
	//nums := []int{1,1,1}
	nums := []int{1,2,2}
	fmt.Println(nums)
	head := NewListNode(nums)
	List2Array(head)
	fmt.Println(deleteDuplicates(head))
	List2Array(head)
}