package golox

type Token struct {
	Ttype   TokenType
	Lexeme  string
	Literal interface{}
	Line    int
	// Start marks the start position of this token on this line
	Start int
	// End marks the end position of this token on this line
	End int
}
