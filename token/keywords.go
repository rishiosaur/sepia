package token

var keywords = map[string]Type{
	"f":       FUNCTION,
	"value":   VALUE,
	"let":     VALUE,
	"true":    TRUE,
	"false":   FALSE,
	"if":      IF,
	"else":    ELSE,
	"return":  RETURN,
	"defined": ASSIGN,
	"be":      ASSIGN,
	"is":      EQ,
	"not":     NOT_EQ,
	"do":      LBRACE,
	"end":     RBRACE,
}

//LookupIdent finds an identifier token type from a string.
func LookupIdent(identifier string) Type {
	if tok, ok := keywords[identifier]; ok {
		return tok
	}

	return IDENT
}
