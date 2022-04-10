package main

import (
	"fmt"
)

type LinkedList struct {
	Head *ListNode
	Size int
}

type ListNode struct {
	Next  *ListNode
	Value string
}

// Insert
func (ll *LinkedList) Insert(insertMe string) {
	newNode := &ListNode{Value: insertMe, Next: nil}
	if ll.Head == nil {
		ll.Head = newNode
		ll.Size++
		return
	}

	current := ll.Head
	var temp *ListNode

	for current != nil {
		if current.Value == insertMe {
			fmt.Println("Item already exist, not inserting.")
			return
		}
		temp = current
		current = current.Next
	}

	temp.Next = newNode
	ll.Size++
	return
}

// Delete
func (ll *LinkedList) Delete(deleteMe string) {
	var current, temp *ListNode
	current = ll.Head

	// If the item we want to delete is the head value
	if ll.Head.Value == deleteMe {
		ll.Head = ll.Head.Next
		return
	}

	// [1, 2, 3]
	// If the item we want to delete is not the head value
	for current != nil {
		if current.Value == deleteMe {
			temp.Next = current.Next
			ll.Size--
			return
		}
		temp = current
		current = current.Next
	}
	return
}

// Search
func (ll *LinkedList) Search(findMe string) bool {
	current := ll.Head
	for current != nil {
		if current.Value == findMe {
			return true
		}
		current = current.Next
	}
	return false
}

// Print
func (ll *LinkedList) Print() {
	current := ll.Head
	for current != nil {
		fmt.Println(current.Value)
		current = current.Next
	}
	return
}

// [1,2,3]
func (ll *LinkedList) Reverse() {
	var current, back, front *ListNode
	current = ll.Head
	if ll.Size == 0 || ll.Size == 1 {
		return
	}

	for current.Next != nil {
		front = current.Next
		current.Next = back
		back = current
		current = front
	}

	current.Next = back
	ll.Head = current
}

func (ll *LinkedList) Get(getMe string) *ListNode {
	current := ll.Head

	for current != nil {
		if current.Value == getMe {
			return current
		}
		current = current.Next
	}
	return nil
}

func (ll *LinkedList) GetLengthRecurrsive() int {
	current := ll.Head

	size := lengthRecurrsive(current)
	return size
}

func lengthRecurrsive(node *ListNode) int {
	if node == nil {
		return 0
	}
	return 1 + lengthRecurrsive(node.Next)
}

func main() {
	ll := &LinkedList{}
	ll.Insert("ryan")
	ll.Insert("ryan")
	ll.Insert("christain")
	ll.Insert("chris")
	ll.Insert("andrew")
	ll.Insert("matthew")

	size := ll.GetLengthRecurrsive()
	fmt.Println("The size of get length recurrsive is: ", size)

	ll.Delete("matthew")
	ll.Delete("andrew")
	ll.Delete("christain")

	ll.Print()

	fmt.Println("==========================================")

	ll.Reverse()

	ll.Print()

	// Find
	// if doesRyanExist := ll.Search("ryan"); !doesRyanExist {
	// 	fmt.Println("ryan was not found")
	// } else {
	// 	fmt.Println("ryan was found")
	// }

	// if doesChristainExist := ll.Search("christain"); !doesChristainExist {
	// 	fmt.Println("christain was not found")
	// } else {
	// 	fmt.Println("christain was found")
	// }

	// if doesNotExist := ll.Search("dontFindMe"); !doesNotExist {
	// 	fmt.Println("dontFindMe was not found")
	// } else {
	// 	fmt.Println("dontFind me was found")
	// }
}
