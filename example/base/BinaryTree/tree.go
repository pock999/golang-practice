package main

import "fmt"

type Node struct{
	Val uint16;
	Left *Node;
	Right *Node;
}

type Tree struct{
	Root *Node;
}

func (tree *Tree)Init() {
	tree.Root = nil;
}

func (tree *Tree)AppendNode(n uint16) {
	newNode := &Node{
		Val: n,
		Left: nil,
		Right: nil,
	};

	current := tree.Root;
	previous := tree.Root;
	for(current != nil) {
		previous = current;
		if (newNode.Val < current.Val){
			current = current.Left;
		} else if(newNode.Val > current.Val) {
			current = current.Right;
		} else {
			fmt.Printf("(%d has already exist), ", n);
			break;
		}
	}

	if(previous == nil) {
		tree.Root = newNode;
		fmt.Printf("%d(root), ", n);
	} else if(newNode.Val > previous.Val) {
		previous.Right = newNode;
		fmt.Printf("%d(right), ", n);
	} else if(newNode.Val < previous.Val) {
		previous.Left = newNode;
		fmt.Printf("%d(left), ", n);
	}
}

func preorderTraversal(rootp *Node) {
	current := rootp;
	if(current != nil) {
		fmt.Printf("%d, ", current.Val);
		preorderTraversal(current.Left);
		preorderTraversal(current.Right);
	}
}
func inorderTraversal(rootp *Node) {
	current := rootp;
	if(current != nil) {
		inorderTraversal(current.Left);
		fmt.Printf("%d, ", current.Val);
		inorderTraversal(current.Right);
	}
}
func postorderTraversal(rootp *Node) {
	current := rootp;
	if(current != nil) {
		postorderTraversal(current.Left);
		postorderTraversal(current.Right);
		fmt.Printf("%d, ", current.Val);
	}
}

func main() {
	tree := new(Tree);
	tree.Init();
	fmt.Printf("insert: ");
	tree.AppendNode(5);
	tree.AppendNode(3);
	tree.AppendNode(2);
	tree.AppendNode(7);
	tree.AppendNode(6);
	tree.AppendNode(3);
	fmt.Println("");
	fmt.Printf("Pre-Order: ");
	preorderTraversal(tree.Root);
	fmt.Println("");
	fmt.Printf("In-Order: ");
	inorderTraversal(tree.Root);
	fmt.Println("");
	fmt.Printf("Post-Order: ");
	postorderTraversal(tree.Root);
	fmt.Println("");
}