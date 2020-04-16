package golox

var reserved_words = map[string]tokenType{
	"and":    andKeyword,
	"class":  classKeyword,
	"else":   elseKeyword,
	"false":  falseKeyword,
	"fun":    funKeyword,
	"for":    forKeyword,
	"if":     ifKeyword,
	"nil":    nilKeyword,
	"or":     orKeyword,
	"print":  printKeyword,
	"return": returnKeyword,
	"super":  superKeyword,
	"this":   thisKeyword,
	"true":   trueKeyword,
	"var":    varKeyword,
	"while":  whileKeyword,
}

type tokenType int

const (
	// Single char tokens
	leftParen  tokenType = 0
	rightParen tokenType = 1
	leftBrace  tokenType = 2
	rightBrace tokenType = 3
	comma      tokenType = 4
	dot        tokenType = 5
	minus      tokenType = 6
	plus       tokenType = 7
	semicolon  tokenType = 8
	slash      tokenType = 9
	star       tokenType = 10

	// One/two char tokens
	bang         tokenType = 11
	bangEqual    tokenType = 12
	equal        tokenType = 13
	equalEqual   tokenType = 14
	greater      tokenType = 15
	greaterEqual tokenType = 16
	less         tokenType = 17
	lessEqual    tokenType = 18

	// Literals
	identifier    tokenType = 19
	stringLiteral tokenType = 20
	number        tokenType = 21

	// Keywords
	andKeyword    tokenType = 22
	classKeyword  tokenType = 23
	elseKeyword   tokenType = 24
	falseKeyword  tokenType = 25
	funKeyword    tokenType = 26
	forKeyword    tokenType = 27
	ifKeyword     tokenType = 28
	nilKeyword    tokenType = 29
	orKeyword     tokenType = 30
	printKeyword  tokenType = 31
	returnKeyword tokenType = 32
	superKeyword  tokenType = 33
	thisKeyword   tokenType = 34
	trueKeyword   tokenType = 35
	varKeyword    tokenType = 36
	whileKeyword  tokenType = 37

	eof tokenType = 38
)

func (ttype tokenType) String() string {
	return "test"
}
