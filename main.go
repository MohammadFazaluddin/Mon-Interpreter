package main

import (
	"Mon/repl"
	"fmt"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()

	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is the Monkey Programming language!\n", user.Username)
	fmt.Printf("Feel free to type in command!\n")

	repl.Start(os.Stdin, os.Stdout)
}
