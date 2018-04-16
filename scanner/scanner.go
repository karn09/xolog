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
	case "!":
		if s.match("=") {
			s.addToken(token.BANG_EQUAL, nil)
		} else {
			s.addToken(token.BANG, nil)
		}
	case "=":
		if s.match("=") {
			s.addToken(token.EQUAL_EQUAL, nil)
		} else {
			s.addToken(token.EQUAL, nil)
		}
	case "<":
		if s.match("=") {
			s.addToken(token.LESS_EQUAL, nil)
		} else {
			s.addToken(token.LESS, nil)
		}
	case ">":
		if s.match("=") {
			s.addToken(token.GREATER_EQUAL, nil)
		} else {
			s.addToken(token.GREATER, nil)
		}
	case "\\":
		if s.match("\\") {
			for s.peek() != "\n" && !s.isAtEnd() {
				s.advance()
			}
			// advance one more time to place current cursor on \n
			s.advance()
		} else {
			s.addToken(token.SLASH, nil)
		}
	case " ":
	case "\t":
	case "\r":
	case "\n":
		s.line++
	default:
		error.Error(s.line, "Unexpected character: "+c)
	}
}

func (s *Scanner) match(expected string) bool {
	if s.isAtEnd() {
		return false
	}
	text := s.source[s.start+1 : s.current+1]
	if text != expected {
		return false
	}
	s.current++
	return true
}

func (s *Scanner) isAtEnd() bool {
	if s.current >= len(s.source) {
		return true
	}
	return false
}

func (s *Scanner) advance() string {
	s.current++
	return string(s.source[s.current-1])
}

func (s *Scanner) peek() string {
	if s.isAtEnd() {
		return "\000"
	}
	c := string(s.source[s.current+1])
	// fmt.Print(c)
	return c
}

func (s *Scanner) addToken(tokenType token.TokenType, literal interface{}) []token.Token {
	text := s.source[s.start:s.current]
	token := token.Token{Type: tokenType, Lexeme: text, Literal: literal, Line: s.line}
	s.tokens[s.current-1] = token
	return s.tokens
}
