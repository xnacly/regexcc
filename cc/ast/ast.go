package ast

import (
	"github.com/xnacly/regexcc/cc/lexer"
)

type Type uint8

const (
	// a literal, may be 0,5,k or ß
	LITERAL Type = iota
)

type Node interface {
	Type() Type
	Token() lexer.Token
	Children() []Node
}
