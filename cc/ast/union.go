package ast

import "github.com/xnacly/regexcc/cc/lexer"

type Union struct {
	T  lexer.Token
	Tt Type
	C  []Node
}

func (n Union) Token() lexer.Token {
	return n.T
}

func (n Union) Type() Type {
	return n.Tt
}

func (n Union) Children() []Node {
	return n.C
}
