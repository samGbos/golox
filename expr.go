package golox

type expr interface {
}

type binaryExpr struct {
	left     expr
	operator Token
	right    expr
}

type unaryExpr struct {
	operator Token
	right    expr
}

type literalExpr struct {
	value interface{}
}

type groupingExpr struct {
	expression expr
}
