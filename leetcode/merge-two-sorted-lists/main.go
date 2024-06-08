package main

import (
	"sort"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	listCode1 := &ListNode{1, &ListNode{2, &ListNode{4, nil}}}
	listCode2 := &ListNode{1, &ListNode{3, &ListNode{4, nil}}}
	// var result *ListNode
	mergeTwoLists(listCode1, listCode2)
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var store []int
	for list1 != nil {
		store = append(store, list1.Val)
		list1 = list1.Next
	}

	for list2 != nil {
		store = append(store, list2.Val)
		list2 = list2.Next
	}
	sort.Sort(sort.Reverse(sort.IntSlice(store)))
	headNode := &ListNode{}
	for _, v := range store {
		temp := headNode
		headNode = &ListNode{Val: v, Next: temp}
	}
	return headNode
}
