Run command ` go get github.com/VenkataBhaskarr/go-std/ `
```
import (
  "github.com/VenkataBhaskarr/go-std/queue"
)

  myStack := stack.NewStack[int]()
	// Push elements onto the Stack
	myStack.Add(1)
	myStack.Add(2)
	myStack.Add(3)

	// Get and print the size of the Stack
	fmt.Printf("Size of the Stack: %d\n", myStack.Size())

	// Peek at the top element of the Stack
	topElement := myStack.Peek()
	fmt.Printf("Top element of the Stack: %d\n", topElement)

	// Pop (remove and return) elements from the Stack
	for myStack.Size() > 0 {
		element := myStack.Pop()
		fmt.Printf("Popped element: %d\n", element)
	}

	// Check if the Stack is empty
	if myStack.Size() == 0 {
		fmt.Println("The Stack is empty.")
	} else {
		fmt.Println("The Stack is not empty.")
	}
}
