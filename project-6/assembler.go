package main

import (
	"sort"
	"strconv"
	"strings"
)

// todo this probably doesn't need to be a struct, since it's just the one method

type Assembler struct {
}

func (a Assembler) assemble(source string) string {
	parser := NewParser(source)

	// build symbol table
	symbols := map[string]int{
		"SP":     0,
		"LCL":    1,
		"ARG":    2,
		"THIS":   3,
		"THAT":   4,
		"SCREEN": 16384,
		"KBD":    24576,
	}

	for i := 0; i < Registers; i++ {
		symbols["R"+strconv.Itoa(i)] = i
	}

	// first-pass - symbols
	pc := 0
	variableSymbols := make(map[string]int)
	symbolCounter := 0

	for parser.HasMoreLines {
		parser.Advance()

		if parser.InstructionType == AInstruction {
			if _, err := strconv.Atoi(parser.Symbol); err != nil {
				if _, ok := symbols[parser.Symbol]; !ok {
					if _, ok := variableSymbols[parser.Symbol]; !ok {
						variableSymbols[parser.Symbol] = symbolCounter
						symbolCounter += 1
					}
				}
			}
		} else if parser.InstructionType == LInstruction {
			symbols[parser.Symbol] = pc

			if _, ok := variableSymbols[parser.Symbol]; ok {
				delete(variableSymbols, parser.Symbol)
			}
		}

		if parser.InstructionType != LInstruction {
			pc += 1
		}
	}

	type SymbolOrder struct {
		symbol string
		order  int
	}

	mapped := make([]SymbolOrder, 0)
	for symbol, order := range variableSymbols {
		mapped = append(mapped, SymbolOrder{symbol, order})
	}

	sort.Slice(mapped, func(i int, j int) bool { return mapped[i].order < mapped[j].order })

	addressCounter := Registers
	for _, orderedSymbol := range mapped {
		symbols[orderedSymbol.symbol] = addressCounter
		addressCounter += 1
	}

	// second-pass - output
	parser = NewParser(source)
	code := Code{}
	sb := strings.Builder{}

	for parser.HasMoreLines {
		parser.Advance()
		if parser.InstructionType == CInstruction {
			encoded := code.CInstruction(parser.Dest, parser.Comp, parser.Jump)
			sb.WriteString(encoded)
			sb.WriteString("\n")
		} else if parser.InstructionType == AInstruction {
			var intVal int
			if _, err := strconv.Atoi(parser.Symbol); err != nil {
				if address, ok := symbols[parser.Symbol]; ok {
					intVal = address
				}
			} else {
				numberValue, err := strconv.Atoi(parser.Symbol)
				if err != nil {
					panic(err)
				}
				intVal = numberValue
			}

			encoded := "0" + code.ToBinary(int64(intVal))
			sb.WriteString(encoded)
			sb.WriteString("\n")
		}
	}

	// return output
	return sb.String()
}
