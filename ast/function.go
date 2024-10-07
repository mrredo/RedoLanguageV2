package ast

import (
	"ast-operation-parser/ast/nodes"
	"ast-operation-parser/lexer/token"
)

func IsFunction(t, t2 *token.Token) bool {
	return t.Type == token.TOKEN_IDENT && t2.Type == token.TOKEN_LPAREN
}
func ParseFunction(t *[]token.Token) (nodes.FunctionCall, error) {
	return nodes.FunctionCall{}, nil
}
