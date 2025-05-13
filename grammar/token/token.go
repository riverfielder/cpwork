package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
	Line    int
	Column  int
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"
	IDENT   = "IDENT"
	NUMBER  = "NUMBER"

	// 运算符
	ASSIGN   = ":="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"
	LT       = "<"
	GT       = ">"
	EQ       = "="
	NEQ      = "!="
	AND      = "&&"
	OR       = "||"

	// 分隔符
	COMMA     = ","
	SEMICOLON = ";"
	DOT       = "."
	LPAREN    = "("
	RPAREN    = ")"

	// 关键字
	PROGRAM = "program"
	BEGIN   = "begin"
	END     = "end"
	IF      = "if"
	THEN    = "then"
	ELSE    = "else"
	WHILE   = "while"
	DO      = "do"
	TRUE    = "true"
	FALSE   = "false"
)

var keywords = map[string]TokenType{
	"program": PROGRAM,
	"begin":   BEGIN,
	"end":     END,
	"if":      IF,
	"then":    THEN,
	"else":    ELSE,
	"while":   WHILE,
	"do":      DO,
	"true":    TRUE,
	"false":   FALSE,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
