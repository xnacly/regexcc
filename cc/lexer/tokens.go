package lexer

type Type uint8

const (
	Epsilon Type = iota
	Symbol
	Pipe
	Star
	BraceLeft
	BraceRight
)

type Token struct {
	T    Type
	Rune rune
}

func From(t Type, r rune) Token {
	return Token{t, r}
}
