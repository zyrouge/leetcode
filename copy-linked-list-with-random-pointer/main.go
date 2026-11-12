package main

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// func copyRandomList(head *Node) *Node {
// 	clones := make(map[*Node]*Node)
// 	var dfs func(node *Node) *Node
// 	dfs = func(node *Node) *Node {
// 		if node == nil {
// 			return nil
// 		}
// 		clone := &Node{node.Val, nil, nil}
// 		clones[node] = clone
// 		clone.Next = dfs(node.Next)
// 		clone.Random = clones[node.Random]
// 		return clone
// 	}
// 	return dfs(head)
// }

// func copyRandomList(head *Node) *Node {
// 	clones := make(map[*Node]*Node)
// 	current := head
// 	for current != nil {
// 		clone := &Node{current.Val, nil, nil}
// 		clones[current] = clone
// 		current = current.Next
// 	}
// 	current = head
// 	for current != nil {
// 		clone := clones[current]
// 		clone.Next = clones[current.Next]
// 		clone.Random = clones[current.Random]
// 		current = current.Next
// 	}
// 	return clones[head]
// }

func copyRandomList(head *Node) *Node {
	clones := make(map[*Node]*Node)
	current := head
	for current != nil {
		clone := &Node{current.Val, nil, nil}
		clones[current] = clone
		current = current.Next
	}
	return clones[head]
}
