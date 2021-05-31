// use struct
package main

import "fmt"

type Node struct {
	Val uint16;
	Next *Node;
}

type LinkedList struct {
	Start *Node;
}

func (list *LinkedList)Init() {
	list.Start = new(Node);
	list.Start.Next = nil;
}

func (list *LinkedList)AppendNode(n uint16) {
	newNode := new(Node);
	newNode.Val = n;
	newNode.Next = nil;

	current := list.Start;
	for(current.Next != nil) {
		current = current.Next;
	}

	current.Next = newNode;
	fmt.Printf("%d->", n);
}

func (list *LinkedList)PrintList() {
	current := list.Start;
	fmt.Printf("LinkedList: ");
	for(current.Next != nil) {
		current = current.Next;
		fmt.Printf("%d->", current.Val);
	}
	fmt.Printf("(null)\n");
}

func main() {
	list := new(LinkedList);
	list.Init();

	fmt.Println("=== append number ===");
	fmt.Printf("append: ");
	list.AppendNode(1);
	list.AppendNode(2);
	list.AppendNode(3);
	list.AppendNode(4);
	fmt.Printf("(null)\n");

	list.PrintList();

	fmt.Println("=== end ===");
}