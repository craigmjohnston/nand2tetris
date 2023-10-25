package main

import (
	"strconv"
	"strings"
)

type Code struct {
}

func (Code) Dest(input string) string {
	output := ""

	output += booltob(strings.Contains(input, "A"))
	output += booltob(strings.Contains(input, "D"))
	output += booltob(strings.Contains(input, "M"))

	return output
}

var CompMap = map[string]string{
	"0":   "0101010",
	"1":   "0111111",
	"-1":  "0111010",
	"D":   "0001100",
	"A":   "0110000",
	"!D":  "0001101",
	"!A":  "0110001",
	"-D":  "0001111",
	"-A":  "0110011",
	"D+1": "0011111",
	"A+1": "0110111",
	"D-1": "0001110",
	"A-1": "0110010",
	"D+A": "0000010",
	"D-A": "0010011",
	"A-D": "0000111",
	"D&A": "0000000",
	"D|A": "0010101",
	"M":   "1110000",
	"!M":  "1110001",
	"-M":  "1110011",
	"M+1": "1110111",
	"M-1": "1110010",
	"D+M": "1000010",
	"D-M": "1010011",
	"M-D": "1000111",
	"D&M": "1000000",
	"D|M": "1010101",
}

func (Code) Comp(input string) string {
	return CompMap[input]
}

var JumpMap = map[string]string{
	"JGT": "001",
	"JEQ": "010",
	"JGE": "011",
	"JLT": "100",
	"JNE": "101",
	"JLE": "110",
	"JMP": "111",
}

func (Code) Jump(input string) string {
	if input == "" {
		return "000"
	}

	return JumpMap[input]
}

func (c Code) CInstruction(dest string, comp string, jump string) string {
	output := "111"

	output += c.Comp(comp)
	output += c.Dest(dest)
	output += c.Jump(jump)

	return output
}

func booltob(value bool) string {
	if value {
		return "1"
	}

	return "0"
}

func (Code) ToBinary(input int64) string {
	formatted := strconv.FormatInt(input, 2)

	if len(formatted) < 15 {
		formatted = strings.Repeat("0", 15-len(formatted)) + formatted
	}

	return formatted
}
