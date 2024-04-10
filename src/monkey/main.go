package main

import (
	"fmt"
	"monkey/repl"
	"os"
	"os/user"
)

func main() {
	var mode string
	if len(os.Args) == 1 {
		fmt.Println("Use default interpreter mode.")
		mode = "interpreter"
	} else {
		mode = os.Args[1]
	}

	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the Monkey programming language!\n",
		user.Username)
	fmt.Printf("Feel free to type in commands\n")

	if mode == "interpreter" {
		fmt.Println("Interpreter Mode!")
		repl.Start(os.Stdin, os.Stdout)
	} else if mode == "vm" {
		fmt.Println("Compiler Mode!")
		repl.StartVM(os.Stdin, os.Stdout)
	}
}
