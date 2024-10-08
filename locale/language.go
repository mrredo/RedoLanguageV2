package locale

import "ast-operation-parser/lexer/token"

var Language = "en"

var Languages = []string{"en", "lv"}

// create a map[string]string with 10 elements
var EnglishToLatvian = map[string]string{
	token.TOKEN_TRUE:  "patiesi",
	token.TOKEN_FALSE: "nepatiesi",
	"if":              "ja",
	"else":            "citādi",
	"else if":         "citādi ja",
	"while":           "kamēr",
	token.TOKEN_VAR:   "mainīgais",
	//iztulkot pareizi latviešu valodā
	"for": "",
}
var EnglishToEnglish = map[string]string{
	token.TOKEN_TRUE:  token.TOKEN_TRUE,
	token.TOKEN_FALSE: token.TOKEN_FALSE,
	"if":              "if",
	"else":            "else",
	"else if":         "else if",
	"while":           "while",
	token.TOKEN_VAR:   token.TOKEN_VAR,
	//iztulkot pareizi latviešu valodā
	"for": "",
}
