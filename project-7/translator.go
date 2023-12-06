package main

import "fmt"

type Translator struct {
}

func (Translator) Translate(source string) (string, error) {
	p := NewParser(source)

	for p.HasMoreLines() {
		p.Advance()
		// todo output command
		fmt.Printf("Type: %v, Arg1: %v, Arg2: %v\n", p.CommandType(), p.Arg1(), p.Arg2())
	}

	// todo return output string
	//panic("not implemented")
	return "", nil
}
