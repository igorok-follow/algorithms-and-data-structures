package main

import "log"

type Node struct {
	Value    int
	Children []*Node
}

var tree = &Node{
	Value: 5,
	Children: []*Node{
		{
			Value: 5,
			Children: []*Node{
				{
					Value:    5,
					Children: []*Node{},
				},
				{
					Value: 1,
					Children: []*Node{
						{
							Value:    5,
							Children: []*Node{},
						},
						{
							Value:    1,
							Children: []*Node{},
						},
					},
				},
				{
					Value:    1,
					Children: []*Node{},
				},
			},
		},
		{
			Value: 1,
			Children: []*Node{
				{
					Value:    5,
					Children: []*Node{},
				},
				{
					Value:    1,
					Children: []*Node{},
				},
				{
					Value: 9,
					Children: []*Node{
						{
							Value:    1,
							Children: []*Node{},
						},
						{
							Value:    1,
							Children: []*Node{},
						},
					},
				},
			},
		},
	},
}

func main() {
	log.Println(roundRecursive([]*Node{tree}))
	log.Println(roundIterative(tree))
}

func roundRecursive(tree []*Node) int {
	sum := 0

	for _, node := range tree {
		sum += node.Value
		if len(node.Children) != 0 {
			sum += roundRecursive(node.Children)
		}
	}
	return sum
}

func roundIterative(tree *Node) int {
	var (
		sum   = tree.Value
		stack = tree.Children
		node  *Node
	)
	for len(stack) != 0 {
		node, stack = stack[len(stack)-1], stack[:len(stack)-1]
		sum += node.Value
		if len(node.Children) != 0 {
			stack = append(stack, node.Children...)
		}
	}

	return sum
}
