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
func ParseValue(t token.Token, tokens []token.Token, idx *int) (nodes.ASTNode, error) {
	if t.Type == token.TOKEN_NUMBER || t.Type == token.TOKEN_STRING || t.Type == token.TOKEN_BOOL {
		return nodes.LiteralNode{
			DataType: nodes.TokenToDataType[t.Type],
			Value:    t.Value,
		}, nil
	} else if t.Type == token.TOKEN_IDENT {
		// Check if the next token is an opening parenthesis `(`, indicating a function call
		if *idx+1 < len(tokens) && tokens[*idx+1].Type == token.TOKEN_LPAREN {
			*idx++ // Move past the identifier (function name)
			*idx++ // Move past the opening parenthesis `(`

			// Parse the function call arguments
			args, err := parseFunctionArguments(tokens, idx)
			if err != nil {
				return nil, err
			}

			return nodes.FunctionCall{
				FunctionName: t.Value,
				Arguments:    args,
			}, nil
		}

		// Otherwise, treat it as a variable
		return nodes.VariableNode{
			Value: t.Value,
		}, nil
	}
	return nil, errors.New("not a value")
}

// Function to parse arguments for function calls
func parseFunctionArguments(tokens []token.Token, idx *int) ([]nodes.ASTNode, error) {
	args := []nodes.ASTNode{}
	parenCount := 1 // To ensure we properly track parentheses and avoid errors
	argTokens := []token.Token{}

	for *idx < len(tokens) {
		currToken := tokens[*idx]

		if currToken.Type == token.TOKEN_LPAREN {
			parenCount++
		} else if currToken.Type == token.TOKEN_RPAREN {
			parenCount--
			if parenCount == 0 {
				// End of function arguments
				//*idx++ // Move past the closing parenthesis
				break
			}
		} else if currToken.Type == token.TOKEN_COMMA && parenCount == 1 {
			// If we encounter a comma at the top level of the argument list, treat it as the end of one argument
			arg, err := InfixToRPN(argTokens)
			if err != nil {
				return nil, err
			}
			args = append(args, arg)
			argTokens = []token.Token{} // Reset for the next argument
			*idx++
			continue
		}

		argTokens = append(argTokens, currToken)
		*idx++
	}

	// Handle the last argument after the loop
	if len(argTokens) > 0 {

		arg, err := InfixToRPN(argTokens)
		if err != nil {
			return nil, err
		}
		args = append(args, arg)
	}

	return args, nil
}
