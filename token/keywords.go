package token

var keywords = map[string]Type{
	"f":        FUNCTION,
	"value":    VALUE,
	"true":     TRUE,
	"false":    FALSE,
	"if":       IF,
	"else":     ELSE,
	"return":   RETURN,
	"is":       EQ,
	"not":      NOT_EQ,
	"end":      CLOSEBLOCK,
	"update":   UPDATE,
	"constant": CONSTANT,
	"and":      AND,
	"or":       OR,
}

//LookupIdent finds an identifier token type from a string.
func LookupIdent(identifier string) Type {
	if tok, ok := keywords[identifier]; ok {
		return tok
	}

	return IDENT
}
