package codegen

import (
	"github.com/xnacly/regexcc/cc/ast"
	"strings"
)

type Generator interface {
	Generate(ast.Node, *strings.Builder) error
}
