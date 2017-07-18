package ast

import (
	"bytes"

	"github.com/alarsyo/monkey/token"
)

// Node of the Abstract Syntax Tree
type Node interface {
	TokenLiteral() string
	String() string
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

func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
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

func (ls *LetStatement) String() string {
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

// Identifier node of the AST
type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode() {}

//TokenLiteral returns the token of the identifier (IDENT)
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

func (i *Identifier) String() string { return i.Value }

// ReturnStatement node of the AST
type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode() {}

// TokenLiteral returns the token of the return statement (RETURN)
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }

func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")

	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}

	out.WriteString(";")

	return out.String()
}

// ExpressionStatement node of the AST
type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}

func (es *ExpressionStatement) statementNode() {}

// TokenLiteral returns the first token of the expression
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }

func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}

type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (il *IntegerLiteral) expressionNode() {}

// TokenLiteral returns the token literal of the integer
func (il *IntegerLiteral) TokenLiteral() string { return il.Token.Literal }

func (il *IntegerLiteral) String() string { return il.Token.Literal }
