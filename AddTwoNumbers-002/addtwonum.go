package AddTwoNumbers_002

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l3 := ListNode{}
	l3.RecursionAddTwoNumbers(l1, l2)
	return &l3
}

func (results *ListNode) RecursionAddTwoNumbers(l1 *ListNode, l2 *ListNode) {
	if l1 == nil || l2 == nil {
		return
	}
	i, y := addTwoNumbersHelper(l1.Val, l2.Val)
	fmt.Printf("Adding %d + %d = %d r%d\n", l1.Val, l2.Val, i, y)
	if y > 0 {
		results.addNodeRemainder(i, y)
	} else {
		results.addNode(i)
	}
	next := results.Next
	next.RecursionAddTwoNumbers(l1.Next, l2.Next)
}

func (results *ListNode) addNode(val int) {
	if results.Next != nil {
		results.Next.Val = results.Next.Val + val
	} else {
		results.Next = &ListNode{
			Val:  val,
			Next: nil,
		}
	}

}

func (results *ListNode) addNodeRemainder(val int, remainder int) {
	results.addNode(val)
	results.Next.addNode(remainder)
}

func addTwoNumbersHelper(i int, y int) (int, int) {
	result := i + y
	if result > 9 {
		return result - 10, 1
	}
	return result, 0

}
