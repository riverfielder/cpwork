package parser

import "fmt"

type ParserError struct {
	Line    int
	Column  int
	Message string
}

func (e ParserError) Error() string {
	return fmt.Sprintf("语法错误: 第%d行第%d列 %s", e.Line, e.Column, e.Message)
}

type ParserErrors []ParserError

func (pe ParserErrors) Error() string {
	var out string
	for _, err := range pe {
		out += err.Error() + "\n"
	}
	return out
}
