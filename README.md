# shell-reader
shell-reader is a really basic Go library to create simple shell environments for VT-1xx compatible terminals. I personally had a need to be able to create a very simple shell to communicate with my services but found it difficult to implement simple features like navigating command history with arrow keys. Other Go libraries I looked into were made for really complex terminal interfaces and I just wanted something super simple. When I started reading into what it takes to accept keys (like arrow up and down) and to communicate with VT-1xx compatible terminals, I decided to expose this feature set as this really simple library.

# Installation
Install and update this go package with `go get -u github.com/chris-tomich/shell-reader`

# Usage
There's a complete example of how to use the library in `cmd/shell_demo/shell_demo.go` but below is a bare minimum version.

The most important points to note in this example are -

1. The `reader` needs to have `Close()` run as NewReader sets the terminal to CBreakMode and failing to run `Close()` will mean the terminal remains in this mode.
1. `shell.NewReader()` shouldn't be called more than once in a single application. (Will look to fix this in the future)
1. `shell.NewReader()` sets up a channel that listens for a SIGINT and will cleanup appropriately. If a SIGKILL is called, there will be no cleanup and this could cause unexpected behaviour in the terminal. Restarting the terminal will fix these issues.

```
import "github.com/chris-tomich/shell-reader/shell"
```
```
	reader, _ := shell.NewReader()

	defer reader.Close()

	for {
		fmt.Printf("$ ")
		line, _ := reader.ReadLine()

		color.Green("%v\n", line)
	}
```