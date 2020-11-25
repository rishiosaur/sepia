package ast

import (
	"bytes"
	"sepia/token"
)

type ValueStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (ls *ValueStatement) statementNode()       {}
func (ls *ValueStatement) TokenLiteral() string { return ls.Token.Literal }
func (ls *ValueStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")

	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}

	out.WriteString(";")

	return out.String()
}

type UpdateStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (ls *UpdateStatement) statementNode()       {}
func (ls *UpdateStatement) TokenLiteral() string { return ls.Token.Literal }
func (ls *UpdateStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")

	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}

	out.WriteString(";")

	return out.String()
}
