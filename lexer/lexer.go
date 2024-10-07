package lexer

import (
	"ast-operation-parser/lexer/token"
	"fmt"
	"strconv"
	"strings"
	"text/scanner"
	"unicode"
)

// Lexer struct
type Lexer struct {
	scanner scanner.Scanner // Use Go's text/scanner package
	token   token.Token     // Current token being parsed
}

// Create a new lexer and initialize the scanner
func NewLexer(input string) *Lexer {
	var s scanner.Scanner
	s.Init(strings.NewReader(input))
	s.Mode = scanner.ScanIdents | scanner.ScanInts | scanner.ScanFloats | scanner.ScanChars
	s.Whitespace = 0 // To track whitespace like newlines

	lexer := &Lexer{scanner: s}
	return lexer
}

// Get the next token from the scanner
func (l *Lexer) NextToken() token.Token {
	r := l.scanner.Scan()
	pos := l.scanner.Pos()

	switch r {
	case '-': // Changed from '–' to '-' for correct minus sign handling
		_, err := strconv.Atoi(string(l.scanner.Peek()))
		if err == nil {
			l.scanner.Scan() // Consume the number
			return token.Token{Type: token.TOKEN_NUMBER, Value: "-" + l.scanner.TokenText(), Pos: pos}
		}
		// If it's not a number, just treat as subtraction operator
		return token.Token{Type: token.TOKEN_MINUS, Value: "-", Pos: pos}

	case scanner.Int, scanner.Float:
		return token.Token{Type: token.TOKEN_NUMBER, Value: l.scanner.TokenText(), Pos: pos}

	case scanner.String:
		return token.Token{Type: token.TOKEN_STRING, Value: l.scanner.TokenText(), Pos: pos}

	case '\n':
		return token.Token{Type: token.TOKEN_NEWLINE, Value: token.TOKEN_NEWLINE, Pos: pos}

	case scanner.EOF:
		return token.Token{Type: token.TOKEN_EOF, Value: token.TOKEN_EOF, Pos: pos} // End of input

	default:
		if scanner.Ident == r {
			switch l.scanner.TokenText() {
			case "var":
				return token.Token{Type: token.TOKEN_VAR, Value: l.scanner.TokenText(), Pos: pos}
			case "true", "false":
				return token.Token{Type: token.TOKEN_BOOL, Value: l.scanner.TokenText(), Pos: pos}
			}
			return token.Token{Type: token.TOKEN_IDENT, Value: l.scanner.TokenText(), Pos: pos}
		}

		if unicode.IsSpace(r) {
			return l.NextToken() // Skip spaces and continue
		} else if isOperator(r) {
			// Handle multi-character operators
			return l.parseOperator(r, pos)
		}
	}

	// If r is not a recognized token type
	if v, ok := token.TokenMap[string(r)]; ok {
		return token.Token{Type: v, Value: l.scanner.TokenText(), Pos: pos}
	}
	return token.Token{Type: token.TOKEN_INVALID, Value: l.scanner.TokenText(), Pos: pos}
}

// Check if a rune is the start of a known operator
func isOperator(r rune) bool {
	return strings.ContainsRune("+-*/&|^~<>!=", r)
}

// Parse multi-character operators, including assignment operators
func (l *Lexer) parseOperator(firstRune rune, pos scanner.Position) token.Token {
	operatorText := string(firstRune)

	// Check for additional characters for multi-character operators
	switch firstRune {
	case '!':
		if l.scanner.Peek() == '=' {
			l.scanner.Next() // Consume the '='
			operatorText = "!="
		}
	case '&':
		next := l.scanner.Peek()
		if next == '&' {
			l.scanner.Next() // Consume the second '&'
			operatorText = "&&"
		} else if next == '=' {
			l.scanner.Next() // Consume the '=' for '&='
			operatorText = "&="
		}
	case '|':
		next := l.scanner.Peek()
		if next == '|' {
			l.scanner.Next() // Consume the second '|'
			operatorText = "||"
		} else if next == '=' {
			l.scanner.Next() // Consume the '=' for '|='
			operatorText = "|="
		}
	case '<':
		next := l.scanner.Peek()
		if next == '=' {
			l.scanner.Next() // Consume the '='
			operatorText = "<="
		} else if next == '<' {
			l.scanner.Next() // Consume the second '<'
			if l.scanner.Peek() == '=' {
				l.scanner.Next() // Consume the '=' for '<<='
				operatorText = "<<="
			} else {
				operatorText = "<<"
			}
		}
	case '>':
		next := l.scanner.Peek()
		if next == '=' {
			l.scanner.Next() // Consume the '='
			operatorText = ">="
		} else if next == '>' {
			l.scanner.Next() // Consume the second '>'
			if l.scanner.Peek() == '=' {
				l.scanner.Next() // Consume the '=' for '>>='
				operatorText = ">>="
			} else {
				operatorText = ">>"
			}
		}
	case '=':
		if l.scanner.Peek() == '=' {
			l.scanner.Next() // Consume the second '='
			operatorText = "=="
		}
	case '+':
		next := l.scanner.Peek()
		if next == '=' {
			l.scanner.Next() // Consume the '='
			operatorText = "+="
		} else if next == '+' {
			l.scanner.Next() // Consume the second '+'
			operatorText = "++"
		}
	case '-':
		next := l.scanner.Peek()
		if next == '=' {
			l.scanner.Next() // Consume the '='
			operatorText = "-="
		} else if next == '-' {
			l.scanner.Next() // Consume the second '-'
			operatorText = "--"
		}
	case '*':
		if l.scanner.Peek() == '=' {
			l.scanner.Next() // Consume the '='
			operatorText = "*="
		}
	case '/':
		if l.scanner.Peek() == '=' {
			l.scanner.Next() // Consume the '='
			operatorText = "/="
		}
	case '%':
		if l.scanner.Peek() == '=' {
			l.scanner.Next() // Consume the '=' for '%='
			operatorText = "%="
		}
	case '^':
		if l.scanner.Peek() == '=' {
			l.scanner.Next() // Consume the '=' for '^='
			operatorText = "^="
		}
	case '~':
		// '~' bitwise NOT operator, no multi-character version
	}

	// Check if the operator exists in the Operators map (including assignment operators)
	if _, exists := token.Operators[operatorText]; exists {
		return token.Token{Type: operatorText, Value: operatorText, Pos: pos}
	}

	// If no valid operator is found, return an INVALID token
	return token.Token{Type: token.TOKEN_INVALID, Value: operatorText, Pos: pos}
}

// Tokenize

// Tokenize function
func Tokenize(input string) []token.Token {
	lexer := NewLexer(input)
	var tokens []token.Token

	for {
		t := lexer.NextToken()
		tokens = append(tokens, t)

		if t.Type == token.TOKEN_EOF {
			break // Stop when nextToken returns nil
		}
	}

	return tokens
}

// ValidateExpression checks if the token sequence forms a valid expression
func ValidateExpression(lexer *Lexer) error {
	expectValue := true // Start by expecting a value (number or variable)

	for {
		t := lexer.NextToken()
		if t.Type == token.TOKEN_EOF {
			break // End of input
		}

		// Check for newlines
		if t.Type == token.TOKEN_NEWLINE {
			continue
		}
		_, ok := token.TokenMap[t.Type]
		// Validate token sequence based on expected type
		if expectValue && t.Type == token.TOKEN_NUMBER {
			expectValue = false // After a value, expect an operator
		} else if !expectValue && ok {
			expectValue = true // After an operator, expect another value
		} else {
			return fmt.Errorf("syntax error at %v: unexpected '%s'", t.Pos, t.Value)
		}
	}

	if expectValue {
		return fmt.Errorf("expression ends with an operator")
	}

	return nil
}
