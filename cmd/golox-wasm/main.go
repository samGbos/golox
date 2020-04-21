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
	js.Global().Set("runParser", js.FuncOf(runParser))
	<-c
}

func convertStep(step golox.Step) map[string]interface{} {
	tokens := make([]interface{}, len(step.Tokens))
	for itok, token := range step.Tokens {
		tokens[itok] = convertToken(token)
	}
	return map[string]interface{}{
		"tokens": tokens,
		"current": step.Current,
		"start": step.Start,
		"line": step.Line,
	}
}

func convertToken(t golox.Token) map[string]interface{} {
	return map[string]interface{}{
		"token_type": t.Ttype.String(),
		"lexeme":     t.Lexeme,
		"literal":    t.Literal,
		"line":       t.Line,
		"start":      t.Start,
		"end":        t.End,
	}
}

func convertExpr(expr golox.Expr) map[string]interface{} {
	exprChildren := expr.Children()
	children := make([]interface{}, len(exprChildren))
	for idx, child := range exprChildren {
		children[idx] = convertExpr(child)
	}
	fmt.Println(expr.Name())
	fmt.Println(children)
	return map[string]interface{}{
		"name":     expr.Name(),
		"children": children,
	}
}

func runScanner(this js.Value, inputs []js.Value) interface{} {
	message := inputs[0].String()
	errorHandler := inputs[1]

    displayError := func(errorMsg string) {
        errorHandler.Invoke(errorMsg)
    }

	steps := golox.RunScannerForSteps(message, displayError)
	serializedSteps := make([]interface{}, len(steps))
	for istep, step := range steps {
		serializedSteps[istep] = convertStep(step)
	}

	jsVal := map[string]interface{}{
		"steps": serializedSteps,
	}
	return jsVal
}

func runParser(this js.Value, inputs []js.Value) interface{} {
	message := inputs[0].String()
	errorHandler := inputs[1]

    displayError := func(errorMsg string) {
        errorHandler.Invoke(errorMsg)
    }

	expr := golox.RunParser(message, displayError)

	jsVal := map[string]interface{}{
		"expr": convertExpr(expr),
	}
	return jsVal
}
