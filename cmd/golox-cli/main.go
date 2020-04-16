// +build !js

package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/samGbos/golox"
)

func main() {
	args := os.Args
	if len(args) > 2 {
		fmt.Println("Usage golox [script]")
		os.Exit(64)
	} else if len(args) == 2 {
		runFile(args[1])
	} else {
		runPrompt()
	}
}

func runPrompt() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		text, _ := reader.ReadString('\n')
		expr := golox.RunParser(text)
		fmt.Println(expr)
	}
}

func runFile(script string) {
	b, err := ioutil.ReadFile(script)
	if err != nil {
		fmt.Print(err)
	}
	tokens := golox.RunScanner(string(b))
	for _, token := range tokens {
		fmt.Printf("%#v\n", token)
	}
	// if hadError {
	// os.Exit(65)
	// }
}
