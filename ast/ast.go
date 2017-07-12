package ast

import "github.com/alarsyo/monkey/token"

// Node of the Abstract Syntax Tree
type Node interface {
	TokenLiteral() string
}

// Statement node of the AST
type Statement interface {
	Node
	statementNode()
}

// Expression node of the AST
type Expression interface {
	Node
	expressionNode()
}

// Program is the root node of the AST
type Program struct {
	Statements []Statement
}

// LetStatement node of the AST
type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {}

// TokenLiteral returns the token of the let statement (LET)
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

// Identifier node of the AST
type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode() {}

//TokenLiteral returns the token of the identifier (IDENT)
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
