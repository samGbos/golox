package golox

type Token struct {
	Ttype   TokenType
	Lexeme  string
	Literal interface{}
	Line    int
}
