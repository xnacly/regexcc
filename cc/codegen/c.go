package codegen

import (
	"errors"
	"strings"

	"github.com/xnacly/regexcc/cc/ast"
)

// C codegen
type C struct{}

func (c *C) Preamble() error {
	// TODO: create <name>.(c|h)
	return errors.New("Unimplemented")
}

func (c *C) Generate(n ast.Node, buf *strings.Builder) error {
	switch n.Type() {
	// something like 5, should just match '5', thus
	case ast.LITERAL:
		buf.WriteString("in[i] == '")
		buf.WriteRune(n.Token().Rune)
		buf.WriteRune('\'')
	}

	return nil
}
