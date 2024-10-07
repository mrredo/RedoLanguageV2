package expressions

import (
	"ast-operation-parser/ast/nodes"
	"ast-operation-parser/lexer/token"
	"errors"
)

func IsValue(tokens *[]token.Token, idx *int) bool {
	t := (*tokens)[*idx]
	if t.Type == token.TOKEN_NUMBER || t.Type == token.TOKEN_STRING || t.Type == token.TOKEN_BOOL {
		return true
	}

	if t.Type == token.TOKEN_IDENT {
		// Check if it's a function call by looking ahead for the `(` token
		if *idx+1 < len(*tokens) && (*tokens)[*idx+1].Type == token.TOKEN_LPAREN {
			return true
		}
		return true // Treat identifiers as variables
	}

	return false
}
func ParseValue(tokens *[]token.Token, i *int) (nodes.ASTNode, error) {
	t := (*tokens)[*i]
	if t.Type == token.TOKEN_NUMBER || t.Type == token.TOKEN_STRING || t.Type == token.TOKEN_BOOL {
		return nodes.LiteralNode{
			DataType: nodes.TokenToDataType[t.Type],
			Value:    t.Value,
		}, nil
	} else if t.Type == token.TOKEN_IDENT {
		// Check if the next token is a left parenthesis indicating a function call
		if *i+1 < len(*tokens) && (*tokens)[*i+1].Type == token.TOKEN_LPAREN {
			*i += 2 // Move past the identifier and left parenthesis
			expectComma := false
			var args []nodes.ASTNode

			for {
				// Ensure we don't exceed the bounds of tokens
				if *i >= len(*tokens) {
					return nil, errors.New("invalid function call: missing closing parenthesis")
				}

				nextT := (*tokens)[*i]

				// Handle end of function call
				if nextT.Type == token.TOKEN_RPAREN {
					*i++ // Move past the closing parenthesis
					break
				}

				// Handle comma expectations
				if expectComma {
					if nextT.Type != token.TOKEN_COMMA {
						return nil, errors.New("invalid function call: expected comma")
					}
					expectComma = false // Reset the expectation for a value
					*i++                // Move past the comma
					continue
				}

				// Parse the argument value
				if IsValue(tokens, i) {
					argNode, err := ParseValue(tokens, i)
					if err != nil {
						return nil, err
					}
					args = append(args, argNode)
					expectComma = true // After a value, we expect a comma or closing parenthesis
				} else {
					return nil, errors.New("invalid function call: expected a value")
				}
			}

			// Create and return the function call AST node
			funcCallNode := nodes.FunctionCall{
				FunctionName: t.Value,
				Arguments:    args,
			}
			return funcCallNode, nil
		} else {
			// If not followed by a left parenthesis, it's an invalid function call
			return nil, errors.New("invalid expression: identifier not followed by function call")
		}
	}

	return nodes.VariableNode{
		Value: t.Value,
	}, nil

}
