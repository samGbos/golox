package golox

import "errors"

type parser struct {
	tokens  []token
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

	for p.match([]tokenType{bangEqual, equalEqual}) {
		operator := p.previous()
		right := p.comparison()
		e = binaryExpr{e, operator, right}
	}

	return e
}

func (p *parser) comparison() expr {
	e := p.addition()
	for p.match([]tokenType{greater, greaterEqual, less, lessEqual}) {

		operator := p.previous()
		right := p.addition()
		e = binaryExpr{e, operator, right}
	}

	return e
}

func (p *parser) addition() expr {
	e := p.multiplication()
	for p.match([]tokenType{minus, plus}) {
		operator := p.previous()
		right := p.multiplication()
		e = binaryExpr{e, operator, right}
	}

	return e
}

func (p *parser) multiplication() expr {
	e := p.unary()

	for p.match([]tokenType{slash, star}) {
		operator := p.previous()
		right := p.unary()
		e = binaryExpr{e, operator, right}
	}

	return e
}

func (p *parser) unary() expr {
	if p.match([]tokenType{bang, minus}) {
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
	if p.match([]tokenType{falseKeyword}) {
		return literalExpr{false}, nil
	}
	if p.match([]tokenType{trueKeyword}) {
		return literalExpr{true}, nil
	}
	if p.match([]tokenType{nilKeyword}) {
		return literalExpr{nil}, nil
	}
	if p.match([]tokenType{number, stringLiteral}) {
		return literalExpr{p.previous().literal}, nil
	}
	if p.match([]tokenType{leftParen}) {
		e := p.expression()
		_, err := p.consume(rightParen, "Expected matching ')'")
		if err != nil {
			// Do nothing for now
		}
		return groupingExpr{e}, nil
	}
	parseError(p.peek(), "Expected expression")
	return nil, errors.New("Expected expression!")
}

func (p *parser) consume(ttype tokenType, message string) (token, error) {
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
		if p.previous().ttype == semicolon {
			return
		}
		switch p.peek().ttype {
		case classKeyword:
			return
		case funKeyword:
			return
		case varKeyword:
			return
		case forKeyword:
			return
		case ifKeyword:
			return
		case whileKeyword:
			return
		case printKeyword:
			return
		case returnKeyword:
			return
		}
		p.advance()
	}
}

func (p *parser) match(ttypes []tokenType) bool {
	for _, ttype := range ttypes {
		if p.check(ttype) {
			p.advance()
			return true
		}
	}
	return false
}

func (p *parser) check(ttype tokenType) bool {
	if p.isAtEnd() {
		return false
	}
	return p.peek().ttype == ttype
}

func (p *parser) advance() token {
	if !p.isAtEnd() {
		p.current++
	}
	return p.previous()
}

func (p *parser) previous() token {
	return p.tokens[p.current-1]
}

func (p *parser) peek() token {
	return p.tokens[p.current]
}
func (p *parser) isAtEnd() bool {
	return p.peek().ttype == eof
}
