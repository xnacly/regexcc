package ast

import "github.com/xnacly/regexcc/cc/lexer"

// Can be either +,* or ?
type Repeat struct {
	T  lexer.Token
	Tt Type
	C  []Node
}

func (n Repeat) Token() lexer.Token {
	return n.T
}

func (n Repeat) Type() Type {
	return n.Tt
}

func (n Repeat) Children() []Node {
	return n.C
}
