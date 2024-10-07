package expressions

import (
	"ast-operation-parser/ast/nodes"
	"ast-operation-parser/lexer/operators"
	"ast-operation-parser/lexer/token"
	"errors"
	"fmt"
	"slices"
)

func InfixToRPNAST(tokens []token.Token) (nodes.ASTNode, error) {
	// Stack for operators
	operatorStack := []token.Token{}
	// Stack for ASTNodes
	output := []nodes.ASTNode{}
	expectValue := true
	parentheseCount := 0

	// Error check for invalid length of tokens
	if len(tokens)%2 == 0 {
		return nil, errors.New("invalid expression: even number of tokens")
	}

	for i, t := range tokens {
		if i > 0 && tokens[i-1].Type == token.TOKEN_RPAREN && t.Type == token.TOKEN_MINUS {
			expectValue = false
		}

		// If the token is a value (number, string, bool, or variable), create the corresponding ASTNode
		if expectValue && IsValue(&tokens, &i) {
			astNode, err := ParseValue(&tokens, &i)
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
				// Handle the operator popping to create BinaryOperation nodes
				err := popOperatorToAST(&output, top)
				if err != nil {
					return nil, err
				}
			}
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

	// Get accepted and returned types for the operator
	acceptedTypes := operators.OperatorAcceptedTypes[operator.Type]
	returnedTypes := operators.OperatorReturnedTypes[operator.Type]

	// Check if operand types are compatible with the operator
	var rightR, leftR = nodes.GetDataType(right), nodes.GetDataType(left)

	// Check if rightR types are accepted by the operator
	accepted := false
	if len(rightR) != 0 {
		for _, rt := range rightR {
			if slices.Contains(acceptedTypes, rt) {
				accepted = true
				break
			}
		}
	} else {
		accepted = true
	}

	if !accepted {
		return errors.New("invalid operator type on the right side")
	}

	// Check if leftR types are accepted by the operator
	accepted = false
	fmt.Println(leftR, acceptedTypes)
	if len(leftR) != 0 {
		for _, lt := range leftR {
			if slices.Contains(acceptedTypes, lt) {
				accepted = true
				break
			}
		}
	} else {
		accepted = true
	}

	if !accepted {
		return errors.New("invalid operator type on the left side")
	}

	// Ensure the right and left types are compatible

	matches := false
	if len(rightR) != 0 && len(leftR) != 0 {
		for _, rt := range rightR {
			for _, lt := range leftR {
				if rt == lt {
					matches = true
				}
			}
		}
	} else {
		matches = true
	}
	if !matches {
		return errors.New("mismatched operand types")
	}

	// If all checks pass, create the BinaryOperation node
	binaryOp, err := nodes.NewBinaryOperation(left, right, operator.Value, acceptedTypes, returnedTypes)
	if err != nil {
		return err
	}

	// Push the new BinaryOperation node back onto the output stack
	*output = append(*output, binaryOp)
	return nil
}
