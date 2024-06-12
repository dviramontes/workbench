package workbench

// AddTwoNumbers stored in a singly-linked list in reserve order
func AddTwoNumbers(l1 *Node, l2 *Node) *Node {
	temp := new(Node)
	curr := temp
	carry := 0

	for l1 != nil || l2 != nil || carry > 0 {
		x, y := 0, 0
		if l1 != nil {
			x = l1.Data
		}

		if l2 != nil {
			y = l2.Data
		}

		sum := carry + x + y
		carry = sum / 10
		data := sum % 10
		curr.Next = &Node{Data: data}
		curr = curr.Next

		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}

	return temp.Next
}
