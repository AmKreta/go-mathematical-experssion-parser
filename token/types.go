package token

type TokenType int8

const (
	// number
	NUMBER TokenType = iota
	PI

	// operators
	PLUS
	MINUS
	MULTIPLY
	DIVIDE
	MODULO
	INTEGER_DIVIDE
	POWER
	LOGARITHM
	FACTORIAL
	SINE
	COSINE
	TANGENT
	COTANGENT
	SECANT
	COSECANT

	// brackets
	PARENTHESIS_OPEN
	PARENTHESIS_CLOSE
	Square_Bracket_Open
	Square_Bracket_Close
	Curly_Bracket_Open
	Curly_Bracket_Close

	EOF
	ILLEGAL
	SPACE
)

var TokenNames = map[TokenType]string{
	NUMBER:               "NUMBER",
	PI:                   "PI",
	PLUS:                 "PLUS",
	MINUS:                "MINUS",
	MULTIPLY:             "MULTIPLY",
	DIVIDE:               "DIVIDE",
	MODULO:               "MODULO",
	INTEGER_DIVIDE:       "INTEGER_DIVIDE",
	POWER:                "POWER",
	LOGARITHM:            "LOGARITHM",
	FACTORIAL:            "FACTORIAL",
	SINE:                 "SINE",
	COSINE:               "COSINE",
	TANGENT:              "TANGENT",
	COTANGENT:            "COTANGENT",
	SECANT:               "SECANT",
	COSECANT:             "COSECANT",
	PARENTHESIS_OPEN:     "PARENTHESIS_OPEN",
	PARENTHESIS_CLOSE:    "PARENTHESIS_CLOSE",
	Square_Bracket_Open:  "Square_Bracket_Open",
	Square_Bracket_Close: "Square_Bracket_Close",
	Curly_Bracket_Open:   "Curly_Bracket_Open",
	Curly_Bracket_Close:  "Curly_Bracket_Close",
	EOF:                  "EOF",
	ILLEGAL:              "ILLEGAL",
	SPACE:                "SPACE",
}

var TokenTypeValMap = map[TokenType]string{
	NUMBER:               "NUMBER",
	PI:                   "PI",
	PLUS:                 "+",
	MINUS:                "-",
	MULTIPLY:             "*",
	DIVIDE:               "/",
	MODULO:               "%",
	INTEGER_DIVIDE:       "//",
	POWER:                "^",
	LOGARITHM:            "log",
	FACTORIAL:            "!",
	SINE:                 "sin",
	COSINE:               "cos",
	TANGENT:              "tan",
	COTANGENT:            "cot",
	SECANT:               "sec",
	COSECANT:             "csc",
	PARENTHESIS_OPEN:     "(",
	PARENTHESIS_CLOSE:    ")",
	Square_Bracket_Open:  "[",
	Square_Bracket_Close: "]",
	Curly_Bracket_Open:   "{",
	Curly_Bracket_Close:  "}",
	EOF:                  "EOF",
	ILLEGAL:              "ILLEGAL",
	SPACE:                "SPACE",
}
