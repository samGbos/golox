package golox

var reserved_words = map[string]TokenType{
	"and":    AndKeyword,
	"class":  ClassKeyword,
	"else":   ElseKeyword,
	"false":  FalseKeyword,
	"fun":    FunKeyword,
	"for":    ForKeyword,
	"if":     IfKeyword,
	"nil":    NilKeyword,
	"or":     OrKeyword,
	"print":  PrintKeyword,
	"return": ReturnKeyword,
	"super":  SuperKeyword,
	"this":   ThisKeyword,
	"true":   TrueKeyword,
	"var":    VarKeyword,
	"while":  WhileKeyword,
}

type TokenType int

const (
	// Single char tokens
	LeftParen  TokenType = 0
	RightParen TokenType = 1
	LeftBrace  TokenType = 2
	RightBrace TokenType = 3
	Comma      TokenType = 4
	Dot        TokenType = 5
	Minus      TokenType = 6
	Plus       TokenType = 7
	Semicolon  TokenType = 8
	Slash      TokenType = 9
	Star       TokenType = 10

	// One/two char tokens
	Bang         TokenType = 11
	BangEqual    TokenType = 12
	Equal        TokenType = 13
	EqualEqual   TokenType = 14
	Greater      TokenType = 15
	GreaterEqual TokenType = 16
	Less         TokenType = 17
	LessEqual    TokenType = 18

	// Literals
	Identifier    TokenType = 19
	StringLiteral TokenType = 20
	Number        TokenType = 21

	// Keywords
	AndKeyword    TokenType = 22
	ClassKeyword  TokenType = 23
	ElseKeyword   TokenType = 24
	FalseKeyword  TokenType = 25
	FunKeyword    TokenType = 26
	ForKeyword    TokenType = 27
	IfKeyword     TokenType = 28
	NilKeyword    TokenType = 29
	OrKeyword     TokenType = 30
	PrintKeyword  TokenType = 31
	ReturnKeyword TokenType = 32
	SuperKeyword  TokenType = 33
	ThisKeyword   TokenType = 34
	TrueKeyword   TokenType = 35
	VarKeyword    TokenType = 36
	WhileKeyword  TokenType = 37

	Eof TokenType = 38
)

func (ttype *TokenType) String() string {
	switch *ttype {
	case LeftParen:
		return "LeftParen"
	case RightParen:
		return "RightParen"
	case LeftBrace:
		return "LeftBrace"
	case RightBrace:
		return "RightBrace"
	case Comma:
		return "Comma"
	case Dot:
		return "Dot"
	case Minus:
		return "Minus"
	case Plus:
		return "Plus"
	case Semicolon:
		return "Semicolon"
	case Slash:
		return "Slash"
	case Star:
		return "Star"

	// One/two char tokens
	case Bang:
		return "Bang"
	case BangEqual:
		return "BangEqual"
	case Equal:
		return "Equal"
	case EqualEqual:
		return "EqualEqual"
	case Greater:
		return "Greater"
	case GreaterEqual:
		return "GreaterEqual"
	case Less:
		return "Less"
	case LessEqual:
		return "LessEqual"

	// Literals
	case Identifier:
		return "Identifier"
	case StringLiteral:
		return "StringLiteral"
	case Number:
		return "Number"

	// Keywords
	case AndKeyword:
		return "AndKeyword"
	case ClassKeyword:
		return "ClassKeyword"
	case ElseKeyword:
		return "ElseKeyword"
	case FalseKeyword:
		return "FalseKeyword"
	case FunKeyword:
		return "FunKeyword"
	case ForKeyword:
		return "ForKeyword"
	case IfKeyword:
		return "IfKeyword"
	case NilKeyword:
		return "NilKeyword"
	case OrKeyword:
		return "OrKeyword"
	case PrintKeyword:
		return "PrintKeyword"
	case ReturnKeyword:
		return "ReturnKeyword"
	case SuperKeyword:
		return "SuperKeyword"
	case ThisKeyword:
		return "ThisKeyword"
	case TrueKeyword:
		return "TrueKeyword"
	case VarKeyword:
		return "VarKeyword"
	case WhileKeyword:
		return "WhileKeyword"

	case Eof:
		return "Eof"
	default:
		return "Unknown"

	}
}
