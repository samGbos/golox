package golox

import (
	"fmt"
)

var hadError bool = false

func parseError(tok token, message string) {
	if tok.ttype == eof {
		report(tok.line, "at end", message)
	} else {
		report(tok.line, fmt.Sprintf("at '%s'", tok.lexeme), message)
	}
}

func reportError(line int, message string) {
	report(line, "", message)
	hadError = true
}

func report(line int, where string, message string) {
	fmt.Println("Error on line ", line, ":", where, " -- ", message)
}

func RunScanner(source string) []token {
    s := scanner{source: source}
	return s.scanTokens()
}

func Run(source string) bool {
    tokens := RunScanner(source)
	p := parser{tokens: tokens}
	e := p.parse()
	fmt.Print(e)
	return hadError
}
