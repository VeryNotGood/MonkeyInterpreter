package main

import (
	"fmt"
	"log"
	"os"
	"os/user"

	"github.com/VeryNotGood/monkey/repl"
)

func main() {

	user, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Hello %s! Welcome to Monkey\n", user.Username)
	repl.ReplStart(os.Stdin, os.Stdout)
}
