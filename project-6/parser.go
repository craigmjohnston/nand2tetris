package main

import (
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
	} else {
		p.parseCInstruction()
	}

	// todo L instruction ??

	if p.pos == len(p.source) {
		p.HasMoreLines = false
	}
}

func (p *Parser) parseAInstruction() {
	p.InstructionType = AInstruction
	p.skip(1) // @
	p.Symbol = p.takeLine()
}

func (p *Parser) parseCInstruction() {
	panic("implement")
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
	return p.take(length)
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
