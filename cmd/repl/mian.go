package main

import (
	"fmt"
	"github.com/watertreestar/go-monkey/repl"
	"os"
	"os/user"
)

func main() {
	u,err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is Monkey programming language!\n",u.Username)
	fmt.Printf("Feel free to type in command\n")
	repl.Start(os.Stdin,os.Stdout)
}
