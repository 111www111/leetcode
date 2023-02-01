package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	//定义一个总链
	returnNode := ListNode{}
	returnNode.Val = -1
	//定义指针
	retNode := &returnNode
	//定义进位
	add := 0
	for l1 != nil || l2 != nil || add != 0 {
		var x, y int
		if l1 != nil {
			x = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			y = l2.Val
			l2 = l2.Next
		}
		sum := x + y + add
		add = sum / 10
		value := sum % 10
		//推进链表
		retNode.Next = &ListNode{Val: value}
		retNode = retNode.Next
	}
	//有进位就再加一个
	if add > 0 {
		retNode.Val = add
	}
	return returnNode.Next
}
