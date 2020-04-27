package golox

import (
	"errors"
)

type ParserStep struct {
	Exprs []Expr
	Logs  []string
}

type parser struct {
	tokens          []Token
	current         int
	expressionCount int
	exprs           []Expr
	steps           []ParserStep
	logs            []string
}

func (p *parser) parse() Expr {
	p.current = 0
	p.expressionCount = 0
	p.expression()
	return p.exprs[0]
}

func (p *parser) parseForSteps() []ParserStep {
	p.current = 0
	p.expressionCount = 0
	p.steps = append(p.steps, ParserStep{Exprs: copyExprs(p.exprs), Logs: copyLogs(p.logs)})
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
func copyLogs(logs []string) []string {
	ls := make([]string, len(logs))
	copy(ls, logs)
	return ls
}

func (p *parser) addLog(log string) {
	p.logs = append(p.logs, log)
	p.steps = append(p.steps, ParserStep{Exprs: copyExprs(p.exprs), Logs: copyLogs(p.logs)})
}

func (p *parser) popLog() {
	newSize := len(p.logs) - 1
	p.logs = p.logs[:newSize]
	p.steps = append(p.steps, ParserStep{Exprs: copyExprs(p.exprs), Logs: copyLogs(p.logs)})
}

func (p *parser) addExpr(expr Expr) {
	if len(p.exprs) > 0 {
		exprToUpdate := p.exprs[len(p.exprs)-1]
		exprToUpdate.UpdateChildExpr(expr)
	}
	p.exprs = append(p.exprs, expr)
	p.steps = append(p.steps, ParserStep{Exprs: copyExprs(p.exprs), Logs: copyLogs(p.logs)})
}

func (p *parser) getExpr() Expr {
	expr := p.exprs[len(p.exprs)-1]
	return expr
}

func (p *parser) popExpr() Expr {
	newSize := len(p.exprs) - 1
	expr := p.exprs[len(p.exprs)-1]
	p.exprs = p.exprs[:newSize]
	if len(p.exprs) > 0 {
		p.steps = append(p.steps, ParserStep{Exprs: copyExprs(p.exprs), Logs: copyLogs(p.logs)})
	}
	return expr
}

func (p *parser) expression() {
	p.addLog("Searching for expresssion")
	p.equality()
	p.popLog()
}

func (p *parser) equality() {
	p.addLog("Searching for equality or higher")
	p.comparison()

	for p.match([]TokenType{BangEqual, EqualEqual}) {
		operator := p.previous()
		right := unknownExpr{p.exprCount()}
		p.addExpr(&binaryExpr{p.popExpr(), operator, &right, p.exprCount()})
		p.comparison()
		p.popExpr()
	}
	p.popLog()
}

func (p *parser) comparison() {
	p.addLog("Searching for comparison or higher")
	p.addition()

	for p.match([]TokenType{Greater, GreaterEqual, Less, LessEqual}) {
		operator := p.previous()
		right := unknownExpr{p.exprCount()}
		p.addExpr(&binaryExpr{p.popExpr(), operator, &right, p.exprCount()})
		p.addition()
		p.popExpr()
	}
	p.popLog()

}

func (p *parser) addition() {
	p.addLog("Searching for addition or higher")

	p.multiplication()

	for p.match([]TokenType{Minus, Plus}) {
		operator := p.previous()
		// For the visualization I want the parent to appear before the unknown value,
		// so tweak the orders to make it look that way
		right := unknownExpr{p.exprCount()}
		p.addExpr(&binaryExpr{p.popExpr(), operator, &right, p.exprCount()})
		p.multiplication()
		p.popExpr()
	}
	p.popLog()

}

func (p *parser) multiplication() {
	p.addLog("Searching for multiplication or higher")

	p.unary()
	for p.match([]TokenType{Slash, Star}) {
		operator := p.previous()
		right := unknownExpr{p.exprCount()}
		p.addExpr(&binaryExpr{p.popExpr(), operator, &right, p.exprCount()})
		p.unary()
		p.popExpr()
	}
	p.popLog()

}

// 2
// 2+
// 2+3 3
// 2+* 3*?
// 2+* 3*4 4
// 2+* 3*4

func (p *parser) unary() {
	p.addLog("Searching for unary or higher")

	if p.match([]TokenType{Bang, Minus}) {
		operator := p.previous()
		right := unknownExpr{p.exprCount()}
		p.addExpr(&unaryExpr{operator, &right, p.exprCount()})
		p.unary()
		p.popExpr()
		return
	}
	err := p.primary()
	if err != nil {
		// Do something
	}
	p.popLog()

}

func (p *parser) primary() error {
	p.addLog("Searching for primary")

	if p.match([]TokenType{FalseKeyword}) {
		p.addExpr(&literalExpr{false, p.exprCount(), p.previous()})
		p.popLog()

		return nil
	}
	if p.match([]TokenType{TrueKeyword}) {
		p.addExpr(&literalExpr{true, p.exprCount(), p.previous()})
		p.popLog()

		return nil
	}
	if p.match([]TokenType{NilKeyword}) {
		p.addExpr(&literalExpr{nil, p.exprCount(), p.previous()})
		p.popLog()

		return nil
	}
	if p.match([]TokenType{Number, StringLiteral}) {
		p.addExpr(&literalExpr{p.previous().Literal, p.exprCount(), p.previous()})
		p.popLog()

		return nil
	}
	if p.match([]TokenType{LeftParen}) {
		expr := unknownExpr{p.exprCount()}
		p.addExpr(&groupingExpr{&expr, p.exprCount(), p.previous()})
		p.expression()
		p.popExpr()
		_, err := p.consume(RightParen, "Expected matching ')'")
		if err != nil {
			// Do nothing for now
		}
		p.popLog()

		return nil
	}
	parseError(p.peek(), "Expected expression")
	p.popLog()

	return errors.New("Expected expression!")
}

func (p *parser) exprCount() int {
	p.expressionCount++
	return p.expressionCount
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
