package lexical

import "regexp"

// 	lexical reserved words
const (
	INT          = "^[-+]?\\d+$"
	DOUBLE       = "^[-+]?\\d*\\.?\\d+$" // supporting exponential numbers ^(\d*(\.\d*)?(e[+-]\d+)?)$
	STRING       = "\"[\\w\\d\\W\\D]*\""
	CHAR         = "'[\\w\\d\\W\\D]?'"
	FOR          = "for"
	IF           = "if"
	ELSE         = "else"
	BREAK        = "break"
	COLON        = ";"
	OPENBRACKET  = "\\["
	CLOSEBRACKET = "\\]"
	OPENCURLY    = "\\{"
	CLOSECURLY   = "\\}"
	SWITCH       = "switch"
	FUNC         = "func"
	RETURN       = "return"
	CASE         = "case"
	BOOL         = "bool"
	TRUE         = "true"
	FALSE        = "false"
	WHILE        = "while"
	PLUS         = "\\+"
	MINUS        = "\\-"
	MULT         = "^\\*{1}$"
	DIV          = "\\/"
	MOD          = "\\%"
	NULL         = "null"
	AND          = "&&"
	OR           = "\\|\\|"
	SEMICOLON    = ":"
	POW          = "^\\*{2}$"
	ASSIGNMENT   = "\\="
)

// IsInt verifies if the lexical expression passed in its argument is an integer
func IsInt(expression string) bool {
	matched, err := regexp.MatchString(INT, expression)
	if err != nil {
		return false
	}
	return matched
}

// IsDouble if the lexical expression passed in its argument is a double
func IsDouble(expression string) bool {
	matched, err := regexp.MatchString(DOUBLE, expression)
	if err != nil {
		return false
	}
	return matched
}

// IsString if the lexical expression passed in its arguments is a string
func IsString(expression string) bool {
	matched, err := regexp.MatchString(STRING, expression)
	if err != nil {
		return false
	}
	return matched
}

// IsChar if the lexical expression passed in its arguments is a char
func IsChar(expression string) bool {
	matched, err := regexp.MatchString(CHAR, expression)
	if err != nil {
		return false
	}
	return matched
}

// IsFor if the lexical expression passed in its arguments is the word for return true, else false
func IsFor(expression string) bool {
	matched, err := regexp.MatchString(FOR, expression)
	if err != nil {
		return false
	}
	return matched
}

// IsIf if the lexical expression passed in its arguments is the word if return true, else false
func IsIf(expression string) bool {
	matched, err := regexp.MatchString(IF, expression)
	if err != nil {
		return false
	}
	return matched
}

// IsElse if the lexical expression passed in its arguments is the word else return true, else false
func IsElse(expression string) bool {
	matched, err := regexp.MatchString(ELSE, expression)
	if err != nil {
		return false
	}
	return matched
}

// IsBreak if the lexical expression passed in its arguments is the word break return true, else false
func IsBreak(expression string) bool {
	matched, err := regexp.MatchString(BREAK, expression)
	if err != nil {
		return false
	}
	return matched
}

// IsColon if the lexical expression passed in its arguments is the identifier ; return true, else false
func IsColon(expression string) bool {
	matched, err := regexp.MatchString(COLON, expression)
	if err != nil {
		return false
	}
	return matched
}

// IsOpenBracket if the lexical expression passed in its arguments is the identifier [ return true, else false
func IsOpenBracket(expression string) bool {
	matched, err := regexp.MatchString(OPENBRACKET, expression)
	if err != nil {
		return false
	}
	return matched
}

// IsCloseBracket if the lexical expression passed in its arguments is the identifier ] return true, else false
func IsCloseBracket(expression string) bool {
	matched, err := regexp.MatchString(CLOSEBRACKET, expression)
	if err != nil {
		return false
	}
	return matched
}

// IsOpenCurly if the lexical expression passed in its arguments is the identifier { return true, else false
func IsOpenCurly(expression string) bool {
	matched, err := regexp.MatchString(OPENCURLY, expression)
	if err != nil {
		return false
	}
	return matched
}

// IsCloseCurly if the lexical expression passed in its arguments is the identifier } return true, else false
func IsCloseCurly(expression string) bool {
	matched, err := regexp.MatchString(CLOSECURLY, expression)
	if err != nil {
		return false
	}
	return matched
}

// IsSwitch if the lexical expression passed in its arguments is the word switch return true, else false
func IsSwitch(expression string) bool {
	matched, err := regexp.MatchString(SWITCH, expression)
	if err != nil {
		return false
	}
	return matched
}

// IsFunc if the lexical expression passed in its arguments is the word func return true, else false
func IsFunc(expression string) bool {
	matched, err := regexp.MatchString(FUNC, expression)
	if err != nil {
		return false
	}
	return matched
}

// IsReturn if the lexical expression passed in its arguments is the word return return true, else false
func IsReturn(expression string) bool {
	matched, err := regexp.MatchString(FOR, expression)
	if err != nil {
		return false
	}
	return matched
}

// IsCase if the lexical expression passed in its arguments is the word case return true, else false
func IsCase(expression string) bool {
	matched, err := regexp.MatchString(CASE, expression)
	if err != nil {
		return false
	}
	return matched
}

// IsBool if the lexical expression passed in its arguments is the word bool return true, else false
func IsBool(expression string) bool {
	matched, err := regexp.MatchString(BOOL, expression)
	if err != nil {
		return false
	}
	return matched
}

// IsTrue if the lexical expression passed in its arguments is the word true return true, else false
func IsTrue(expression string) bool {
	matched, err := regexp.MatchString(TRUE, expression)
	if err != nil {
		return false
	}
	return matched
}

// IsFalse if the lexical expression passed in its arguments is the word false return true, else false
func IsFalse(expression string) bool {
	matched, err := regexp.MatchString(FALSE, expression)
	if err != nil {
		return false
	}
	return matched
}

// IsWhile if the lexical expression passed in its arguments is the word while return true, else false
func IsWhile(expression string) bool {
	matched, err := regexp.MatchString(FOR, expression)
	if err != nil {
		return false
	}
	return matched
}

// IsPlus if the lexical expression passed in its arguments is the symbol + return true, else false
func IsPlus(expression string) bool {
	matched, err := regexp.MatchString(PLUS, expression)
	if err != nil {
		return false
	}
	return matched
}

// IsMinus if the lexical expression passed in its arguments is the symbol - return true, else false
func IsMinus(expression string) bool {
	matched, err := regexp.MatchString(MINUS, expression)
	if err != nil {
		return false
	}
	return matched
}

// IsMult if the lexical expression passed in its arguments is the symbol * return true, else false
func IsMult(expression string) bool {
	matched, err := regexp.MatchString(MULT, expression)
	if err != nil {
		return false
	}
	return matched
}

// IsDiv if the lexical expression passed in its arguments is the symbol / return true, else false
func IsDiv(expression string) bool {
	matched, err := regexp.MatchString(DIV, expression)
	if err != nil {
		return false
	}
	return matched
}

// IsMod if the lexical expression passed in its arguments is the symbol % return true, else false
func IsMod(expression string) bool {
	matched, err := regexp.MatchString(MOD, expression)
	if err != nil {
		return false
	}
	return matched
}

// IsNull if the lexical expression passed in its arguments is the word null return true, else false
func IsNull(expression string) bool {
	matched, err := regexp.MatchString(NULL, expression)
	if err != nil {
		return false
	}
	return matched
}

// IsAnd if the lexical expression passed in its arguments is the symbol && return true, else false
func IsAnd(expression string) bool {
	matched, err := regexp.MatchString(AND, expression)
	if err != nil {
		return false
	}
	return matched
}

// IsOr if the lexical expression passed in its arguments is the symbol || return true, else false
func IsOr(expression string) bool {
	matched, err := regexp.MatchString(OR, expression)
	if err != nil {
		return false
	}
	return matched
}

// IsSemiColon if the lexical expression passed in its arguments is the symbol : return true, else false
func IsSemiColon(expression string) bool {
	matched, err := regexp.MatchString(SEMICOLON, expression)
	if err != nil {
		return false
	}
	return matched
}

// IsPow if the lexical expression passed in its arguments is the operand ** return true, else false
func IsPow(expression string) bool {
	matched, err := regexp.MatchString(POW, expression)
	if err != nil {
		return false
	}
	return matched
}

// IsAssignment if the lexical expression passed in its arguments is the operand = return true, else false
func IsAssignment(expression string) bool {
	matched, err := regexp.MatchString(ASSIGNMENT, expression)
	if err != nil {
		return false
	}
	return matched
}

// Classify return the string representation of the gived value
func Classify(expression string) string {
	if IsInt(expression) {
		return "int"
	}
	if IsDouble(expression) {
		return "double"
	}
	if IsString(expression) {
		return "string"
	}
	if IsChar(expression) {
		return "char"
	}
	if IsFor(expression) {
		return FOR
	}
	if IsIf(expression) {
		return IF
	}
	if IsElse(expression) {
		return ELSE
	}
	if IsBreak(expression) {
		return BREAK
	}
	if IsColon(expression) {
		return COLON
	}
	if IsOpenBracket(expression) {
		return "["
	}
	if IsCloseBracket(expression) {
		return "]"
	}
	if IsOpenCurly(expression) {
		return "{"
	}
	if IsCloseCurly(expression) {
		return "}"
	}
	if IsSwitch(expression) {
		return SWITCH
	}
	if IsFunc(expression) {
		return FUNC
	}
	if IsReturn(expression) {
		return RETURN
	}
	if IsCase(expression) {
		return CASE
	}
	if IsBool(expression) {
		return BOOL
	}
	if IsTrue(expression) {
		return TRUE
	}
	if IsFalse(expression) {
		return FALSE
	}
	if IsWhile(expression) {
		return WHILE
	}
	if IsPlus(expression) {
		return "+"
	}
	if IsMinus(expression) {
		return "-"
	}
	if IsMult(expression) {
		return "*"
	}
	if IsDiv(expression) {
		return "/"
	}
	if IsMod(expression) {
		return "%"
	}
	if IsNull(expression) {
		return NULL
	}
	if IsAnd(expression) {
		return "&&"
	}
	if IsOr(expression) {
		return "||"
	}
	if IsSemiColon(expression) {
		return SEMICOLON
	}
	if IsPow(expression) {
		return "**"
	}
	if IsAssignment(expression) {
		return "="
	}
	return "IDENTIFIER"
}
