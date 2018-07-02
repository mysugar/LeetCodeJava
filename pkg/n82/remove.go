package main

import (
	"fmt"
)


type ListNode struct {
   Val int
   Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// 处理 head 重复情况
	// 上面刚刚检查过了，head != nil && head.Next != nil
	if head.Val == head.Next.Val {
		val := head.Val
		head = head.Next.Next
		for head != nil && head.Val == val {
			head = head.Next
		}
		// 值为 val 的 node，已经全部删除了
		return deleteDuplicates(head)
	}
	// 处理 head 后面元素出现重复的情况
	head.Next = deleteDuplicates(head.Next)
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
	nums := []int{1,1,1,2,2,3}
	fmt.Println(nums)
	head := NewListNode(nums)
	List2Array(head)
	ret := deleteDuplicates(head)
	fmt.Println(ret)
	List2Array(ret)
}