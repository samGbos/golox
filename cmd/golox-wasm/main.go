// +build js,wasm

package main

import (
	"fmt"
	"syscall/js"

	"github.com/samGbos/golox"
)

func main() {
	fmt.Println("Hello, WebAssembly!")
	c := make(chan bool)
	js.Global().Set("runScanner", js.FuncOf(runScanner))
	<-c
}

func convert(t golox.Token) map[string]interface{} {
	return map[string]interface{}{
		"token_type": t.Ttype.String(),
		"lexeme":     t.Lexeme,
		"literal":    t.Literal,
		"line":       t.Line,
	}
}

func runScanner(this js.Value, inputs []js.Value) interface{} {
	message := inputs[0].String()
	callback := inputs[1]

	tokens := golox.RunScanner(message)
	toks := make([]interface{}, len(tokens))
	for i, tok := range tokens {
		toks[i] = convert(tok)
	}

	fmt.Println(toks)

	jsVal := map[string]interface{}{
		"tokens": toks,
	}
	callback.Invoke(jsVal)
	return jsVal
}
