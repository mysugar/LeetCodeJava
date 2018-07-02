package n143

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

 type ListNode struct {
	     Val int
	     Next *ListNode
	     }
func reorderList(head *ListNode)  {
	if head==nil || head.Next==nil   {
		return
	}
	array := list2Array(head)
	lth := len(array)
	if lth<2 {
		return
	}
	//arr2 := array
	for i:=0;i<lth;i++ {
		if lth-i-1==i+1 {//完成对整个链表的操作
			array[lth-i-1].Next=nil

			return
		}
		if i==lth-i-1 {
			fmt.Println(i)
			fmt.Println(lth-i-1)
			fmt.Println(array[i].Val)
			array[i].Next=nil //0
			array[lth-i].Next=array[i] //n
			fmt.Println("@@")
			for _,v := range array{
				fmt.Println("@",v.Val)
			}
			return
		}

		//插入后移
		//array[lth-i-1].Next=array[i+1]
		array[i].Next=array[lth-i-1] //0
		array[lth-i-1].Next=array[i+1] //n
	}
	return
}
func list2Array(node *ListNode) []*ListNode {
	var arr []*ListNode
	arr=append(arr,node)
	for ;node.Next!=nil ;  {
		arr=append(arr,node.Next)
		node=node.Next
	}
	return arr
}

