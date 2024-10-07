package expressions

import (
	"ast-operation-parser/ast/nodes"
	"ast-operation-parser/lexer/operators"
	"ast-operation-parser/lexer/token"
	"errors"
)

func InfixToRPN(tokens []token.Token) (nodes.ASTNode, error) {
	// Stack for operators
	operatorStack := []token.Token{}
	// Stack for ASTNodes
	output := []nodes.ASTNode{}
	expectValue := true
	parentheseCount := 0

	for i := 0; i < len(tokens); i++ {
		t := tokens[i]
		if i > 0 && tokens[i-1].Type == token.TOKEN_RPAREN && t.Type == token.TOKEN_MINUS {
			expectValue = false
		}
		// If the token is a value (number, string, bool, or variable), create the corresponding ASTNode
		if expectValue && IsValue(&tokens, &i) {
			astNode, err := ParseValue(t, tokens, &i) // Now ParseValue also handles function calls
			if err != nil {
				return nil, err
			}
			output = append(output, astNode)
			expectValue = false
			continue
		} else if t.Type == token.TOKEN_LPAREN {
			parentheseCount++
			operatorStack = append(operatorStack, t)
			continue
		} else if t.Type == token.TOKEN_RPAREN {
			parentheseCount--
			// Pop from operator stack until we find the left parenthesis
			for len(operatorStack) > 0 {
				top := operatorStack[len(operatorStack)-1]
				operatorStack = operatorStack[:len(operatorStack)-1]
				if top.Type == token.TOKEN_LPAREN {
					break
				}
				err := popOperatorToAST(&output, top)
				if err != nil {
					return nil, err
				}
			}
			continue
		} else if t.Type == token.TOKEN_COMMA {
			// Comma is used to separate function arguments; it should be ignored here in RPN
			continue
		} else if _, ok := token.Operators[t.Type]; ok {
			// Handle operators
			for len(operatorStack) > 0 {
				top := operatorStack[len(operatorStack)-1]
				if top.Type == token.TOKEN_LPAREN {
					break
				}

				o1 := operators.OperatorPrecedence[top.Type]
				o2 := operators.OperatorPrecedence[t.Type]

				if o1 > o2 {
					operatorStack = operatorStack[:len(operatorStack)-1]
					err := popOperatorToAST(&output, top)
					if err != nil {
						return nil, err
					}
				} else {
					break
				}
			}
			operatorStack = append(operatorStack, t)
			expectValue = true
		} else if t.Type == token.TOKEN_IDENT && i+1 < len(tokens) && tokens[i+1].Type == token.TOKEN_LPAREN {
			// This is the start of a function call
			i++ // Move past the identifier and parenthesis
			i++ // Skip past '('

			// Parse the function arguments using the helper function
			args, err := parseFunctionArguments(tokens, &i)
			if err != nil {
				return nil, err
			}

			// Create a FunctionCall AST node
			funcCall := nodes.FunctionCall{
				FunctionName: t.Value,
				Arguments:    args,
			}

			output = append(output, funcCall)
			expectValue = false
			continue
		}
	}

	if parentheseCount != 0 {
		return nil, errors.New("incorrectly placed parenthesis")
	}

	// Pop all the operators from the stack to the output
	for len(operatorStack) > 0 {
		err := popOperatorToAST(&output, operatorStack[len(operatorStack)-1])
		if err != nil {
			return nil, err
		}
		operatorStack = operatorStack[:len(operatorStack)-1]
	}

	// Final check: Ensure that the output has a valid single root AST node
	if len(output) != 1 {
		return nil, errors.New("invalid RPN expression: multiple root nodes")
	}

	return output[0], nil // Return the single root ASTNode
}

// Helper function to create a BinaryOperation AST node from the top operator
func popOperatorToAST(output *[]nodes.ASTNode, operator token.Token) error {
	if len(*output) < 2 {
		return errors.New("insufficient operands for binary operation")
	}

	// Pop the two top nodes from the output stack
	right := (*output)[len(*output)-1]
	left := (*output)[len(*output)-2]
	*output = (*output)[:len(*output)-2]

	// Accepted types and returned types for the binary operation
	acceptedTypes := operators.OperatorAcceptedTypes[operator.Type]
	returnedTypes := operators.OperatorReturnedTypes[operator.Type]

	// Create the BinaryOperation AST node
	binaryOp, err := nodes.NewBinaryOperation(left, right, operator.Value, acceptedTypes, returnedTypes)
	if err != nil {
		return err
	}

	// Push the binary operation result back onto the output stack
	*output = append(*output, binaryOp)
	return nil
}
