package main

import (
	"fmt"
	"go-monkey/interpreter/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s!, this is Monekey-Go ", user.Username)
	fmt.Println("Try this ")
	repl.Start(os.Stdin, os.Stdout)
}
