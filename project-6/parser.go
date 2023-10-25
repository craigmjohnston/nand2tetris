package main

import (
	"strings"
	"unicode"
)

type Parser struct {
	source string
	pos    int

	HasMoreLines    bool
	InstructionType InstructionType
	Symbol          string
	Dest            string
	Comp            string
	Jump            string
}

type InstructionType int

const (
	AInstruction InstructionType = iota
	CInstruction
	LInstruction
)

func NewParser(source string) *Parser {
	return &Parser{
		source:       source,
		pos:          0,
		HasMoreLines: true,
	}
}

// Advance parses the next line of the source and sets the
// values on the Parser struct
func (p *Parser) Advance() {
	p.skipWhitespace()

	if p.peek(2) == "//" {
		p.skipLine()
	}

	if p.peek(1) == "@" {
		p.parseAInstruction()
	} else if p.peek(1) == "(" {
		p.parseLInstruction()
	} else {
		p.parseCInstruction()
	}

	if p.pos == len(p.source) {
		p.HasMoreLines = false
	}
}

func (p *Parser) parseAInstruction() {
	p.InstructionType = AInstruction
	p.skip(1) // @
	p.Symbol = p.takeLine()
}

func (p *Parser) parseLInstruction() {
	p.InstructionType = LInstruction
	p.skip(1) // (
	line := p.takeLine()

	if line[len(line)-1:] != ")" {
		panic("no closing brace for L-instruction")
	}

	p.Symbol = line[:len(line)-1]
}

func (p *Parser) parseCInstruction() {
	p.InstructionType = CInstruction
	line := p.takeLine()
	start := 0
	for i := 0; i < len(line); i++ {
		if unicode.IsSpace(rune(line[i])) {
			continue
		}

		if line[i] == '=' {
			if p.Dest != "" {
				panic("dest already defined")
			}

			p.Dest = trimSpace(line[start:i])
			start = i + 1
		} else if line[i] == ';' {
			if p.Comp != "" {
				panic("comp already defined")
			}

			p.Comp = trimSpace(line[start:i])
			start = i + 1
		}
	}

	if start < len(line)-1 {
		if p.Comp != "" {
			p.Jump = trimSpace(line[start:len(line)])
		} else {
			p.Comp = trimSpace(line[start:len(line)])
		}
	}
}

func trimSpace(value string) string {
	return strings.TrimFunc(value, unicode.IsSpace)
}

// peek returns a slice of a specified length from the current
// position in the source, but does not advance the position
func (p *Parser) peek(length int) string {
	return p.source[p.pos : p.pos+length]
}

// skip advances the position by the specified length
func (p *Parser) skip(length int) {
	p.pos += length
}

func (p *Parser) skipLine() {
	p.skipUntil("\n")
}

func (p *Parser) skipUntil(match string) bool {
	matched := 0
	for ; p.pos < len(p.source) && matched < len(match); p.pos++ {
		if p.source[p.pos] == match[matched] {
			matched += 1
		} else {
			matched = 0
		}
	}

	return matched == len(match)
}

// take returns a slice of a specified length from the current
// position in the source, advancing the position
func (p *Parser) take(length int) string {
	output := p.source[p.pos : p.pos+length]
	p.pos += length
	return output
}

// takeLine consumes the source from the current position
// until the end of the line, advancing the position
func (p *Parser) takeLine() string {
	length := 0
	for ; p.pos+length < len(p.source) && p.source[p.pos+length] != '\n'; length++ {
	}
	return strings.TrimRightFunc(p.take(length), unicode.IsSpace)
}

// skipWhitespace advances the position until it reaches a
// non-whitespace character
func (p *Parser) skipWhitespace() {
	for ; p.pos < len(p.source); p.pos++ {
		if !unicode.IsSpace([]rune(p.source)[p.pos]) {
			break
		}
	}
}
