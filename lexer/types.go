package lexer

// TokenType 定义词法单元类型
type TokenType int

const (
	EOF TokenType = iota
	ERROR

	// 运算符
	PLUS     // +
	MINUS    // -
	MULTIPLY // *
	DIVIDE   // /
	ASSIGN   // =
	GT       // >
	LT       // <
	EQ       // ==

	// 界符
	LPAREN    // (
	RPAREN    // )
	SEMICOLON // ;

	// 关键字
	IF
	THEN
	ELSE

	// 其他类型
	IDENTIFIER // 标识符
	NUMBER     // 数字常量
	COMMENT    // 注释

)

// Token 结构体定义
type Token struct {
	Type   TokenType
	Lexeme string
	Value  interface{}
	Line   int
}

// 在 TokenType 定义后添加
func (tt TokenType) String() string {
	names := [...]string{
		"EOF", "ERROR",
		"PLUS", "MINUS", "MULTIPLY", "DIVIDE", "ASSIGN", "GT", "LT", "EQ",
		"LPAREN", "RPAREN", "SEMICOLON",
		"IF", "THEN", "ELSE",
		"IDENTIFIER", "NUMBER", "COMMENT",
	}
	if tt < 0 || int(tt) >= len(names) {
		return "UNKNOWN"
	}
	return names[tt]
}

// SymbolTable 各类单词表结构
type SymbolTable struct {
	Keywords    map[string]TokenType // 关键字表
	Identifiers map[string]int       // 标识符表
	Constants   map[string]float64   // 常数表
	Operators   map[string]TokenType // 运算符表
	Delimiters  map[string]TokenType // 界符表
}

// Lexer 词法分析器结构
type Lexer struct {
	input    string
	position int
	readPos  int
	ch       byte
	line     int
	symbols  *SymbolTable
}

// NumberValue 存储数字常量的值和索引
type NumberValue struct {
	Value float64
	Index int
}
