package parser

import (
	"github.com/alarsyo/monkey/ast"
	"github.com/alarsyo/monkey/lexer"
	"github.com/alarsyo/monkey/token"
)

// Parser is responsible for organizing the Tokens logically into an AST
type Parser struct {
	l *lexer.Lexer

	curToken  token.Token
	peekToken token.Token
}

// New Parser
func New(l *lexer.Lexer) (p *Parser) {
	p = &Parser{l: l}

	// Read two tokens, so curToken and peekToken are both set
	p.nextToken()
	p.nextToken()
	return
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

// ParseProgram builds the Program's entire AST
func (p *Parser) ParseProgram() *ast.Program {
	return nil
}
