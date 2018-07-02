package n143

import (
	"testing"
	"fmt"
)

func TestBackend(t *testing.T) {
	e := &ListNode{5,nil}
	d := &ListNode{4,e}
	c := &ListNode{3,d}
	b := &ListNode{2,c}
	a := &ListNode{1,b}

	//a := nil
	//lprint(a)
	reorderList(a)
	//fmt.Println("x")
	for _,v := range list2Array(a){
		fmt.Println("@",v.Val)
	}
}

func lprint(node *ListNode)  {
	fmt.Println(node.Val)
	fmt.Println("s")
	n := node
	for i:=0;n.Next!=nil ;i++  {
		fmt.Print(node.Val)
		n=n.Next
	}
	fmt.Print(n.Val)
	fmt.Println()
}

