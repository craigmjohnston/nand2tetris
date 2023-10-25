package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// input
	args := os.Args
	if len(args) < 2 {
		panic("Missing args.")
	}

	filename := args[1]
	if _, err := os.Stat(filename); err != nil && os.IsNotExist(err) {
		panic("File does not exist")
	}

	fileext := filepath.Ext(filename)
	if fileext != ".asm" {
		panic("Wrong file extension (should be .asm)")
	}

	source, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	parser := NewParser(string(source))

	for parser.HasMoreLines {
		parser.Advance()
		if parser.InstructionType == CInstruction {
			fmt.Printf("%d DEST: %s, COMP: %s, JUMP: %s", parser.InstructionType, parser.Dest, parser.Comp, parser.Jump)
			fmt.Println()
		} else {
			fmt.Println(parser.InstructionType, parser.Symbol)
		}
	}

	// output
	outputFilename := strings.TrimSuffix(filename, fileext) + ".hack"
	fmt.Println(outputFilename)
}
