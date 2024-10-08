package token

import "text/scanner"

const (
	// Token types
	TOKEN_NUMBER = "NUMBER"
	TOKEN_STRING = "STRING"
	TOKEN_BOOL   = "BOOL"

	//true/false tokens not used in code, but are for translations and lexer
	TOKEN_TRUE  = "true"
	TOKEN_FALSE = "false"

	TOKEN_PLUS          = "+"
	TOKEN_MINUS         = "-"
	TOKEN_MUL           = "*"
	TOKEN_DIV           = "/"
	TOKEN_MODULO        = "%"
	TOKEN_LPAREN        = "("
	TOKEN_RPAREN        = ")"
	TOKEN_NEWLINE       = "NEWLINE"
	TOKEN_INVALID       = "INVALID"
	TOKEN_IDENT         = "IDENT"
	TOKEN_VAR           = "var"
	TOKEN_EQ            = "=="
	TOKEN_NEQ           = "!="
	TOKEN_LT            = "<"
	TOKEN_LTE           = "<="
	TOKEN_GT            = ">"
	TOKEN_GTE           = ">="
	TOKEN_AND           = "&&"
	TOKEN_OR            = "||"
	TOKEN_BIT_AND       = "&"
	TOKEN_BIT_OR        = "|"
	TOKEN_BIT_XOR       = "^"
	TOKEN_BIT_NOT       = "~"
	TOKEN_LEFT_SHIFT    = "<<"
	TOKEN_RIGHT_SHIFT   = ">>"
	TOKEN_INCREMENT     = "++"
	TOKEN_DECREMENT     = "--"
	TOKEN_ASSIGN        = "="
	TOKEN_ASSIGN_ADD    = "+="
	TOKEN_ASSIGN_SUB    = "-="
	TOKEN_ASSIGN_MUL    = "*="
	TOKEN_ASSIGN_DIV    = "/="
	TOKEN_ASSIGN_MOD    = "%="
	TOKEN_ASSIGN_AND    = "&="
	TOKEN_ASSIGN_OR     = "|="
	TOKEN_ASSIGN_XOR    = "^="
	TOKEN_ASSIGN_LSHIFT = "<<="
	TOKEN_ASSIGN_RSHIFT = ">>="

	TOKEN_SEMICOLON = ";"
	TOKEN_COMMA     = ","
	TOKEN_DOT       = "."
	TOKEN_COLON     = ":"
	TOKEN_EOF       = "EOF" // end of file
)

// TokenMap[Token]Token
var TokenMap = map[string]string{
	TOKEN_NUMBER:        TOKEN_NUMBER,
	TOKEN_PLUS:          TOKEN_PLUS,
	TOKEN_MINUS:         TOKEN_MINUS,
	TOKEN_MUL:           TOKEN_MUL,
	TOKEN_DIV:           TOKEN_DIV,
	TOKEN_MODULO:        TOKEN_MODULO,
	TOKEN_LPAREN:        TOKEN_LPAREN,
	TOKEN_RPAREN:        TOKEN_RPAREN,
	TOKEN_NEWLINE:       TOKEN_NEWLINE,
	TOKEN_INVALID:       TOKEN_INVALID,
	TOKEN_IDENT:         TOKEN_IDENT,
	TOKEN_VAR:           TOKEN_VAR,
	TOKEN_EQ:            TOKEN_EQ,
	TOKEN_NEQ:           TOKEN_NEQ,
	TOKEN_LT:            TOKEN_LT,
	TOKEN_LTE:           TOKEN_LTE,
	TOKEN_GT:            TOKEN_GT,
	TOKEN_GTE:           TOKEN_GTE,
	TOKEN_AND:           TOKEN_AND,
	TOKEN_OR:            TOKEN_OR,
	TOKEN_BIT_AND:       TOKEN_BIT_AND,
	TOKEN_BIT_OR:        TOKEN_BIT_OR,
	TOKEN_BIT_XOR:       TOKEN_BIT_XOR,
	TOKEN_BIT_NOT:       TOKEN_BIT_NOT,
	TOKEN_LEFT_SHIFT:    TOKEN_LEFT_SHIFT,
	TOKEN_RIGHT_SHIFT:   TOKEN_RIGHT_SHIFT,
	TOKEN_INCREMENT:     TOKEN_INCREMENT,
	TOKEN_DECREMENT:     TOKEN_DECREMENT,
	TOKEN_ASSIGN:        TOKEN_ASSIGN,
	TOKEN_ASSIGN_ADD:    TOKEN_ASSIGN_ADD,
	TOKEN_ASSIGN_SUB:    TOKEN_ASSIGN_SUB,
	TOKEN_ASSIGN_MUL:    TOKEN_ASSIGN_MUL,
	TOKEN_ASSIGN_DIV:    TOKEN_ASSIGN_DIV,
	TOKEN_ASSIGN_MOD:    TOKEN_ASSIGN_MOD,
	TOKEN_ASSIGN_AND:    TOKEN_ASSIGN_AND,
	TOKEN_ASSIGN_OR:     TOKEN_ASSIGN_OR,
	TOKEN_ASSIGN_XOR:    TOKEN_ASSIGN_XOR,
	TOKEN_ASSIGN_LSHIFT: TOKEN_ASSIGN_LSHIFT,
	TOKEN_ASSIGN_RSHIFT: TOKEN_ASSIGN_RSHIFT,
	TOKEN_SEMICOLON:     TOKEN_SEMICOLON,
	TOKEN_COMMA:         TOKEN_COMMA,
	TOKEN_DOT:           TOKEN_DOT,
	TOKEN_COLON:         TOKEN_COLON,
	TOKEN_EOF:           TOKEN_EOF,
	TOKEN_STRING:        TOKEN_STRING,
	TOKEN_BOOL:          TOKEN_BOOL,
}

// Operators[Token]bool
var Operators = map[string]bool{
	TOKEN_PLUS:          true,
	TOKEN_MINUS:         true,
	TOKEN_MUL:           true,
	TOKEN_DIV:           true,
	TOKEN_MODULO:        true,
	TOKEN_EQ:            true,
	TOKEN_NEQ:           true,
	TOKEN_LT:            true,
	TOKEN_LTE:           true,
	TOKEN_GT:            true,
	TOKEN_GTE:           true,
	TOKEN_AND:           true,
	TOKEN_OR:            true,
	TOKEN_BIT_AND:       true,
	TOKEN_BIT_OR:        true,
	TOKEN_BIT_XOR:       true,
	TOKEN_BIT_NOT:       true,
	TOKEN_LEFT_SHIFT:    true,
	TOKEN_RIGHT_SHIFT:   true,
	TOKEN_INCREMENT:     true,
	TOKEN_DECREMENT:     true,
	TOKEN_ASSIGN:        true,
	TOKEN_ASSIGN_ADD:    true,
	TOKEN_ASSIGN_SUB:    true,
	TOKEN_ASSIGN_MUL:    true,
	TOKEN_ASSIGN_DIV:    true,
	TOKEN_ASSIGN_MOD:    true,
	TOKEN_ASSIGN_AND:    true,
	TOKEN_ASSIGN_OR:     true,
	TOKEN_ASSIGN_XOR:    true,
	TOKEN_ASSIGN_LSHIFT: true,
	TOKEN_ASSIGN_RSHIFT: true,
	// TOKEN_SEMICOLON:    false, // Not an operator, just included for completeness
	// TOKEN_COMMA:        false,
	// TOKEN_DOT:          false,
	// TOKEN_COLON:        false,
}

// Token structure to hold the type and value of a token
type Token struct {
	Type  string
	Value string
	Pos   scanner.Position // Position in input for error reporting
}
