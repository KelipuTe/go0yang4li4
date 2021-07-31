package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
} //链表结点

func main() {
	ln5 := ListNode{Val: 3}
	ln4 := ListNode{Val: 3, Next: &ln5}
	ln3 := ListNode{Val: 2, Next: &ln4}
	ln2 := ListNode{Val: 1, Next: &ln3}
	ln1 := ListNode{Val: 1, Next: &ln2}

	tpln := deleteDuplicates(&ln1)
	for tpln != nil {
		fmt.Printf("%d,", tpln.Val)
		tpln = tpln.Next
	}
}

//存在一个按升序排列的链表，删除所有重复的元素，使每个元素只出现一次

//83-删除排序链表中的重复元素，重复元素只保留一个结点(82,83)
func deleteDuplicates(head *ListNode) *ListNode {
	var tplnLink *ListNode = head       //连接指针
	var tplnCheck *ListNode = head      //校验指针
	var tplnQuery *ListNode = head.Next //遍历指针

	//链表为空或者只有一个结点
	if head == nil || head.Next == nil {
		return head
	}

	//遍历到最后一个结点
	for tplnQuery.Next != nil {
		if tplnQuery.Val != tplnCheck.Val {
			//校验结点不重复，连接
			tplnLink.Next = tplnQuery
			tplnLink = tplnQuery
			tplnCheck = tplnQuery
		}
		tplnQuery = tplnQuery.Next
	}

	//处理最后一个结点
	if tplnQuery.Val != tplnCheck.Val {
		tplnLink.Next = tplnQuery
	} else {
		//断开连接结点后面的链表
		tplnLink.Next = nil
	}
	return head
}
