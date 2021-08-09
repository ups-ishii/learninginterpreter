package main

import (
	"fmt"
	"log"
	"os"
	"os/user"

	"github.com/ups-ishii/learninginterpreter/monkey/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Hello %s! This is the Monkey programming language!\n", user.Username)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
