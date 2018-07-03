package main


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

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil  {
		return head
	}
	if head.Next == nil && head.Val==val{
		return nil
	}

	if head.Val == val {
		head = head.Next
		for head != nil && head.Val == val {
			head = head.Next
		}
		// 值为 val 的 node，已经全部删除了
		return removeElements(head,val)
	}

	head.Next = removeElements(head.Next,val)
	return head
}

func removeElements2(head *ListNode, val int) *ListNode {
	curr:=head
	var prev *ListNode =nil
	for curr!=nil {
		if curr.Val==val {
			if prev==nil {
				head=head.Next
			}else {
				prev.Next=curr.Next
			}
			curr=curr.Next
			continue
		}
		prev=curr
		curr=curr.Next
	}
	return head
}

func main() {
	nums := []int{6,6,6,3,4,5,6}
	//nums := []int{1,2,6,3,4,5,6}
	//nums := []int{6}
	fmt.Println(nums)
	head := NewListNode(nums)
	List2Array(head)
	ret := removeElements2(head,6)
	fmt.Println(ret)
	List2Array(ret)
}


