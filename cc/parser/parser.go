// parser implements the following grammar:
// Regex     → Alt
// Union       → Concat ('|' Concat)*
// Concat    → Repeat+
// Repeat    → Atom ('*' | '+' | '?')?
// Symbol      → Char | '(' Regex ')'
package parser

import (
	"go/ast"

	"github.com/xnacly/regexcc/cc/lexer"
)

func From(tokens []lexer.Token) ast.Node {
	return nil
}
