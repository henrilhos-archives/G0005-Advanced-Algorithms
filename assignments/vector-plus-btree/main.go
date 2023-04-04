package main

import (
	"fmt"
	"strings"
)

type Vector struct {
	letter string
	root   *Node
}

type Node struct {
	value       string
	left, right *Node
}

func insert(root *Node, value string) *Node {
	if root == nil {
		return &Node{value: value}
	}

	if value < root.value {
		root.left = insert(root.left, value)
	} else {
		root.right = insert(root.right, value)
	}

	return root
}

func traverse(root *Node) {
	if root != nil {
		traverse(root.left)
		fmt.Printf("%s ", root.value)
		traverse(root.right)
	}
}

func addName(vector []Vector, name string) {
	firstLetter := strings.ToLower(name[0:1])

	for i, v := range vector {
		if v.letter == firstLetter {
			vector[i].root = insert(vector[i].root, name)
			return
		}
	}
}

func printNames(vector []Vector, includeEmpty bool) {
	for _, v := range vector {
		if v.root != nil || includeEmpty {
			fmt.Printf("%s: ", v.letter)
			traverse(v.root)
			fmt.Println()
		}
	}
}

func menu(vector []Vector) {
	var choice int
	for {
		fmt.Println()
		fmt.Println("Choose an option:")
		fmt.Println("1. Add a name")
		fmt.Println("2. Print names (excluding empty vectors)")
		fmt.Println("3. Print names (including empty vectors)")
		fmt.Println("4. Exit")

		fmt.Scanln(&choice)

		switch choice {
		case 1:
			var name string
			fmt.Println("Enter a name:")
			fmt.Scanln(&name)
			addName(vector, name)
		case 2:
			printNames(vector, false)
		case 3:
			printNames(vector, true)
		case 4:
			fmt.Println("Bye!")
			return
		default:
			fmt.Println("Invalid option!")
		}
	}
}

func main() {
	alphabet := "abcdefghijklmnopqrstuvwxyz"

	var vector []Vector
	for _, letter := range strings.Split(alphabet, "") {
		vector = append(vector, Vector{letter: letter})
	}

	menu(vector)
}
