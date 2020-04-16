package golox

import (
	"fmt"
)

var hadError bool = false

func parseError(token Token, message string) {
	if token.Ttype == Eof {
		report(token.Line, "at end", message)
	} else {
		report(token.Line, fmt.Sprintf("at '%s'", token.Lexeme), message)
	}
}

func reportError(line int, message string) {
	report(line, "", message)
	hadError = true
}

func report(line int, where string, message string) {
	fmt.Println("Error on line ", line, ":", where, " -- ", message)
}

func RunScanner(source string) []Token {
	s := scanner{source: source}
	return s.scanTokens()
}

func RunParser(source string) Expr {
	tokens := RunScanner(source)
	p := parser{tokens: tokens}
	expr := p.parse()
	return expr
}
