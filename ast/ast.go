// Package ast defines our Abstract Syntax Tree
package ast

import "monkey/token"

type Node interface {
	TokenLiteral() string
}

// Statement - Every valid Monkey program is a series of statements.
// These statements are contained in the Program.Statements, which is just
// a slice of AST nodes that implement the Statement interface.
type Statement interface {
	Node
	statementNode()
}

// Expression - a syntactic combination of variables, constants, operators and functions
// that the system evaluates to produce a single value
type Expression interface {
	Node
	expressionNode()
}

// Program node is going to be the root node of every AST
// our parser producces.
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

// LetStatement have an identifier and a value with an assignment between them
type LetStatement struct {
	Token token.Token // the token.LET token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

// ReturnStatement consist solely of the keyword 'return' and an expression
type ReturnStatement struct {
	Token       token.Token // the 'return' token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }

type Identifier struct {
	Token token.Token // the token.Ident token
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
