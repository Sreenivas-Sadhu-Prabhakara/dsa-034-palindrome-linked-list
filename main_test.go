package main

import (
	"reflect"
	"testing"
)

func toSlice(head *Node) []int {
	out := []int{}
	for n := head; n != nil; n = n.Next {
		out = append(out, n.Val)
	}
	return out
}

func TestIsPalindrome(t *testing.T) {
	cases := []struct {
		in   []int
		want bool
	}{
		{[]int{1, 2, 2, 1}, true},
		{[]int{1, 2, 3, 2, 1}, true},
		{[]int{1, 2}, false},
		{[]int{1}, true},
		{[]int{}, true},
		{[]int{1, 2, 3, 1}, false},
	}
	for _, c := range cases {
		if got := IsPalindrome(fromSlice(c.in)); got != c.want {
			t.Errorf("IsPalindrome(%v) = %v, want %v", c.in, got, c.want)
		}
	}
}

func TestListRestored(t *testing.T) {
	head := fromSlice([]int{1, 2, 3, 2, 1})
	IsPalindrome(head)
	if got := toSlice(head); !reflect.DeepEqual(got, []int{1, 2, 3, 2, 1}) {
		t.Errorf("list not restored after check: %v", got)
	}
	head2 := fromSlice([]int{1, 2, 3, 4})
	IsPalindrome(head2)
	if got := toSlice(head2); !reflect.DeepEqual(got, []int{1, 2, 3, 4}) {
		t.Errorf("list not restored after negative check: %v", got)
	}
}
