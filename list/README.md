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
    myList.Add("Hello")
    myList.Add("World")

    // Get elements
    value1 := myList.Get(0) // Returns "Hello"
    value2 := myList.Get(1) // Returns "World"

    // Modify elements
    myList.Set(1, "New World")

    // Remove elements
    myList.Remove(0)

    // Check size
    size := myList.Size() // Returns the number of elements in the list

    // Check if empty
    isEmpty := myList.IsEmpty() // Returns true if the list is empty

    // Clear the list
    myList.Clear()

    // Check if an element is in the list
    contains := myList.Contains("Search Value")

    // Find the index of an element
    index := myList.IndexOf("Search Value")

    // Create a sublist
    sublist := myList.SubList(1, 3)

    // Sort the list
    myList.Sort()

    // Reverse the list
    myList.Reverse()
}
```


