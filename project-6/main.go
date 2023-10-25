package main

import (
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
	code := Code{}

	output := ""

	for parser.HasMoreLines {
		parser.Advance()
		if parser.InstructionType == CInstruction {
			//fmt.Printf("%d DEST: %s, COMP: %s, JUMP: %s", parser.InstructionType, parser.Dest, parser.Comp, parser.Jump)
			//fmt.Println()
			encoded := code.CInstruction(parser.Dest, parser.Comp, parser.Jump)
			output += encoded
			//fmt.Println(encoded)
		} else {
			//fmt.Println(parser.InstructionType, parser.Symbol)
			encoded := "0" + code.ToBinary(parser.Symbol)
			output += encoded
			//fmt.Println(encoded)
		}

		output += "\n"
	}

	//fmt.Println(output)

	// output
	outputFilename := strings.TrimSuffix(filename, fileext) + ".hack"
	//fmt.Println(outputFilename)

	if err := os.WriteFile(outputFilename, []byte(output), os.ModePerm); err != nil {
		panic(err)
	}
}
