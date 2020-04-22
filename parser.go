package golox

import (
	"errors"
	"fmt"
)

type ParserStep struct {
	Exprs []Expr
}

type parser struct {
	tokens  []Token
	current int
	exprs   []Expr
	steps   []ParserStep
}

func (p *parser) parse() Expr {
	p.current = 0
	p.expression()
	return p.exprs[0]
}

func (p *parser) parseForSteps() []ParserStep {
	p.current = 0
	p.expression()
	return p.steps
}

func copyExprs(exprs []Expr) []Expr {
	es := make([]Expr, len(exprs))
	for i, e := range exprs {
		es[i] = e.Copy()
	}
	return es
}

func (p *parser) addExpr(expr Expr) {
	if len(p.exprs) > 0 {
		exprToUpdate := p.exprs[len(p.exprs)-1]
		fmt.Println("Update ", exprToUpdate.Name(), " to ", expr.Name())
		exprToUpdate.UpdateChildExpr(expr)
	}
	p.exprs = append(p.exprs, expr)
	fmt.Println("Push ", expr.Name())
	p.steps = append(p.steps, ParserStep{Exprs: copyExprs(p.exprs)})
}

func (p *parser) getExpr() Expr {
	expr := p.exprs[len(p.exprs)-1]
	return expr
}

func (p *parser) popExpr() Expr {
	newSize := len(p.exprs) - 1
	expr := p.exprs[len(p.exprs)-1]
	p.exprs = p.exprs[:newSize]
	fmt.Println("Pop ", expr.Name())
	p.steps = append(p.steps, ParserStep{Exprs: copyExprs(p.exprs)})
	return expr
}

func (p *parser) expression() {
	p.equality()
}

func (p *parser) equality() {
	p.comparison()

	for p.match([]TokenType{BangEqual, EqualEqual}) {
		operator := p.previous()
		right := unknownExpr{}
		p.addExpr(&binaryExpr{p.popExpr(), operator, &right})
		p.comparison()
		p.popExpr()
	}
}

func (p *parser) comparison() {
	p.addition()

	for p.match([]TokenType{Greater, GreaterEqual, Less, LessEqual}) {
		operator := p.previous()
		right := unknownExpr{}
		p.addExpr(&binaryExpr{p.popExpr(), operator, &right})
		p.addition()
		p.popExpr()
	}
}

func (p *parser) addition() {
	p.multiplication()

	for p.match([]TokenType{Minus, Plus}) {
		operator := p.previous()
		right := unknownExpr{}

		p.addExpr(&binaryExpr{p.popExpr(), operator, &right})
		p.multiplication()
		p.popExpr()
	}
}

func (p *parser) multiplication() {
	p.unary()
	for p.match([]TokenType{Slash, Star}) {
		operator := p.previous()
		right := unknownExpr{}

		p.addExpr(&binaryExpr{p.popExpr(), operator, &right})
		p.unary()
		p.popExpr()
	}

}

// 2
// 2+
// 2+3 3
// 2+* 3*?
// 2+* 3*4 4
// 2+* 3*4

func (p *parser) unary() {
	if p.match([]TokenType{Bang, Minus}) {
		operator := p.previous()
		right := unknownExpr{}
		p.addExpr(&unaryExpr{operator, &right})
		p.unary()
		p.popExpr()
		return
	}
	err := p.primary()
	if err != nil {
		// Do something
	}
}

func (p *parser) primary() error {
	if p.match([]TokenType{FalseKeyword}) {
		p.addExpr(&literalExpr{false})
		return nil
	}
	if p.match([]TokenType{TrueKeyword}) {
		p.addExpr(&literalExpr{true})
		return nil
	}
	if p.match([]TokenType{NilKeyword}) {
		p.addExpr(&literalExpr{nil})
		return nil
	}
	if p.match([]TokenType{Number, StringLiteral}) {
		p.addExpr(&literalExpr{p.previous().Literal})
		return nil
	}
	if p.match([]TokenType{LeftParen}) {
		expr := unknownExpr{}
		p.addExpr(&groupingExpr{&expr})
		p.expression()
		p.popExpr()
		_, err := p.consume(RightParen, "Expected matching ')'")
		if err != nil {
			// Do nothing for now
		}
		return nil
	}
	parseError(p.peek(), "Expected expression")
	return errors.New("Expected expression!")
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
