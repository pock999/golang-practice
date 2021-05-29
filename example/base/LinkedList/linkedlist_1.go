package main

import "fmt"

type Node struct {
	Val uint16;
	Next *Node;
}

func Init(startp *Node) {
	startp.Next = nil;
}

func AppendNode(n uint16, startp *Node) {
	newNode := new(Node);
	newNode.Val = n;
	newNode.Next = nil;

	current := startp;
	for(current.Next != nil) {
		current = current.Next;
	}

	current.Next = newNode;
	fmt.Printf("%d->", n);
}

func PrintList(startp *Node) {
	current := startp;
	fmt.Printf("LinkedList: ");
	for(current.Next != nil) {
		current = current.Next;
		fmt.Printf("%d->", current.Val);
	}
	fmt.Printf("(null)\n");
}

func main() {
	start := new(Node);
	Init(start);

	fmt.Println("=== append number ===");
	fmt.Printf("append: ");
	AppendNode(1, start);
	AppendNode(2, start);
	AppendNode(3, start);
	AppendNode(4, start);
	fmt.Printf("(null)\n");

	PrintList(start);

	fmt.Println("=== end ===");
}
