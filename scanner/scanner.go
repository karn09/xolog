package scanner

import (
	"unicode"
	"unicode/utf8"
	"xolog/error"
	"xolog/token"
)

type Scanner struct {
	source   []byte
	start    int
	current  int
	line     int
	tokens   []token.Token
	HadError bool
}

// NewScanner accepts a string, and returns a pointer to the initialized Scanner struct.
func NewScanner(source string) *Scanner {
	return &Scanner{source: []byte(source), start: 0, current: 0, line: 1, tokens: []token.Token{}}
}

// ScanTokens will return an array of source length tokens of type Token.
func (s *Scanner) ScanTokens() []token.Token {
	for s.current < len(s.source) {
		s.start = s.current
		s.scanToken()
	}
	eof := token.Token{Type: token.EOF, Lexeme: string('\000'), Literal: nil, Line: s.line}
	s.tokens = append(s.tokens, eof)
	return s.tokens
}

func (s *Scanner) scanToken() {
	c := s.advance()
	switch c {
	case '(':
		s.addToken(token.LEFT_PAREN, nil)
	case ')':
		s.addToken(token.RIGHT_PAREN, nil)
	case '{':
		s.addToken(token.LEFT_BRACE, nil)
	case '}':
		s.addToken(token.RIGHT_BRACE, nil)
	case ',':
		s.addToken(token.COMMA, nil)
	case '.':
		s.addToken(token.DOT, nil)
	case '-':
		s.addToken(token.MINUS, nil)
	case '+':
		s.addToken(token.PLUS, nil)
	case ';':
		s.addToken(token.SEMICOLON, nil)
	case '*':
		s.addToken(token.STAR, nil)
	case '!':
		if s.match('=') {
			s.addToken(token.BANG_EQUAL, nil)
		} else {
			s.addToken(token.BANG, nil)
		}
	case '=':
		if s.match('=') {
			s.addToken(token.EQUAL_EQUAL, nil)
		} else {
			s.addToken(token.EQUAL, nil)
		}
	case '<':
		if s.match('=') {
			s.addToken(token.LESS_EQUAL, nil)
		} else {
			s.addToken(token.LESS, nil)
		}
	case '>':
		if s.match('=') {
			s.addToken(token.GREATER_EQUAL, nil)
		} else {
			s.addToken(token.GREATER, nil)
		}
	case '\\':
		if s.match('\\') {
			for s.peek() != '\n' && !s.isAtEnd() {
				s.advance()
			}
		} else {
			s.addToken(token.SLASH, nil)
		}
	case ' ':
	case '\t':
	case '\r':
	case '\n':
		s.line++
	case '"':
		s.string()
	case '\'':
		s.string()
	default:
		if s.isDigit(c) {
			s.number()
		} else {
			error.Error(s.line, "Unexpected character: "+string(c))
			s.HadError = true
		}
	}
}

// match will compare unconsumed character with expected character
func (s *Scanner) match(expected rune) bool {
	if s.isAtEnd() {
		return false
	}

	current, size := utf8.DecodeRune(s.source[s.current:])
	if current != expected {
		return false
	}
	s.current = s.current + size
	return true
}

// string will consume all tokens contained within " or ', add string token, with literal string.
func (s *Scanner) string() {
	literal := make([]rune, 0)
	for s.peek() != '"' && s.peek() != '\'' && !s.isAtEnd() {
		if s.peek() == '\n' {
			s.line++
		}
		literal = append(literal, s.advance())
	}
	if s.isAtEnd() {
		error.Error(s.line, "Unterminated string.")
		return
	}
	s.addToken(token.STRING, literal)
}

// isAtEnd will return whether current is last.
func (s *Scanner) isAtEnd() bool {
	if s.current >= len(s.source) {
		return true
	}
	return false
}

// advance will consume the current token, and return the consumed token.
func (s *Scanner) advance() rune {
	r, size := utf8.DecodeRune(s.source[s.current:])
	s.current = s.current + size
	return r
}

// peek will return the current token, without consuming
func (s *Scanner) peek() rune {
	if s.isAtEnd() {
		return '\000'
	}
	r, _ := utf8.DecodeRune(s.source[s.current:])
	return r
}

// peekNext will return the next token without consuming.
func (s *Scanner) peekNext() rune {
	if s.current+1 >= len(s.source) {
		return '\000'
	}
	_, size := utf8.DecodeRune(s.source[s.current:])
	nextRune, _ := utf8.DecodeRune(s.source[s.current+size:])
	return nextRune
}

// isDigit will check whether character is an integer.
func (s *Scanner) isDigit(c rune) bool {
	return unicode.IsDigit(c)
}

// number will consume, and finally add a number token, with float64 literal.
func (s *Scanner) number() {
	for s.isDigit(s.peek()) {
		s.advance()
	}
	if s.peek() == '.' && s.isDigit(s.peekNext()) {
		s.advance()
		for s.isDigit(s.peek()) {
			s.advance()
		}
	}
	// this seems a bit convoluted just to cast []byte into []rune
	literal := []rune(string(s.source[s.start:s.current]))
	s.addToken(token.NUMBER, literal)
}

// addToken will add token type and lexeme to returned tokens array.
func (s *Scanner) addToken(tokenType token.TokenType, literal []rune) []token.Token {
	var lexeme string
	if literal != nil {
		lexeme = string(literal)
	} else {
		lexeme = string(s.source[s.start:s.current])
	}
	token := token.Token{Type: tokenType, Lexeme: lexeme, Literal: literal, Line: s.line}
	s.tokens = append(s.tokens, token)
	return s.tokens
}
