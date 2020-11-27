package token

//Type is a token type
type Type string

const (
	// Endings
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"
	// Identifiers + literals
	IDENT = "IDENT" // add, foobar, x, y, ...
	INT   = "INT"   // 1343456
	FLOAT = "FLOAT"
	// Operators
	ASSIGN = "="
	PLUS   = "+"
	EQ     = "=="
	NOT_EQ = "!="
	// Delimiters
	COMMA       = ","
	SEMICOLON   = ";"
	LPAREN      = "("
	RPAREN      = ")"
	LBRACE      = "{"
	RBRACE      = "}"
	LBRACKET    = "["
	RBRACKET    = "]"
	COLON       = ":"
	DOUBLECOLON = "::"
	BAR         = "|"
	// Keywords
	FUNCTION   = "FUNCTION"
	VALUE      = "VALUE"
	UPDATE     = "UPDATE"
	CONSTANT   = "CONSTANT"
	TRUE       = "TRUE"
	FALSE      = "FALSE"
	IF         = "IF"
	ELSE       = "ELSE"
	RETURN     = "RETURN"
	STRING     = "STRING"
	MATCH      = "MATCH"
	MINUS      = "-"
	BANG       = "!"
	ASTERISK   = "*"
	SLASH      = "/"
	LT         = "<"
	GT         = ">"
	LTEQ       = "<="
	GTEQ       = ">="
	INCREMENT  = "++"
	DECREMENT  = "--"
	MINUSEQ    = "-="
	PLUSEQ     = "+="
	MULEQ      = "*="
	SLASHEQ    = "/="
	OR         = "||"
	AND        = "&&"
	OPENBLOCK  = "->"
	CLOSEBLOCK = "end"
	DEFAULT    = "default"
)
