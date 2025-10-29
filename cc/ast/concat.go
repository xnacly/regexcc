package ast

import "github.com/xnacly/regexcc/cc/lexer"

type Concat struct {
	T  lexer.Token
	Tt Type
	C  []Node
}

func (n Concat) Token() lexer.Token {
	return n.T
}

func (n Concat) Type() Type {
	return n.Tt
}

func (n Concat) Children() []Node {
	return n.C
}
