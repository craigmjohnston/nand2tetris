package main

import "fmt"

type Translator struct {
}

// todo pass errors through from parser/coder so we can return them from here

func (Translator) Translate(source string) (string, error) {
	p := NewParser(source)
	c := NewCoder()

	for p.HasMoreLines() {
		p.Advance()

		switch p.CommandType() {
		case ArithmeticCommand:
			c.WriteArithmetic(p.Arg1())
		case PushCommand, PopCommand:
			c.WritePushPop(p.CommandType(), p.Arg1(), p.Arg2())
		}

		fmt.Printf("Type: %v, Arg1: %v, Arg2: %v\n", p.CommandType(), p.Arg1(), p.Arg2())
	}

	output := c.String()
	return output, nil
}
