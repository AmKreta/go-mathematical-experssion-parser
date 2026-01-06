package lexer

func New(input string) *Lexer {
	return &Lexer{input: input, position: 0, inputLength: len(input)}
}
