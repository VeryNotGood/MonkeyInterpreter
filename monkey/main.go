package main

import (
	"fmt"
	"log"
	"monkey/repl"
	"os"
	"os/user"
)

func main() {

	user, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Hello %s! Welcome to Monkey\n", user.Username)
	repl.ReplStart(os.Stdin, os.Stdout)
}
