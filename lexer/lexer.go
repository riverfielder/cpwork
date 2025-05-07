package lexer

import (
	"fmt"
	"strconv"
)

// 辅助函数
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

// NewLexer 创建新的词法分析器
func NewLexer(input string) *Lexer {
	l := &Lexer{
		input:   input,
		line:    1,
		symbols: NewSymbolTable(),
	}
	l.readChar()
	return l
}

// readChar 读取下一个字符
func (l *Lexer) readChar() {
	if l.readPos >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPos]
	}
	l.position = l.readPos
	l.readPos += 1
}

// peekChar 预读下一个字符
func (l *Lexer) peekChar() byte {
	if l.readPos >= len(l.input) {
		return 0
	}
	return l.input[l.readPos]
}

// skipWhitespace 跳过空白字符
func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		if l.ch == '\n' {
			l.line++
		}
		l.readChar()
	}
}

// skipComment 跳过注释
func (l *Lexer) skipComment() {
	// 处理单行注释 //
	if l.peekChar() == '/' {
		for l.ch != '\n' && l.ch != 0 {
			l.readChar()
		}
		return
	}

	// 处理多行注释 /* ... */
	if l.peekChar() == '*' {
		l.readChar() // 跳过第一个 *
		for {
			l.readChar()
			if l.ch == 0 { // 文件结束
				return
			}
			if l.ch == '*' && l.peekChar() == '/' {
				l.readChar() // 跳过 *
				l.readChar() // 跳过 /
				return
			}
		}
	}
}

// readNumber 读取数字
func (l *Lexer) readNumber() (string, error) {
	position := l.position
	dotCount := 0

	// 检查数字开头是否为0且后面还有数字（非法的八进制表示）
	if l.ch == '0' && l.peekChar() >= '0' && l.peekChar() <= '9' {
		// 读取整个非法数字直到非数字字符
		for isDigit(l.ch) {
			l.readChar()
		}
		return "", fmt.Errorf("illegal number format: leading zeros not allowed")
	}

	for {
		if isDigit(l.ch) {
			l.readChar()
		} else if l.ch == '.' {
			dotCount++
			if dotCount > 1 {
				return "", fmt.Errorf("illegal number format: multiple decimal points")
			}
			l.readChar()
			// 小数点后必须有数字
			if !isDigit(l.ch) {
				return "", fmt.Errorf("illegal number format: decimal point must be followed by digits")
			}
		} else {
			break
		}
	}

	// 检查数字后是否紧跟字母（非法标识符）
	if isLetter(l.ch) {
		return "", fmt.Errorf("illegal identifier: cannot start with number")
	}

	return l.input[position:l.position], nil
}

// readIdentifier 读取标识符
func (l *Lexer) readIdentifier() (string, error) {
	position := l.position

	// 第一个字符必须是字母或下划线
	if !isLetter(l.ch) {
		return "", fmt.Errorf("illegal identifier: must start with letter or underscore")
	}

	for isLetter(l.ch) || isDigit(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position], nil
}

// NextToken 获取下一个词法单元
func (l *Lexer) NextToken() Token {
	l.skipWhitespace()

	tok := Token{Line: l.line}

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			l.readChar()
			tok.Lexeme = "=="
			tok.Type = EQ
		} else {
			tok.Lexeme = string(l.ch)
			tok.Type = ASSIGN
			l.readChar()

			// 检查赋值符号后是否直接跟着分号或换行
			l.skipWhitespace()
			if l.ch == ';' || l.ch == '\n' || l.ch == 0 {
				return Token{
					Type:   ERROR,
					Lexeme: "",
					Value:  "missing expression after '='",
					Line:   l.line,
				}
			}
			return tok
		}
		l.readChar()
	case '+':
		tok = Token{Type: PLUS, Lexeme: string(l.ch), Line: l.line} // 添加行号
		l.readChar()
	case '-':
		tok = Token{Type: MINUS, Lexeme: string(l.ch), Line: l.line} // 添加行号
		l.readChar()
	case '*':
		tok = Token{Type: MULTIPLY, Lexeme: string(l.ch), Line: l.line} // 添加行号
		l.readChar()
	case '/':
		// 检查是否是注释
		if l.peekChar() == '/' || l.peekChar() == '*' {
			l.skipComment()
			return l.NextToken() // 递归调用获取下一个有效token
		}
		tok = Token{Type: DIVIDE, Lexeme: string(l.ch), Line: l.line} // 添加行号
		l.readChar()
	case '>':
		tok = Token{Type: GT, Lexeme: string(l.ch), Line: l.line}
		l.readChar()
	case '<':
		tok = Token{Type: LT, Lexeme: string(l.ch), Line: l.line}
		l.readChar()
	case '(':
		tok = Token{Type: LPAREN, Lexeme: string(l.ch), Line: l.line}
		l.readChar()
	case ')':
		tok = Token{Type: RPAREN, Lexeme: string(l.ch), Line: l.line}
		l.readChar()
	case ';':
		tok = Token{Type: SEMICOLON, Lexeme: string(l.ch), Line: l.line}
		l.readChar()
	case '#':
		tok.Type = EOF
		tok.Lexeme = "#"
		l.readChar()
	case 0:
		tok.Type = EOF
	default:
		if isLetter(l.ch) {
			// 处理标识符
			var err error
			tok.Lexeme, err = l.readIdentifier()
			if err != nil {
				tok.Type = ERROR
				tok.Value = err.Error()
				return tok
			}

			// 检查是否是关键字
			if keywordType, ok := l.symbols.Keywords[tok.Lexeme]; ok {
				tok.Type = keywordType
			} else {
				tok.Type = IDENTIFIER
				// 添加到标识符表并获取索引
				index := l.symbols.AddIdentifier(tok.Lexeme)
				tok.Value = index
			}
			return tok

		} else if isDigit(l.ch) {
			// 处理数字常量
			var err error
			tok.Lexeme, err = l.readNumber()
			if err != nil {
				tok.Type = ERROR
				tok.Value = err.Error()
				l.readChar() // 确保读取位置推进
				return tok
			}

			tok.Type = NUMBER
			// 将字符串转换为float64
			value, _ := strconv.ParseFloat(tok.Lexeme, 64)
			// 添加到常量表并获取索引
			index := l.symbols.AddConstant(tok.Lexeme, value)
			tok.Value = struct {
				Value float64
				Index int
			}{
				Value: value,
				Index: index,
			}
			return tok

		} else {
			// 非法字符处理
			tok.Type = ERROR
			tok.Lexeme = string(l.ch)
			tok.Value = "illegal character"
			l.readChar()
		}
	}

	return tok
}
