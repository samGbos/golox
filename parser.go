package golox

import "errors"

type parser struct {
	tokens  []Token
	current int
}

func (p *parser) parse() expr {
	p.current = 0
	return p.expression()
}

func (p *parser) expression() expr {
	return p.equality()
}

func (p *parser) equality() expr {
	e := p.comparison()

	for p.match([]TokenType{BangEqual, EqualEqual}) {
		operator := p.previous()
		right := p.comparison()
		e = binaryExpr{e, operator, right}
	}

	return e
}

func (p *parser) comparison() expr {
	e := p.addition()
	for p.match([]TokenType{Greater, GreaterEqual, Less, LessEqual}) {

		operator := p.previous()
		right := p.addition()
		e = binaryExpr{e, operator, right}
	}

	return e
}

func (p *parser) addition() expr {
	e := p.multiplication()
	for p.match([]TokenType{Minus, Plus}) {
		operator := p.previous()
		right := p.multiplication()
		e = binaryExpr{e, operator, right}
	}

	return e
}

func (p *parser) multiplication() expr {
	e := p.unary()

	for p.match([]TokenType{Slash, Star}) {
		operator := p.previous()
		right := p.unary()
		e = binaryExpr{e, operator, right}
	}

	return e
}

func (p *parser) unary() expr {
	if p.match([]TokenType{Bang, Minus}) {
		operator := p.previous()
		right := p.unary()
		return unaryExpr{operator, right}
	}
	e, err := p.primary()
	if err != nil {
		// Do something
	}
	return e
}

func (p *parser) primary() (expr, error) {
	if p.match([]TokenType{FalseKeyword}) {
		return literalExpr{false}, nil
	}
	if p.match([]TokenType{TrueKeyword}) {
		return literalExpr{true}, nil
	}
	if p.match([]TokenType{NilKeyword}) {
		return literalExpr{nil}, nil
	}
	if p.match([]TokenType{Number, StringLiteral}) {
		return literalExpr{p.previous().Literal}, nil
	}
	if p.match([]TokenType{LeftParen}) {
		e := p.expression()
		_, err := p.consume(RightParen, "Expected matching ')'")
		if err != nil {
			// Do nothing for now
		}
		return groupingExpr{e}, nil
	}
	parseError(p.peek(), "Expected expression")
	return nil, errors.New("Expected expression!")
}

func (p *parser) consume(ttype TokenType, message string) (Token, error) {
	if p.check(ttype) {
		return p.advance(), nil
	}
	tok := p.peek()
	parseError(tok, message)
	return p.peek(), errors.New(message)
}

func (p *parser) synchronize() {
	p.advance()

	for !p.isAtEnd() {
		if p.previous().Ttype == Semicolon {
			return
		}
		switch p.peek().Ttype {
		case ClassKeyword:
			return
		case FunKeyword:
			return
		case VarKeyword:
			return
		case ForKeyword:
			return
		case IfKeyword:
			return
		case WhileKeyword:
			return
		case PrintKeyword:
			return
		case ReturnKeyword:
			return
		}
		p.advance()
	}
}

func (p *parser) match(ttypes []TokenType) bool {
	for _, ttype := range ttypes {
		if p.check(ttype) {
			p.advance()
			return true
		}
	}
	return false
}

func (p *parser) check(ttype TokenType) bool {
	if p.isAtEnd() {
		return false
	}
	return p.peek().Ttype == ttype
}

func (p *parser) advance() Token {
	if !p.isAtEnd() {
		p.current++
	}
	return p.previous()
}

func (p *parser) previous() Token {
	return p.tokens[p.current-1]
}

func (p *parser) peek() Token {
	return p.tokens[p.current]
}
func (p *parser) isAtEnd() bool {
	return p.peek().Ttype == Eof
}
