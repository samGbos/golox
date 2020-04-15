package main

import "fmt"

type Scanner struct {
	source  string
	tokens  []Token
	start   int
	current int
	line    int
}

func (s *Scanner) scanTokens() []Token {
	s.start = 0
	s.current = 0
	s.line = 1
	for !s.isAtEnd() {
		s.start = s.current
		s.scanToken()
	}

	token := Token{Eof, "", "", s.line}
	s.tokens = append(s.tokens, token)
	return s.tokens
}

func (s *Scanner) isAtEnd() bool {
	return s.current >= len(s.source)
}

func (s *Scanner) scanToken() {
	c := s.advance()
	switch c {
	case "(":
		s.addToken(LeftParen)
	case ")":
		s.addToken(RightParen)
	case "{":
		s.addToken(LeftBrace)
	case ",":
		s.addToken(Comma)
	case ".":
		s.addToken(Minus)
	case "+":
		s.addToken(Plus)
	case ";":
		s.addToken(Semicolon)
	case "*":
		s.addToken(Star)
	case "!":
		var t TokenType
		if s.match("=") {
			t = BangEqual
		} else {
			t = Bang
		}
		s.addToken(t)
	case "=":
		var t TokenType
		if s.match("=") {
			t = EqualEqual
		} else {
			t = Equal
		}
		s.addToken(t)
	case "<":
		var t TokenType
		if s.match("=") {
			t = LessEqual
		} else {
			t = Less
		}
		s.addToken(t)
	case ">":
		var t TokenType
		if s.match("=") {
			t = GreaterEqual
		} else {
			t = Greater
		}
		s.addToken(t)
	case "/":
		if s.match("/") {
			for s.peek() != "\n" && !s.isAtEnd() {
				s.advance()
			}
		} else {
			s.addToken(Slash)

		}
	case " ":
	case "\r":
	case "\t":
	case "\n":
		s.line += 1

	default:
		reportError(s.line, fmt.Sprintf("Unexpected character %s", c))
	}
}

func (s *Scanner) match(expected string) bool {
	if s.isAtEnd() {
		return false
	}
	if string(s.source[s.current]) != expected {
		return false
	}
	s.current += 1
	return true
}

func (s *Scanner) peek() string {
	if s.isAtEnd() {
		return "\000"
	}
	return string(s.source[s.current])
}

func (s *Scanner) advance() string {
	s.current = s.current + 1
	return string(s.source[s.current-1])
}

func (s *Scanner) addToken(ttype TokenType) {
	s.addTokenWithLiteral(ttype, "")
}

func (s *Scanner) addTokenWithLiteral(ttype TokenType, literal string) {
	text := s.source[s.start:s.current]
	token := Token{ttype, text, literal, s.line}
	s.tokens = append(s.tokens, token)
}
