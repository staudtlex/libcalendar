// Copyright (C) 2022  Alexander Staudt
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

// This file implements a minimal JSON parser sufficient to unmarshal
// JSON-serialized Date objects

package libcalendar

import (
	"bufio"
	"bytes"
	"io"
	"strconv"
	"strings"
)

// Tokens
type token int

const (
	// Special tokens
	ILLEGAL token = iota
	EOF
	WS

	// Literals
	NAME_OR_STRING
	CALENDAR
	COMPONENTS
	ARRAY
	DIGITS

	// Misc characters
	COMMA // ,
	COLON // :
)

// End of file rune
const eof = rune(0)

// isWhitespace returns true if a rune is either a blank, a tab, or a newline
// character, and returns false otherwise.
func isWhitespace(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\n'
}

// isLetter returns true if a rune is a lower- of uppercase latin character,
// and false otherwise.
func isLetter(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') ||
		(ch >= 'A' && ch <= 'Z') ||
		(ch >= 'À' && ch <= 'ÿ') ||
		(ch == '"')
}

// isDigit returns true if a rune represents a digit, and false otherwise.
func isDigit(ch rune) bool {
	return ch >= '0' && ch <= '9'
}

// Scanner represents a lexical scanner.
type Scanner struct {
	reader *bufio.Reader
}

// NewScanner returns a new instance of Scanner.
func NewScanner(reader io.Reader) *Scanner {
	return &Scanner{reader: bufio.NewReader(reader)}
}

// read reads the next rune from the buffered reader.
// Returns the rune(0) if an error occurs (or io.EOF is returned).
func (s *Scanner) read() rune {
	ch, _, err := s.reader.ReadRune()
	if err != nil {
		return eof
	}
	return ch
}

// unread places the previously read rune back on the reader.
func (s *Scanner) unread() {
	_ = s.reader.UnreadRune()
}

// scanWhitespace consumes the current rune and all contiguous whitespace.
func (s *Scanner) scanWhitespace() (tok token, str string) {
	// Create a buffer and read the current character from s into it.
	var buf bytes.Buffer
	buf.WriteRune(s.read())
	// Read every subsequent whitespace character into the buffer.
	// Exit the loop upon finding a non-whitespace characters or EOF.
	for {
		if ch := s.read(); ch == eof {
			break
		} else if !isWhitespace(ch) {
			s.unread()
			break
		} else {
			buf.WriteRune(ch)
		}
	}
	return WS, buf.String()
}

// skipWhitespace advances the reader whenever a character is a whitespace.
func (s *Scanner) skipWhitespace() {
	for {
		if ch := s.read(); !isWhitespace(ch) {
			s.unread()
			break
		}
	}
}

// scanArray returns a token representing a JSON-array.
func (s *Scanner) scanArray() (tok token, str string) {
	var buf bytes.Buffer
	s.skipWhitespace()
	for {
		if ch := s.read(); ch == ']' {
			break
		} else {
			if ch != '\t' && ch != '\n' {
				buf.WriteRune(ch)
			}
		}
	}
	return ARRAY, buf.String()
}

// scanDigits returns a token representing digits.
func (s *Scanner) scanDigits() (tok token, str string) {
	var buf bytes.Buffer
	for {
		if ch := s.read(); ch == eof {
			break
		} else if !isDigit(ch) {
			s.unread()
			break
		} else {
			_, _ = buf.WriteRune(ch)
		}
	}
	return DIGITS, buf.String()
}

// scanName consumes the current rune and all contiguous name runes.
func (s *Scanner) scanName() (tok token, str string) {
	var buf bytes.Buffer
	for {
		if ch := s.read(); ch == eof {
			s.unread()
			break
		} else if !isLetter(ch) && !isDigit(ch) && ch != ' ' {
			s.unread()
			break
		} else {
			_, _ = buf.WriteRune(ch)
		}
	}
	nameOrString := strings.ReplaceAll(buf.String(), "\"", "")
	switch nameOrString {
	case "calendar":
		return CALENDAR, nameOrString
	case "components":
		return COMPONENTS, nameOrString
	default:
		return NAME_OR_STRING, nameOrString
	}
}

// Scan returns the next token and literal value.
func (s *Scanner) Scan() (tok token, str string) {
	// Read the next rune.
	ch := s.read()
	// If we see whitespace then consume all contiguous whitespace.
	// If we see a letter then consume as an ident or reserved word.
	if isWhitespace(ch) {
		s.unread()
		return s.scanWhitespace()
	} else if ch == '[' {
		return s.scanArray()
	} else if isLetter(ch) {
		s.unread()
		return s.scanName()
	} else if isDigit(ch) {
		s.unread()
		return s.scanDigits()
	}
	// Otherwise read the individual character.
	switch ch {
	case eof:
		return EOF, ""
	case ',':
		return COMMA, string(ch)
	case ':':
		return COLON, string(ch)
	}
	return ILLEGAL, string(ch)
}

type tokenString struct {
	token
	string
}

// tokenize tokenizes a string and returns all elements representing either a
// JSON key, string value, or array.
func tokenize(s string) []tokenString {
	scan := NewScanner(strings.NewReader(s))
	tokens := []tokenString{}
	for tok, str := scan.Scan(); tok != EOF; tok, str = scan.Scan() {
		if tok == NAME_OR_STRING || tok == CALENDAR || tok == COMPONENTS || tok == ARRAY {
			tokens = append(tokens, tokenString{tok, str})
		}
	}
	return tokens
}

// parseNumArray parses a string representing a JSON array of numeric values
// and returns a slice of float64.
func parseNumArray(s string) []float64 {
	scan := NewScanner(strings.NewReader(s))
	array := []float64{}
	for tok, str := scan.Scan(); tok != EOF; tok, str = scan.Scan() {
		if tok == DIGITS {
			num, _ := strconv.ParseFloat(str, 64)
			array = append(array, num)
		}
	}
	return array
}

// isValidCalendar returns true if the given string matches the name of a
// supported calendar, and false otherwise.
func isValidCalendar(calendar string) bool {
	switch calendar {
	case "gregorian",
		"julian",
		"iso",
		"islamic",
		"hebrew",
		"mayanLongCount",
		"mayanHaab",
		"mayanTzolkin",
		"french",
		"oldHinduSolar",
		"oldHinduLunar":
		return true
	default:
		return false
	}
}

// JsonToDate unmarshals a JSON-serialized Date object into a Date struct.
// Unmarshals only elements "calendar" and "components".
func JsonToDate(json string) Date {
	tokens := tokenize(json)
	data := map[token]string{}
	for i := 0; i < len(tokens); i++ {
		switch {
		case tokens[i].token == CALENDAR:
			if isValidCalendar(tokens[i+1].string) {
				data[tokens[i].token] = tokens[i+1].string
			} else {
				data[tokens[i].token] = ""
			}
		case tokens[i].token == COMPONENTS:
			data[tokens[i].token] = tokens[i+1].string
		}
	}
	return Date{
		Calendar:   data[CALENDAR],
		Components: parseNumArray(data[COMPONENTS]),
	}
}
