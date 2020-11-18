package token

var keywords = map[string]Type{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

//LookupIdent finds an identifier token type from a string.
func LookupIdent(identifier string) Type {
	if tok, ok := keywords[identifier]; ok {
		return tok
	}

	return IDENT
}
