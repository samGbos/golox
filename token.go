package main

type Token struct {
	ttype   TokenType
	lexeme  string
	literal string
	line    int
}
