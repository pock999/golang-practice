package main

import "fmt"

type Node struct {
	Val uint16;
	Current uint16; // current height
	Above *Node;
	Under *Node;
}

type Stack struct {
	Limit uint16;
	Top *Node;
	Bottom *Node;
}

func (stack *Stack) Init(n uint16) {
	stack.Limit = n;
	stack.Top = &Node{
		Current: 0,
		Above: nil,
		Under: nil,
	};
	stack.Bottom = &Node{
		Current: 0,
		Above: nil,
		Under: nil,
	};
}

func (stack *Stack) Push(n uint16) {
	current := stack.Bottom;
	currentHeight := stack.Bottom.Current;

	for(current.Above != nil) {
		current = current.Above;
		currentHeight = current.Current;
	}

	if(currentHeight >= stack.Limit) {
		fmt.Println("\nover stack limit");
	} else {
		newNode := Node{
			Val: n,
			Above: nil,
			Under: nil,
		};
    
		newNode.Current = currentHeight + 1;
    newNode.Under = current;
		current.Above = &newNode;
		stack.Top.Under = &newNode;
    fmt.Printf("%d-", n);
	}
}

func (stack *Stack) Pop() (uint16, bool) {
  current := stack.Top.Under;

  if(current.Val == 0) {
    return 0, false;
  }
  
  pop := current.Val;
  stack.Top.Under.Above = nil;
  stack.Top.Under = current.Under;
  return pop, true;
}

func main() {
	stack := new(Stack);
	stack.Init(5);
  fmt.Println("=== Push start ===");
  stack.Push(1);
  stack.Push(2);
  stack.Push(4);
  stack.Push(5);
  stack.Push(6);
  stack.Push(7);
  stack.Push(3);
  fmt.Println("=== Push  end  ===");
  fmt.Println("=== Pop start ===");
  v, ok := stack.Pop();
  fmt.Printf("%d(%t)-", v, ok);
  v, ok = stack.Pop();
  fmt.Printf("%d(%t)-", v, ok);
  v, ok = stack.Pop();
  fmt.Printf("%d(%t)-", v, ok);
  v, ok = stack.Pop();
  fmt.Printf("%d(%t)-", v, ok);
  v, ok = stack.Pop();
  fmt.Printf("%d(%t)-", v, ok);
  v, ok = stack.Pop();
  fmt.Printf("%d(%t)-", v, ok);
  
  fmt.Println("=== Pop  end  ===");
}