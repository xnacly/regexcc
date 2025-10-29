package nfa

import "github.com/xnacly/regexcc/cc/ast"

type State int

type Edge struct {
	From  State
	To    State
	Label *[2]byte
}

type NFA struct {
	States int
	Edges  []Edge
	Start  State
	Accept map[State]bool
}

func (n *NFA) StatesInOrder() []State {
	states := make([]State, n.States)
	for i := 0; i < n.States; i++ {
		states[i] = State(i)
	}
	return states
}

func (n *NFA) EdgesFor(state State) []Edge {
	matches := make([]Edge, 0, len(n.Edges))
	for _, e := range n.Edges {
		if e.From == state {
			matches = append(matches, e)
		}
	}
	return matches
}

func From(_ ast.Node) NFA {
	// TODO: replace this to lower the AST to a NFA

	// ab*c
	n := NFA{
		Accept: make(map[State]bool),
	}

	s0 := State(0)
	s1 := State(1)
	s2 := State(2)

	n.Start = s0

	n.Edges = []Edge{
		{From: s0, To: s1, Label: &[2]byte{'a', 'a'}},
		{From: s1, To: s1, Label: &[2]byte{'b', 'b'}},
		{From: s1, To: s2, Label: nil},
	}

	n.Accept[s2] = true

	s3 := State(3)

	n.Edges = append(n.Edges, Edge{From: s2, To: s3, Label: &[2]byte{'c', 'c'}})
	n.Accept[s3] = true

	n.States = 4
	return n
}
