package ast

import (
	"github.com/xnacly/regexcc/cc/lexer"
)

type Type uint8

const (
	N_Empty    Type = iota
	N_Symbol        // a,b,c,ß,#,0
	N_Union         // union of Type: a|b, c|ß
	N_Concat        // concact of Type: abc, cßa 0123
	N_Star          // *
	N_Plus          // +
	N_Question      // ?
)

type Node interface {
	Type() Type
	Token() lexer.Token
	Children() []Node
}

func From(t lexer.Token, tt Type, children ...Node) Node {
	switch tt {
	case N_Empty:
		return nil
	case N_Symbol:
		return Symbol{
			t,
			tt,
		}
	case N_Union:
		return Union{
			t,
			tt,
			children,
		}
	case N_Concat:
		return Concat{
			t,
			tt,
			children,
		}
	case N_Star:
		return Repeat{
			t,
			tt,
			children,
		}
	default:
		panic("Impossible node state")
	}
}
