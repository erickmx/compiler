package lexical

import (
	"testing"
)

func TestIsInt(t *testing.T) {
	matched := IsInt("10")
	if !matched {
		t.Error("The gived value 10 isn't an integer data type")
	}
	matched = IsInt("-1")
	if !matched {
		t.Error("The gived value -1 wasn't identified as an integer")
	}
	matched = IsInt("0")
	if !matched {
		t.Error("The gived value 0 isn't an integer data type")
	}
	matched = IsInt("+123")
	if !matched {
		t.Error("The gived value +123 wasn't identified as an integer data type")
	}
	matched = IsInt("-0")
	if !matched {
		t.Error("The gived value -0 isn't an integer data type")
	}
	matched = IsInt("+0")
	if !matched {
		t.Error("The gived value +0 isn't an integer data type")
	}
	matched = IsInt("abc")
	if matched {
		t.Error("The gived value abc was identified as an integer data type")
	}
	matched = IsInt("1.5")
	if matched {
		t.Error("The gived value 1.5 was identified as an integer data type")
	}
	matched = IsInt("")
	if matched {
		t.Error("The gived empty value was classified as an integer data type")
	}
	matched = IsInt("1 23 34")
	if matched {
		t.Error("The gived value 1 23 34 was classified as an integer data type")
	}
}

func TestIsDouble(t *testing.T) {
	matched := IsDouble("1.45")
	if !matched {
		t.Error("The gived value 1.45 isn't a double data type")
	}
	matched = IsDouble(".523")
	if !matched {
		t.Error("The gived value .523 isn't a double data type")
	}
	matched = IsDouble("0.678")
	if !matched {
		t.Error("The gived value 0.678 isn't a double data type")
	}
	matched = IsDouble("21")
	if !matched {
		t.Error("The gived value 21 isn't a double data type")
	}
	matched = IsDouble("0")
	if !matched {
		t.Error("The gived value 0 isn't a double data type")
	}
	matched = IsDouble("+1.45")
	if !matched {
		t.Error("The gived value +1.45 isn't a double data type")
	}
	matched = IsDouble("+.523")
	if !matched {
		t.Error("The gived value +.523 isn't a double data type")
	}
	matched = IsDouble("+0.678")
	if !matched {
		t.Error("The gived value +0.678 isn't a double data type")
	}
	matched = IsDouble("+21")
	if !matched {
		t.Error("The gived value +21 isn't a double data type")
	}
	matched = IsDouble("+0")
	if !matched {
		t.Error("The gived value 0 isn't a double data type")
	}
	matched = IsDouble("-1.45")
	if !matched {
		t.Error("The gived value -1.45 isn't a double data type")
	}
	matched = IsDouble("-.523")
	if !matched {
		t.Error("The gived value -.523 isn't a double data type")
	}
	matched = IsDouble("-0.678")
	if !matched {
		t.Error("The gived value -0.678 isn't a double data type")
	}
	matched = IsDouble("-21")
	if !matched {
		t.Error("The gived value -21 isn't a double data type")
	}
	matched = IsDouble("-0")
	if !matched {
		t.Error("The gived value 0 isn't a double data type")
	}
	matched = IsDouble("")
	if matched {
		t.Error("The gived empty value was classfified as a double data type")
	}
	matched = IsDouble("abc")
	if matched {
		t.Error("The gived value abc was identified as an double data type")
	}
	matched = IsDouble("1 23 34")
	if matched {
		t.Error("The gived value 1 23 34 was classified as a double data type")
	}
	matched = IsDouble("1 23.89")
	if matched {
		t.Error("The gived value 1 23.89 was classified as a double data type")
	}
}

func TestIsString(t *testing.T) {
	matched := IsString("\"hola\"")
	if !matched {
		t.Error("The gived string \"hola\" was not identified as an string data type")
	}
	matched = IsString("\"\"")
	if !matched {
		t.Error("The gived string \"\" was not identified as an empty string")
	}
	matched = IsString("\"123\"")
	if !matched {
		t.Error("The gived string \"123\" was not identified as an empty string")
	}
	matched = IsString("1234")
	if matched {
		t.Error("The gived expression was identified as an string data type")
	}
	matched = IsString("'hola'")
	if matched {
		t.Error("The gived value was identified as an string")
	}
	matched = IsString("\"(efmei#$#%.)mfm32423r\"")
	if !matched {
		t.Error("The gived value \"(efmei#$#%.)mfm32423r\" was not identified as an string data type")
	}
	matched = IsString("\"")
	if matched {
		t.Error("The gived value \" was identified as a string data type")
	}
	matched = IsString("hola")
	if matched {
		t.Error("The gived value hola was identified as a string data type")
	}
	matched = IsString("\"hello world!\"")
	if !matched {
		t.Error("The gived value \"hello world!\" was not identified as a string data type")
	}
	matched = IsString("\" \"")
	if !matched {
		t.Error("The gived value \" \" was not identified as a string data type")
	}
	matched = IsString("\"_\"")
	if !matched {
		t.Error("The gived value \"_\" was not identified as a string data type")
	}
	matched = IsString("\"_Hello world_. #1321\"")
	if !matched {
		t.Error("The gived value \"_Hello world_. #1321\" was not identified as a string data type")
	}
	matched = IsString("\"\\\"quotation\\\"\"")
	if !matched {
		t.Error("The gived value \"\\\"quotation\\\"\" was not identified as a string data type")
	}
}

func TestIsChar(t *testing.T) {
	matched := IsChar("'a'")
	if !matched {
		t.Error("The gived value 'a' was not identified as a char data type")
	}
	matched = IsChar("''")
	if !matched {
		t.Error("The gived value '' was not identified as a char data type")
	}
	matched = IsChar("'")
	if matched {
		t.Error("The gived value ' was identified as a char data type")
	}
	matched = IsChar("\"a\"")
	if matched {
		t.Error("The gived value \"a\" was identified as a char data type")
	}
	matched = IsChar("'sa'")
	if matched {
		t.Error("The gived value 'sa' was identified as a char data type")
	}
	matched = IsChar("'1'")
	if !matched {
		t.Error("The gived value '1' was not identified as an integer data type")
	}
	matched = IsChar("'('")
	if !matched {
		t.Error("The gived value '(' was not identified as a char data type")
	}
	matched = IsChar("' '")
	if !matched {
		t.Error("The gived value ' ' was not identified as a char data type")
	}
	matched = IsChar("'\"'")
	if !matched {
		t.Error("The gived value '\"' was not identified as a char data type")
	}
	matched = IsChar("'\\'")
	if !matched {
		t.Error("The gived value '\\' was not identified as a char data type")
	}
}

func TestIsFor(t *testing.T) {
	matched := IsFor("for")
	if !matched {
		t.Error("The gived value for was not identified as for reserved word")
	}
	matched = IsFor("For")
	if matched {
		t.Error("The gived value For was identified as for reserved word")
	}
	matched = IsFor("f or")
	if matched {
		t.Error("The gived value f or was identified as for reserved word")
	}
	matched = IsFor("f0r")
	if matched {
		t.Error("The gived value f0r was identified as for reserved word")
	}
	matched = IsFor("FOR")
	if matched {
		t.Error("The gived value FOR was identified as for reserved word")
	}
}

func TestIsIf(t *testing.T) {
	matched := IsIf("if")
	if !matched {
		t.Error("The gived value if was not identified as if reserved word")
	}
	matched = IsIf("If")
	if matched {
		t.Error("The gived value If was identified as if reserved word")
	}
	matched = IsIf("i f")
	if matched {
		t.Error("The gived value i f was identified as if reserved word")
	}
	matched = IsIf("1f")
	if matched {
		t.Error("The gived value 1f was identified as if reserved word")
	}
	matched = IsIf("IF")
	if matched {
		t.Error("The gived value IF was identified as if reserved word")
	}
}

func TestIsElse(t *testing.T) {
	matched := IsElse("else")
	if !matched {
		t.Error("The gived value else was not identified as else reserved word")
	}
	matched = IsElse("Else")
	if matched {
		t.Error("The gived value Else was identified as else reserved word")
	}
	matched = IsElse("e lse")
	if matched {
		t.Error("The gived value e lse was identified as else reserved word")
	}
	matched = IsElse("el5e")
	if matched {
		t.Error("The gived value el5e was identified as else reserved word")
	}
	matched = IsElse("ELSE")
	if matched {
		t.Error("The gived value ELSE was identified as else reserved word")
	}
}

func TestIsBreak(t *testing.T) {
	matched := IsBreak("break")
	if !matched {
		t.Error("The gived value break was not identified as break reserved word")
	}
	matched = IsBreak("Break")
	if matched {
		t.Error("The gived value Break was identified as break reserved word")
	}
	matched = IsBreak("b reak")
	if matched {
		t.Error("The gived value b reak was identified as break reserved word")
	}
	matched = IsBreak("br34k")
	if matched {
		t.Error("The gived value br34k was identified as break reserved word")
	}
	matched = IsBreak("BREAK")
	if matched {
		t.Error("The gived value BREAK was identified as break reserved word")
	}
}
