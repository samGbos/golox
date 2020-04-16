package golox

type Expr interface {
    Name() interface{}
    Children() []Expr
}

type binaryExpr struct {
	left     Expr
	operator Token
	right    Expr
}

func (expr binaryExpr) Name() interface{} {
    return expr.operator.Lexeme
}

func (expr binaryExpr) Children() []Expr {
    return []Expr{expr.left, expr.right}
}


type unaryExpr struct {
	operator Token
	right    Expr
}

func (expr unaryExpr) Name() interface{} {
    return expr.operator.Lexeme
}

func (expr unaryExpr) Children() []Expr {
    return []Expr{expr.right}
}


type literalExpr struct {
	value interface{}
}

func (expr literalExpr) Name() interface{} {
    return expr.value
}

func (expr literalExpr) Children() []Expr {
    return nil
}


type groupingExpr struct {
	expression Expr
}

func (expr groupingExpr) Name() interface{} {
    return "()"
}

func (expr groupingExpr) Children() []Expr {
    return []Expr{expr.expression}
}