package main

import (
	"fmt"
	"strconv"
	"unicode"
)

type Parser interface {
	HasMoreLines() bool
	Advance()
	CommandType() CommandType
	Arg1() string
	Arg2() int
}

type CommandType int

const (
	ArithmeticCommand CommandType = iota
	PushCommand
	PopCommand
	LabelCommand
	GotoCommand
	IfCommand
	FunctionCommand
	ReturnCommand
	CallCommand
)

var commandToType = map[string]CommandType{
	"add":  ArithmeticCommand,
	"sub":  ArithmeticCommand,
	"neg":  ArithmeticCommand,
	"eq":   ArithmeticCommand,
	"gt":   ArithmeticCommand,
	"lt":   ArithmeticCommand,
	"and":  ArithmeticCommand,
	"or":   ArithmeticCommand,
	"not":  ArithmeticCommand,
	"push": PushCommand,
	"pop":  PopCommand,
}

type parser struct {
	source      string
	sourceRunes []rune
	pos         int

	hasMoreLines bool
	commandType  CommandType
	arg1         string
	arg2         int
}

func NewParser(source string) Parser {
	return &parser{
		source:       source,
		sourceRunes:  []rune(source),
		hasMoreLines: len(source) > 0,
	}
}

func (p *parser) HasMoreLines() bool       { return p.hasMoreLines }
func (p *parser) CommandType() CommandType { return p.commandType }
func (p *parser) Arg1() string             { return p.arg1 }
func (p *parser) Arg2() int                { return p.arg2 }

func (p *parser) Advance() {
	// skip whitespace (and empty lines)
	p.skip(while(unicode.IsSpace))

	// skip comments
	for p.peekequals("//") {
		if !p.skip(until(IsNewline)) {
			p.end()
			return
		}
		p.pos += 1 // skip the newline

		// skip any whitespace which could be between comments
		p.skip(while(unicode.IsSpace))
	}

	// check if this is the end of the file
	if p.pos == len(p.sourceRunes) {
		p.end()
		return
	}

	// get the command
	command, found := p.take(until(unicode.IsSpace))
	if !found {
		panic("not sure what this means") // todo figure out what this means
	}
	p.skip(while(unicode.IsSpace))

	commandType, ok := commandToType[command]
	if !ok {
		panic(fmt.Errorf("unsupported command: %s", command))
	}
	p.commandType = commandType

	switch commandType {
	case ArithmeticCommand:
		p.arg1 = command
	case PushCommand, PopCommand:
		// get the segment
		segment, found := p.take(until(unicode.IsSpace))
		if !found {
			panic("not sure what this means") // todo figure out what this means
		}
		p.skip(while(unicode.IsSpace))
		p.arg1 = segment

		// get the value
		valuestring, found := p.take(until(unicode.IsSpace))
		if !found {
			panic("not sure what this means") // todo figure out what this means
		}
		p.skip(while(unicode.IsSpace))
		value, err := strconv.Atoi(valuestring)
		if err != nil {
			panic(fmt.Errorf("invalid value: %s", valuestring))
		}

		p.arg2 = value
	}

	// check if this is the end of the file
	if p.pos == len(p.sourceRunes) {
		p.end()
	}
}

func (p *parser) end() {
	p.pos = len(p.sourceRunes)
	p.hasMoreLines = false
}

func (p *parser) peekequals(value string) bool {
	valueRunes := []rune(value)

	if p.pos+len(valueRunes) > len(p.sourceRunes) {
		return false
	}

	for i := 0; i < len(valueRunes); i++ {
		if p.sourceRunes[p.pos+i] != valueRunes[i] {
			return false
		}
	}

	return true
}

func (p *parser) skip(iterator RuneIterator) bool {
	found, pos := iterator(p)
	p.pos = pos
	return found
}

func (p *parser) take(iterator RuneIterator) (string, bool) {
	found, pos := iterator(p)
	output := p.sourceRunes[p.pos:pos]
	p.pos = pos
	return string(output), found
}

type RuneCondition func(rune) bool
type RuneIterator func(p *parser) (bool, int)

func until(condition RuneCondition) RuneIterator {
	return func(p *parser) (bool, int) {
		var found = false
		var i = p.pos
		for ; i < len(p.sourceRunes); i++ {
			if condition(p.sourceRunes[i]) {
				found = true
				break
			}
		}

		return found, i
	}
}

func while(condition RuneCondition) RuneIterator {
	return func(p *parser) (bool, int) {
		var found = false
		var i = p.pos
		for ; i < len(p.sourceRunes); i++ {
			if !condition(p.sourceRunes[i]) {
				found = true
				break
			}
		}

		return found, i
	}
}

func IsNewline(r rune) bool {
	return r == '\n'
}
