package main

import (
	"flag"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/xnacly/regexcc/cc/ast"
	"github.com/xnacly/regexcc/cc/codegen"
	"github.com/xnacly/regexcc/cc/config"
	"github.com/xnacly/regexcc/cc/lexer"
	"github.com/xnacly/regexcc/cc/nfa"
)

func main() {
	conf := config.Config{}
	flag.StringVar(&conf.Name, "name", fmt.Sprint("regex", time.Now().Unix()), "Define a name for both source files and function to generate")
	flag.StringVar(&conf.Generator, "gen", "c", "Which codegen backend to use")
	debug := flag.Bool("debug", false, "view debug output")
	flag.Parse()
	if *debug {
		slog.SetDefault(slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		})))
	} else {
		slog.SetDefault(slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		})))
	}

	conf.Input = flag.Arg(0)
	if len(conf.Input) == 0 {
		slog.Error("no regex to compile")
		os.Exit(1)
	}

	var g codegen.Generator
	switch conf.Generator {
	case "c":
		g = &codegen.C{}
	case "rust":
		panic("Unimplemented")
	case "go":
		panic("Unimplemented")
	default:
		slog.Error("invalid codegen backend selected", "backend", conf.Generator, "expected", "c,rust,go")
		os.Exit(1)
	}

	slog.Info("starting compilation", "regex", conf.Input, "codegen", conf.Generator)

	// TODO: lex
	// TODO: parse

	// equiv to /ab/
	tree := ast.From(
		lexer.From(lexer.Epsilon, 0),
		ast.N_Concat,
		ast.From(lexer.From(lexer.Symbol, 'a'), ast.N_Symbol),
		ast.From(lexer.From(lexer.Symbol, 'b'), ast.N_Symbol),
	)

	// TODO: lower from ast to nfa via thompson
	nfa := nfa.From(tree)
	g.Generate(nfa, conf)
}
