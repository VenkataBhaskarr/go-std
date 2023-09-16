## This is a ArrayList package in golang

## Usage 

Run command ` go get "github.com/VenkataBhaskarr/go-std" `

```
import (
    "github.com/VenkataBhaskarr/go-std/list"
)
func main() {
    // Create a new list of type string
    myList := list.NewList[string]()

    // Add elements
    myList.add("Hello")
    myList.add("World")

    // Get elements
    value1 := myList.get(0) // Returns "Hello"
    value2 := myList.get(1) // Returns "World"

    // Modify elements
    myList.set(1, "New World")

    // Remove elements
    myList.remove(0)

    // Check size
    size := myList.size() // Returns the number of elements in the list

    // Check if empty
    isEmpty := myList.isEmpty() // Returns true if the list is empty

    // Clear the list
    myList.clear()

    // Check if an element is in the list
    contains := myList.contains("Search Value")

    // Find the index of an element
    index := myList.indexOf("Search Value")

    // Create a sublist
    sublist := myList.subList(1, 3)

    // Sort the list
    myList.sort()

    // Reverse the list
    myList.reverse()
}
```


