package golox

type expr interface {
}

type binaryExpr struct {
	left     expr
	operator token
	right    expr
}

type unaryExpr struct {
	operator token
	right    expr
}

type literalExpr struct {
	value interface{}
}

type groupingExpr struct {
	expression expr
}
