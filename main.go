package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Print("Please input an arguement")
		return
	}

	command := os.Args[1]

	switch command {
	case "add":

	case "list":
	case "update":
	case "delete":
	case "in progress":
	case "done":
	default:

	}

}
