package ast

import (
	"testing"

	"github.com/arjunmalhotra1/monkey-interpreter/token"
)

// let myVar = anotherVar
func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{
					Type:    token.LET,
					Literal: "let",
				},
				Name: &Identifier{
					Token: token.Token{
						Type:    token.IDENT,
						Literal: "myVar",
					},
					Value: "myVar",
				},
				Value: &Identifier{ // Value is an expression and Identifier implements the expression interface.
					Token: token.Token{
						Type:    token.IDENT,
						Literal: "anotherVar",
					},
					Value: "anotherVar",
				},
			},
		},
	}
	if program.String() != "let myVar = anotherVar;" {
		t.Errorf("program.String() wrong. got = %q", program.String())
	}
}
