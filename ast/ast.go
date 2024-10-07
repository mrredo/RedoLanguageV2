package ast

import (
	"ast-operation-parser/ast/nodes"
	"ast-operation-parser/lexer/token"
)

// // Example function to print the AST (for debugging purposes)
// func PrintAST(node *ASTNode, depth int) {
// 	if node == nil {
// 		return
// 	}
// 	for i := 0; i < depth; i++ {
// 		fmt.Print("  ")
// 	}
// 	fmt.Printf("Node: { Type: %s, Value: %s }\n", node.Type, node.Value)
// 	PrintAST(node.Left, depth+1)
// 	PrintAST(node.Right, depth+1)
// }

// Precedence map for Go operators

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
func GroupTokens(tokens []token.Token) ([][]token.Token, error) {
	groups := [][]token.Token{}

	return groups, nil
}

// func main() {
// 	a := []Token{
// 		Token{Type: Number, Value: "1"},
// 		Token{Type: Operator, Value: "*"},
// 		Token{Type: Number, Value: "2"},
// 		Token{Type: Operator, Value: "+"},
// 		Token{Type: Number, Value: "5"},
// 	}
// 	b := []Token{
// 		Token{Type: Number, Value: "1"},
// 		Token{Type: Operator, Value: "*"},
// 		Token{Type: RParen, Value: "("},
// 		Token{Type: Number, Value: "2"},
// 		Token{Type: Operator, Value: "+"},
// 		Token{Type: Number, Value: "5"},
// 		Token{Type: LParen, Value: ")"},
// 	}
// 	fmt.Println(infixToRPN(a), infixToRPN(b))
// }
