Run command ` go get github.com/VenkataBhaskarr/go-std/ `
```
import (
  "github.com/VenkataBhaskarr/go-std/queue"
)

  myQueue := queue.NewQueue[int]()
	// Add elements to the Queue
	myQueue.Add(1)
	myQueue.Add(2)
	myQueue.Add(3)

	// Get and print the size of the Queue
	fmt.Printf("Size of the Queue: %d\n", myQueue.Size())

	// Peek at the front element of the Queue
	frontElement := myQueue.Peek()
	fmt.Printf("Front element of the Queue: %d\n", frontElement)

	// Poll (remove and return) elements from the Queue
	for !myQueue.IsEmpty() {
		element := myQueue.Poll()
		fmt.Printf("Popped element: %d\n", element)
	}

	// Check if the Queue is empty
	if myQueue.IsEmpty() {
		fmt.Println("The Queue is empty.")
	} else {
		fmt.Println("The Queue is not empty.")
	}
}

```






