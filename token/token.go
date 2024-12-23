package token

type TokenType string

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDENT = "IDENT" // varaibles... foo, x, y?
	INT   = "INT"   //integer value

	//Operator
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"

	LESSER  = "<"
	GREATER = ">"
)

type Token struct {
	Type    TokenType
	Literal string
}

var keywords = map[string]TokenType{}
