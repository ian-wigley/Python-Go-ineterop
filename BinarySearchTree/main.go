package main

/* Go Library, to build as a .so -> go build -buildmode=c-shared -o _interop.so */

import (
	"C"
	"fmt"
	"strconv"
	"unsafe"
)

var (
	oldResult = 0
	newResult = 0
	startValue = 0
	wasSequence = false
	startingValue = false
	george = ""
)

// Struct to hold each numeric value
type Node struct {
	word   int
	pLeft  *Node
	pRight *Node
}

func main() {}

//export binary_tree
func binary_tree(cArray *C.int, cSize C.int, i C.int) *C.char  {

    gSlice := (*[1 << 30]C.int)(unsafe.Pointer(cArray))[:cSize:cSize]

	fmt.Print("The values passed from Python are : ")
	var tree *Node = nil
	for _, value := range gSlice {
		fmt.Print(value)
		fmt.Print(", ")
		tree = add(tree, int(value))
	}
	fmt.Print("\n")
	populateRangeStart(tree)
	getRanges(tree)
	populateRangeEnd()

	return C.CString(george)
}

// Function to create a new Node
func add(tree *Node, word int) *Node {
	tempPtr := Node{
		word:   word,
		pLeft:  nil,
		pRight: nil,
	}
	return addNode(tree, &tempPtr)
}

// Function to add the words and frequencies to the newly created Node
func addNode(tree *Node, toAdd *Node) *Node {
	if tree == nil {
		return toAdd
	}
	if toAdd.word != -1 {
		// Insert recursively into Tree
		if toAdd.word < tree.word && toAdd.word != tree.word {
			tree.pLeft = addNode(tree.pLeft, toAdd)
			return tree
		}

		if toAdd.word > tree.word && toAdd.word != tree.word {
			tree.pRight = addNode(tree.pRight, toAdd)
			return tree
		}
	}
	return tree
}

func getRanges(tree *Node) {
	if tree != nil {
		getRanges(tree.pLeft)
		newResult = tree.word
		getRanges(tree.pRight)
	}

	// End of a number sequence
	if newResult-oldResult != 1 {
		if newResult != oldResult {
			if wasSequence {
				wasSequence = false
				startingValue = false
				george += strconv.Itoa(startValue) + "-" + strconv.Itoa(oldResult) + ","
			} else {
				george += strconv.Itoa(oldResult) + ","
			}
		}
		// Start of a number sequence
	} else {
		if !startingValue {
			if oldResult == 0 {
				startValue = newResult
			} else {
				startValue = min(oldResult, newResult)
			}
			startingValue = true
		}
		wasSequence = true
	}
	oldResult = newResult
}

func populateRangeStart(tree *Node) {
	if tree != nil {
		startValue = tree.word
		oldResult = 0
		newResult = 0
	}
}

func populateRangeEnd() {
	if wasSequence {
		george += strconv.Itoa(startValue) + "-" + strconv.Itoa(newResult)
	} else {
		george += strconv.Itoa(newResult)
	}
}

func min(value_one int, value_two int) int {
	if value_one < value_two {
		return value_one
	} else {
		return value_two
	}
}
