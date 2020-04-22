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

func RunScanner(source string, displayError func(string)) []Token {
	s := scanner{source: source}
	tokens, err := s.scanTokens(displayError)
	if err != nil {
		// do nothing
	}
	return tokens
}

func RunScannerForSteps(source string, displayError func(string)) []ScannerStep {
	s := scanner{source: source}
	steps, err := s.scanTokensForSteps(displayError)
	if err != nil {
		// do nothing
	}
	return steps
}

func RunParser(source string, displayError func(string)) Expr {
	tokens := RunScanner(source, displayError)
	p := parser{tokens: tokens}
	return p.parse()
}

func RunParserForSteps(source string, displayError func(string)) []ParserStep {
	tokens := RunScanner(source, displayError)
	p := parser{tokens: tokens}
	return p.parseForSteps()
}
