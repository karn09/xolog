package scanner

import (
	"xolog/error"
	"xolog/token"
)

// type tokens []token.Token

type Scanner struct {
	source  string
	start   int
	current int
	line    int
	tokens  []token.Token
}

func NewScanner(source string) *Scanner {
	return &Scanner{source: source, start: 0, current: 0, line: 1, tokens: make([]token.Token, len(source)+1)}
}

func (s *Scanner) scanTokens() []token.Token {
	for s.current < len(s.source) {
		s.start = s.current
		s.scanToken()
	}
	eof := token.Token{Type: token.EOF, Lexeme: "", Literal: nil, Line: s.line}
	s.tokens[s.current] = eof
	return s.tokens
}

func (s *Scanner) scanToken() {
	c := s.advance()
	switch c {
	case "(":
		s.addToken(token.LEFT_PAREN, nil)
	case ")":
		s.addToken(token.RIGHT_PAREN, nil)
	case "{":
		s.addToken(token.LEFT_BRACE, nil)
	case "}":
		s.addToken(token.RIGHT_BRACE, nil)
	case ",":
		s.addToken(token.COMMA, nil)
	case ".":
		s.addToken(token.DOT, nil)
	case "-":
		s.addToken(token.MINUS, nil)
	case "+":
		s.addToken(token.PLUS, nil)
	case ";":
		s.addToken(token.SEMICOLON, nil)
	case "*":
		s.addToken(token.STAR, nil)
	default:
		error.Error(s.line, "Unexpected character.")
	}
}

func (s *Scanner) advance() string {
	s.current++
	return string(s.source[s.current-1])
}

func (s *Scanner) addToken(tokenType token.TokenType, literal interface{}) []token.Token {
	text := s.source[s.start:s.current]
	token := token.Token{Type: tokenType, Lexeme: text, Literal: literal, Line: s.line}
	s.tokens[s.current-1] = token
	return s.tokens
}