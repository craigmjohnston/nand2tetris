package main

import (
	"os"
	"path/filepath"
	"strings"
)

const Registers int = 16

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

	dat, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	source := string(dat)

	assembler := Assembler{}
	output := assembler.assemble(source)

	// output
	outputFilename := strings.TrimSuffix(filename, fileext) + ".hack"

	if err := os.WriteFile(outputFilename, []byte(output), os.ModePerm); err != nil {
		panic(err)
	}
}
