package AddTwoNumbers_002

import (
	"fmt"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	a, b, c := 2, 4, 3
	l1 := &ListNode{
		Val:  a,
		Next: nil,
	}
	l1.Next = &ListNode{
		Val:  b,
		Next: nil,
	}
	l1.Next.Next = &ListNode{
		Val:  c,
		Next: nil,
	}

	d, e, f := 5, 6, 4
	l2 := &ListNode{
		Val:  d,
		Next: nil,
	}
	l2.Next = &ListNode{
		Val:  e,
		Next: nil,
	}
	l2.Next.Next = &ListNode{
		Val:  f,
		Next: nil,
	}

	l3 := AddTwoNumbers(l1, l2)
	fmt.Println()
	for n := l3; n != nil; n = n.Next {
		fmt.Printf("%d -> ", n.Val)
	}

}

func BenchmarkAddTwoNumbers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a, b, c := 2, 4, 3
		l1 := &ListNode{
			Val:  a,
			Next: nil,
		}
		l1.Next = &ListNode{
			Val:  b,
			Next: nil,
		}
		l1.Next.Next = &ListNode{
			Val:  c,
			Next: nil,
		}

		d, e, f := 5, 6, 4
		l2 := &ListNode{
			Val:  d,
			Next: nil,
		}
		l2.Next = &ListNode{
			Val:  e,
			Next: nil,
		}
		l2.Next.Next = &ListNode{
			Val:  f,
			Next: nil,
		}

		l3 := AddTwoNumbers(l1, l2)
		fmt.Println()
		for n := l3; n != nil; n = n.Next {
			fmt.Printf("%d -> ", n.Val)
		}
	}
}
