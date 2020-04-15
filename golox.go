package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

var hadError bool = false

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
		run(text)
		hadError = true
	}
}

func runFile(script string) {
	b, err := ioutil.ReadFile(script)
	if err != nil {
		fmt.Print(err)
	}
	run(string(b))
	if hadError {
		os.Exit(65)
	}
}

func reportError(line int, message string) {
	report(line, "", message)
	hadError = true
}

func report(line int, where string, message string) {
	fmt.Println("Error on line ", line, ":", where, " -- ", message)
}

func run(source string) {
	scanner := Scanner{source: source}
	tokens := scanner.scanTokens()
	for _, token := range tokens {
		fmt.Printf("%#v\n", token)
	}
	fmt.Print(source)
	fmt.Println("hadError", hadError)
}
