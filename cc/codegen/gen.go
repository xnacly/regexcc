package codegen

import (
	"github.com/xnacly/regexcc/cc/config"
	"github.com/xnacly/regexcc/cc/nfa"
)

type Generator interface {
	Generate(nfa.NFA, config.Config) error
}
