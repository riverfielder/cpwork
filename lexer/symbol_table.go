package lexer

// NewSymbolTable 初始化符号表
func NewSymbolTable() *SymbolTable {
	st := &SymbolTable{
		Keywords:    make(map[string]TokenType),
		Identifiers: make(map[string]int),
		Constants:   make(map[string]float64),
		Operators:   make(map[string]TokenType),
		Delimiters:  make(map[string]TokenType),
	}

	return st
}

// AddKeyword 动态添加关键字
func (st *SymbolTable) AddKeyword(keyword string, tokenType TokenType) {
	if _, exists := st.Keywords[keyword]; !exists {
		st.Keywords[keyword] = tokenType
	}
}

// AddOperator 动态添加运算符
func (st *SymbolTable) AddOperator(op string, tokenType TokenType) {
	if _, exists := st.Operators[op]; !exists {
		st.Operators[op] = tokenType
	}
}

// AddDelimiter 动态添加界符
func (st *SymbolTable) AddDelimiter(delim string, tokenType TokenType) {
	if _, exists := st.Delimiters[delim]; !exists {
		st.Delimiters[delim] = tokenType
	}
}

// AddIdentifier 添加标识符到符号表
func (st *SymbolTable) AddIdentifier(ident string) int {
	if index, exists := st.Identifiers[ident]; exists {
		return index // 已存在则返回现有索引
	}
	index := len(st.Identifiers)
	st.Identifiers[ident] = index
	return index
}

// GetIdentifierIndex 获取标识符索引
func (st *SymbolTable) GetIdentifierIndex(ident string) (int, bool) {
	index, exists := st.Identifiers[ident]
	return index, exists
}

// AddConstant 添加常量到符号表
func (st *SymbolTable) AddConstant(constant string, value float64) int {
	if index, exists := st.GetConstantIndex(constant); exists {
		return index // 已存在则返回现有索引
	}
	index := len(st.Constants)
	st.Constants[constant] = value
	return index
}

// GetConstantIndex 获取常量索引
func (st *SymbolTable) GetConstantIndex(constant string) (int, bool) {
	// 需要线性搜索因为常量是按值存储的
	i := 0
	for k, _ := range st.Constants {
		if k == constant {
			return i, true
		}
		i++
	}
	return -1, false
}

// GetConstantValue 获取常量值
func (st *SymbolTable) GetConstantValue(constant string) (float64, bool) {
	val, exists := st.Constants[constant]
	return val, exists
}

// IdentifierCount 返回标识符数量
func (st *SymbolTable) IdentifierCount() int {
	return len(st.Identifiers)
}

// ConstantCount 返回常量数量
func (st *SymbolTable) ConstantCount() int {
	return len(st.Constants)
}
