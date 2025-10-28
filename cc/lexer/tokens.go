package lexer

type Type uint8

const (
	LITERAL Type = iota
)

type Token struct {
	Pos  uint
	Rune rune
}
