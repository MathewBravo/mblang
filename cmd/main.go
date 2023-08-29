package main

import (
	"fmt"
	"os"

	"github.com/MathewBravo/mblang/pkg/repl"
)

func main() {
	fmt.Println("Welcome to the MBLang Programming Language")
	fmt.Println("Feel free to type in commands")
	repl.Start(os.Stdin, os.Stdout)
}
