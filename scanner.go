package golox

import (
	"errors"
	"fmt"
	"strconv"
)

type ScannerStep struct {
	Tokens  []Token
	Current int
	Start   int
	Line    int
}

type scanner struct {
	source         string
	tokens         []Token
	start          int
	current        int
	line           int
	lineStart      int
	calculateSteps bool
	steps          []ScannerStep
}

// displayError is a callback to show any errors found during scanning
func (s *scanner) scanTokens(displayError func(string)) ([]Token, error) {
	s.start = 0
	s.current = 0
	s.lineStart = 0
	s.line = 1
	hadError := false
	for !s.isAtEnd() {
		s.start = s.current
		err := s.scanToken(displayError)
		if err != nil {
			hadError = true
		}
	}

	s.start = s.current
	s.addTokenWithLiteral(Eof, "")
	if hadError {
		return s.tokens, errors.New("Error during scanning")
	}
	return s.tokens, nil
}

// displayError is a callback to show any errors found during scanning
func (s *scanner) scanTokensForSteps(displayError func(string)) ([]ScannerStep, error) {
	s.start = 0
	s.current = 0
	s.lineStart = 0
	s.line = 1
	s.calculateSteps = true
	hadError := false
	for !s.isAtEnd() {
		s.start = s.current
		s.addStep()
		err := s.scanToken(displayError)
		if err != nil {
			hadError = true
		}
	}

	s.start = s.current
	s.addTokenWithLiteral(Eof, "")
	if hadError {
		return s.steps, errors.New("Error during scanning")
	}
	return s.steps, nil
}

func (s *scanner) isAtEnd() bool {
	return s.current >= len(s.source)
}

func (s *scanner) scanToken(displayError func(string)) error {
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
	case "-":
		s.addToken(Minus)
	case ".":
		s.addToken(Dot)
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
		err := s.handleString(displayError)
		if err != nil {
			return err
		}
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
		s.incrementLine()

	default:
		if isDigit(c) {
			err := s.handleNumber(displayError)
			if err != nil {
				return err
			}
		} else if isAlpha(c) {
			s.handleIdentifier()
		} else {
			errorMsg := fmt.Sprintf("Unexpected character '%s' on line %d", c, s.line)
			displayError(errorMsg)
			return errors.New(errorMsg)
		}
	}
	return nil
}

func (s *scanner) match(expected string) bool {
	if s.isAtEnd() {
		return false
	}
	if string(s.source[s.current]) != expected {
		return false
	}
	s.current++
	s.addStep()
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

func (s *scanner) handleNumber(displayError func(string)) error {
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
		errorMsg := fmt.Sprintf("Couldn't parse number on line %d", s.line)
		displayError(errorMsg)
		return errors.New(errorMsg)
	}
	s.addTokenWithLiteral(Number, num)
	return nil
}

func (s *scanner) handleString(displayError func(string)) error {
	for s.peek() != "\"" && !s.isAtEnd() {
		if s.peek() == "\n" {
			s.incrementLine()
		}
		s.advance()
	}
	if s.isAtEnd() {
		errorMsg := fmt.Sprintf("Unterminated string on line %d", s.line)
		displayError(errorMsg)
		return errors.New(errorMsg)
	}
	s.advance()
	s.addTokenWithLiteral(StringLiteral, s.source[s.start+1:s.current-1])
	return nil
}

func (s *scanner) incrementLine() {
	s.line++
	s.lineStart = s.current
}

func (s *scanner) advance() string {
	s.current++
	s.addStep()
	return string(s.source[s.current-1])
}

func (s *scanner) addToken(ttype TokenType) {
	s.addTokenWithLiteral(ttype, nil)
}

func (s *scanner) addTokenWithLiteral(ttype TokenType, literal interface{}) {
	text := s.source[s.start:s.current]
	token := Token{ttype, text, literal, s.line, s.start - s.lineStart, s.current - s.lineStart}
	s.tokens = append(s.tokens, token)
	s.addStep()
}

func (s *scanner) addStep() {
	if s.calculateSteps {
		step := ScannerStep{Tokens: s.tokens, Current: s.current - s.lineStart, Start: s.start - s.lineStart, Line: s.line}
		s.steps = append(s.steps, step)
	}
}
