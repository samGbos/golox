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
		hadError := golox.Run(text)
		if hadError {
		    fmt.Print("Error ocurred")
		}
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
