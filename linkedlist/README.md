## usage

Run command ` go get github.com/VenkataBhaskarr/go-std `
```
import (
  "github.com/VenkataBhaskarr/go-std/linkedlist
)
myList := linkedlist.NewLinkedList[int]()

    // Add elements to the LinkedList
    myList.Add(1)
    myList.Add(2)
    myList.Add(3)

    // Get and print the size of the LinkedList
    fmt.Printf("Size of the LinkedList: %d\n", myList.Size())

    // Check if the LinkedList is empty
    if myList.IsEmpty() {
        fmt.Println("The LinkedList is empty.")
    } else {
        fmt.Println("The LinkedList is not empty.")
    }

    // Check if the LinkedList contains a specific element
    if myList.Contains(2) {
        fmt.Println("The LinkedList contains the element 2.")
    } else {
        fmt.Println("The LinkedList does not contain the element 2.")
    }

    // Get and print elements from the LinkedList
    fmt.Printf("Element at index 1: %d\n", myList.Get(1))

    // Remove an element from the LinkedList
    myList.Remove(0)

    // Get and print the updated size of the LinkedList
    fmt.Printf("Size of the LinkedList after removal: %d\n", myList.Size())

```
