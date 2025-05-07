package lexer

import (
	"bufio"
	"fmt"
)

// PrintTables 输出各类单词表
func (l *Lexer) PrintTables(writer *bufio.Writer) {
	// 1. 关键字表
	fmt.Fprintf(writer, "\n=== Keyword Table ===\n")
	for keyword := range l.symbols.Keywords {
		fmt.Fprintf(writer, "%s\n", keyword)
	}

	// 2. 标识符表
	fmt.Fprintf(writer, "\n=== Identifier Table ===\n")
	for ident, index := range l.symbols.Identifiers {
		fmt.Fprintf(writer, "%d: %s\n", index, ident)
	}

	// 3. 常数表
	fmt.Fprintf(writer, "\n=== Constant Table ===\n")
	for constant, value := range l.symbols.Constants {
		fmt.Fprintf(writer, "%f: %s\n", value, constant)
	}

	// 4. 运算符表
	fmt.Fprintf(writer, "\n=== Operator Table ===\n")
	for op := range l.symbols.Operators {
		fmt.Fprintf(writer, "%s\n", op)
	}

	// 5. 界符表
	fmt.Fprintf(writer, "\n=== Delimiter Table ===\n")
	for delim := range l.symbols.Delimiters {
		fmt.Fprintf(writer, "%s\n", delim)
	}
}

// PrintToken 增强版的词法单元输出函数
func PrintToken(writer *bufio.Writer, token Token) {
	fmt.Fprintf(writer, "Line %d: Type: %v, Lexeme: %s", token.Line, token.Type.String(), token.Lexeme)

	switch token.Type {
	case IDENTIFIER:
		if index, ok := token.Value.(int); ok {
			fmt.Fprintf(writer, ", Index: %d", index)
		}

	case NUMBER:
		if numVal, ok := token.Value.(NumberValue); ok {
			fmt.Fprintf(writer, ", Value: %f, Index: %d", numVal.Value, numVal.Index)
		}

	case ERROR:
		fmt.Fprintf(writer, " --- ERROR: %v", token.Value)

	case IF, THEN, ELSE:
		fmt.Fprintf(writer, " (keyword)")

	}

	fmt.Fprintln(writer) // 换行
}
