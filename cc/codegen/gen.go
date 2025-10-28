package codegen

import (
	"github.com/xnacly/regexcc/cc/ast"
	"strings"
)

// TODO: make file agnostic and pass CcConfig in, containing filename and stuff
type Generator interface {
	Generate(ast.Node, *strings.Builder) error
	Preamble() error
}
