package ast

import "github.com/xnacly/regexcc/cc/lexer"

type Symbol struct {
	T  lexer.Token
	Tt Type
}

func (n Symbol) Token() lexer.Token {
	return n.T
}

func (n Symbol) Type() Type {
	return n.Tt
}

func (n Symbol) Children() []Node {
	return nil
}
