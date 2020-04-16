package golox

import (
	"fmt"
	"strconv"
)

type scanner struct {
	source  string
	tokens  []Token
	start   int
	current int
	line    int
}

func (s *scanner) scanTokens() []Token {
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

func (s *scanner) isAtEnd() bool {
	return s.current >= len(s.source)
}

func (s *scanner) scanToken() {
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
	case "\"":
		s.handleString()
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
		s.line++

	default:
		if isDigit(c) {
			s.handleNumber()
		} else if isAlpha(c) {
			s.handleIdentifier()
		} else {
			reportError(s.line, fmt.Sprintf("Unexpected character %s", c))
		}
	}
}

func (s *scanner) match(expected string) bool {
	if s.isAtEnd() {
		return false
	}
	if string(s.source[s.current]) != expected {
		return false
	}
	s.current++
	return true
}

func (s *scanner) peek() string {
	if s.isAtEnd() {
		return "\000"
	}
	return string(s.source[s.current])
}

func (s *scanner) peekNext() string {
	if s.current+1 >= len(s.source) {
		return "\000"
	}
	return string(s.source[s.current+1])
}

func isAlpha(c string) bool {
	return (c >= "a" && c <= "z") ||
		(c >= "A" && c <= "Z") ||
		c == "_"
}

func isDigit(c string) bool {
	return c >= "0" && c <= "9"
}

func isAlphanumeric(c string) bool {
	return isAlpha(c) || isDigit(c)
}

func (s *scanner) handleIdentifier() {
	for isAlphanumeric(s.peek()) {
		s.advance()
	}

	text := s.source[s.start:s.current]
	ttype, ok := reserved_words[text]
	if ok {
		s.addToken(ttype)
	} else {
		s.addToken(Identifier)
	}
}

func (s *scanner) handleNumber() {
	for isDigit(s.peek()) {
		s.advance()
	}
	if s.peek() == "." && isDigit(s.peekNext()) {
		s.advance() // consume the "."

		for isDigit(s.peek()) {
			s.advance()
		}
	}
	num, err := strconv.ParseFloat(s.source[s.start:s.current], 64)
	if err != nil {
		reportError(s.line, "Couldn't parse number")
	}
	s.addTokenWithLiteral(Number, num)
}

func (s *scanner) handleString() {
	for s.peek() != "\"" && !s.isAtEnd() {
		if s.peek() == "\n" {
			s.line++
		}
		s.advance()
	}
	if s.isAtEnd() {
		reportError(s.line, "Unterminated string")
		return
	}
	s.advance()
	s.addTokenWithLiteral(StringLiteral, s.source[s.start+1:s.current-1])
}

func (s *scanner) advance() string {
	s.current = s.current + 1
	return string(s.source[s.current-1])
}

func (s *scanner) addToken(ttype TokenType) {
	s.addTokenWithLiteral(ttype, nil)
}

func (s *scanner) addTokenWithLiteral(ttype TokenType, literal interface{}) {
	text := s.source[s.start:s.current]
	token := Token{ttype, text, literal, s.line}
	s.tokens = append(s.tokens, token)
}
