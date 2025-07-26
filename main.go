package main

import (
	"fmt"
	"github.com/nayeem-bd/Todo-App/cmd"
	"os"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Usage: go run main.go serve")
		return
	}
	if args[1] == "serve" {
		cmd.Serve()
	}

	if args[1] == "work" {
		cmd.Work()
	}
}
