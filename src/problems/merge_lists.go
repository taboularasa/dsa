package problems

type LinkedListNode struct {
	value int
	next  *LinkedListNode
}

func MergeLists(lists []*LinkedListNode) *LinkedListNode {
	var res *LinkedListNode
	for i := 0; i < len(lists); i++ {
		res = merge(res, lists[i])
	}
	return res
}

func merge(a, b *LinkedListNode) *LinkedListNode {
	head := &LinkedListNode{}
	n := head

	for a != nil && b != nil {
		if a.value < b.value {
			n.next = a
			a = a.next
		} else {
			n.next = b
			b = b.next
		}
		n = n.next
	}

	for a != nil {
		n.next = a
		a = a.next
		n = n.next
	}

	for b != nil {
		n.next = b
		b = b.next
		n = n.next
	}

	return head.next
}
