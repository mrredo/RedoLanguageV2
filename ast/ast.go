package ast

import (
	"ast-operation-parser/ast/nodes"
	"ast-operation-parser/lexer/token"
)

// BuildAST constructs the AST from the given token
func BuildAST(tokens []token.Token) (*nodes.ASTStructure, error) {
	structure := &nodes.ASTStructure{
		Children: []nodes.ASTNode{},
	}
	for i := 0; i < len(tokens); i++ {
		t := tokens[i]
		if t.Type == token.TOKEN_EOF {
			break
		}
		if TokenEnd(t) {
			continue
		}
		switch t.Type {
		case token.TOKEN_VAR:

			g, err := MakeVariableDefinition(&tokens, &i)
			if err != nil {

				return structure, err
			}
			structure.Children = append(structure.Children, *g)

		case token.TOKEN_IDENT:
			if tokens[i+1].Type != token.TOKEN_LPAREN {
				varA, err := MakeVariableAssigning(&tokens, &i)
				if err != nil {

					return structure, err
				}
				structure.Children = append(structure.Children, *varA)

			}
		}
	}
	return structure, nil
}
func TokenEnd(t token.Token) bool {
	return t.Type == token.TOKEN_NEWLINE || t.Type == token.TOKEN_SEMICOLON || t.Type == token.TOKEN_EOF
}
