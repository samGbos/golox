package golox

type Expr interface {
	Name() interface{}
	Children() []Expr
	UpdateChildExpr(Expr)
	Copy() Expr
}

type unknownExpr struct{}

func (expr unknownExpr) Name() interface{} {
	return "??"
}

func (expr unknownExpr) Children() []Expr {
	return nil
}

func (expr *unknownExpr) UpdateChildExpr(child Expr) {
	// do nothing
}

func (expr *unknownExpr) Copy() Expr {
	return &unknownExpr{}
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

func (expr *binaryExpr) UpdateChildExpr(child Expr) {
	expr.right = child
}

func (expr *binaryExpr) Copy() Expr {
	return &binaryExpr{expr.left.Copy(), expr.operator, expr.right.Copy()}
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

func (expr *unaryExpr) UpdateChildExpr(child Expr) {
	expr.right = child
}

func (expr *unaryExpr) Copy() Expr {
	return &unaryExpr{expr.operator, expr.right.Copy()}
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

func (expr literalExpr) UpdateChildExpr(child Expr) {
	// do nothing
}

func (expr *literalExpr) Copy() Expr {
	return &literalExpr{expr.value}
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

func (expr *groupingExpr) UpdateChildExpr(child Expr) {
	// do nothing
	expr.expression = child
}

func (expr *groupingExpr) Copy() Expr {
	return &groupingExpr{expr.expression.Copy()}
}
