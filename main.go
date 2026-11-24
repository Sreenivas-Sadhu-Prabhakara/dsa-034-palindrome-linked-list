package main

import "fmt"

// Node is a singly linked list node.
type Node struct {
	Val  int
	Next *Node
}

// IsPalindrome reports whether the list reads the same both ways, in O(1)
// space: find the middle, reverse the second half, compare, then restore
// the list before returning.
func IsPalindrome(head *Node) bool {
	if head == nil || head.Next == nil {
		return true
	}
	// Middle: slow ends at the last node of the first half.
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	second := reverse(slow.Next)

	ok := true
	for p, q := head, second; q != nil; p, q = p.Next, q.Next {
		if p.Val != q.Val {
			ok = false
			break
		}
	}

	slow.Next = reverse(second) // restore the original links
	return ok
}

func reverse(head *Node) *Node {
	var prev *Node
	for cur := head; cur != nil; {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
}

func fromSlice(vals []int) *Node {
	dummy := &Node{}
	cur := dummy
	for _, v := range vals {
		cur.Next = &Node{Val: v}
		cur = cur.Next
	}
	return dummy.Next
}

func main() {
	fmt.Println(IsPalindrome(fromSlice([]int{1, 2, 2, 1})))
}
