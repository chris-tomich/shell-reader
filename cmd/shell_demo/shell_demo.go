package main

import (
	"fmt"

	"github.com/chris-tomich/shell-reader/shell"
	"github.com/fatih/color"
)

func main() {
	r, err := shell.NewReader()

	if err != nil {
		panic(err)
	}

	defer r.Close()

	fmt.Println("demo 'echo shell' which will echo anything you type")
	fmt.Println("use 'exit' or Ctrl+C to close the shell")

	for {
		fmt.Printf("$ ")
		line, err := r.ReadLine()

		if err != nil {
			panic(err)
		}

		if line == "exit" {
			break
		}

		color.Green("%v\n", line)
	}
}
