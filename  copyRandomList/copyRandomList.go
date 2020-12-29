package main

type Node struct {
     Val int
     Next *Node
     Random *Node
 }

func copyRandomList(head *Node) *Node {
	if head == nil{
		return head
	}
	n := head
	for n != nil{
		node := &Node{
			Val:    n.Val,
			Next:   nil,
			Random: nil,
		}
		tmp := n.Next
		n.Next = node
		node.Next = tmp
		n = tmp
	}
	n1 := head
	n2 := head.Next
	for n1 != nil{
		r := n1.Random
		if r != nil{
			n2.Random = r.Next
		}else{
			n2.Random = nil
		}
		n1 = n2.Next
		if n1 != nil{
			n2 = n1.Next
		}
	}
	n1 = head
	res := n1.Next
	for n1 != nil{
		n2 = n1.Next
		n1.Next = n2.Next
		if n1.Next != nil{
			n2.Next = n1.Next.Next
		}else{
			n2.Next = nil
		}
		n1 = n1.Next
	}
	return res
}

