package main

import (
	"bufio"
	"fmt"
	"os"

	"mini-lexer/lexer"
)

func main() {
	// 读取源文件
	sourceCode, err := os.ReadFile("source.txt")
	if err != nil {
		fmt.Println("Error reading source file:", err)
		return
	}

	// 创建词法分析器
	l := lexer.NewLexer(string(sourceCode))

	// 创建输出文件
	outputFile, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outputFile.Close()

	writer := bufio.NewWriter(outputFile)

	// 词法分析过程
	fmt.Fprintf(writer, "=== Lexical Analysis Results  ===\n")
	for {
		token := l.NextToken()
		lexer.PrintToken(writer, token)

		if token.Type == lexer.EOF {
			break
		}
	}

	// 输出各类单词表
	l.PrintTables(writer)
	writer.Flush()

	fmt.Println("Lexical analysis completed. Results saved to output.txt")
}
