package main

import (
	"fmt"
	"mini-parser/parser" // 修改后
	"mini-parser/token"  // 修改后
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("使用方法: mini_parser <文件路径>")
		os.Exit(1)
	}

	// 读取输入文件
	filename := os.Args[1]
	input, err := os.ReadFile(filename) // Replace ioutil.ReadFile with os.ReadFile
	if err != nil {
		fmt.Printf("读取文件错误: %v\n", err)
		os.Exit(1)
	}

	// 初始化词法分析器和语法分析器
	tokenizer := token.New(string(input))
	parser := parser.New(tokenizer)

	// 执行语法分析
	program := parser.ParseProgram()

	// 输出分析结果
	if len(parser.Errors()) > 0 {
		fmt.Println("语法分析发现错误:")
		for _, err := range parser.Errors() {
			fmt.Println(err)
		}
		os.Exit(1)
	}

	fmt.Println("语法分析成功! 程序结构:")
	fmt.Println(program.String())
}
