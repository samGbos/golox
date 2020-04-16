package golox

type token struct {
	ttype   tokenType
	lexeme  string
	literal interface{}
	line    int
}
