package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Coder interface {
	WriteArithmetic(string)
	WritePushPop(CommandType, string, int)
	String() string
}

type coder struct {
	builder strings.Builder
}

func NewCoder() Coder {
	return &coder{
		builder: strings.Builder{},
	}
}

var segmentToName = map[string]string{
	"local":    "LCL",
	"argument": "ARG",
	"this":     "THIS",
	"that":     "THAT",
	"temp":     "TEMP",
}

const RamTempOffset int = 5

func (c *coder) WriteArithmetic(command string) {
	c.comment(command)
	binary := false
	switch command {
	case "add", "sub", "eq", "gt", "lt", "and", "or":
		binary = true
	}

	c.pop("D")
	if binary {
		// this leaves us pointing at the second operand without having to store it somewhere
		c.decrement()
	}

	switch command {
	case "add":
		c.assign("M", "D+M")
		c.assign("D", "M")
		c.push("D")
	case "sub":
		c.assign("M", "D-M")
		c.assign("D", "M")
		c.push("D")
	}

	c.builder.WriteRune('\n')
}

func (c *coder) WritePushPop(commandType CommandType, segment string, value int) {
	if commandType == PushCommand {
		c.comment(fmt.Sprintf("push %s %d", segment, value))

		var address string

		if segment == "constant" {
			address = strconv.Itoa(value)
		} else {
			address = getAddress(segment, value)
		}

		c.address(address)
		if segment == "constant" {
			c.assign("D", "A")
		} else {
			c.assign("D", "M")
		}
		c.pop("D")
	} else {
		c.comment(fmt.Sprintf("pop %s %d", segment, value))

		c.push("D")
		address := getAddress(segment, value)
		c.address(address)
		c.assign("M", "D")
	}

	c.builder.WriteRune('\n')
}

func getAddress(segment string, value int) string {
	if segment == "static" {
		return "Static." + strconv.Itoa(value)
	}
	if segment == "temp" {
		return "R" + strconv.Itoa(value+RamTempOffset)
	}
	if segment == "pointer" {
		if value == 0 {
			segment = "this"
		} else {
			segment = "that"
		}
	}

	address, ok := segmentToName[segment]
	if !ok {
		panic(fmt.Errorf("unsupported segment: %s", segment))
	}
	return address
}

func (c *coder) comment(value string) {
	c.builder.WriteString("// ")
	c.builder.WriteString(value)
	c.builder.WriteRune('\n')
}

func (c *coder) address(value string) {
	c.builder.WriteRune('@')
	c.builder.WriteString(value)
	c.builder.WriteRune('\n')
}

func (c *coder) assign(target string, value string) {
	c.builder.WriteString(target)
	c.builder.WriteRune('=')
	c.builder.WriteString(value)
	c.builder.WriteRune('\n')
}

func (c *coder) push(from string) {
	c.pointToStackEnd()
	c.assign("M", from)
	c.address("SP")
	c.assign("M", "M+1")
	c.assign("A", "M") // point to the end again
}

func (c *coder) decrement() {
	c.address("SP")
	c.assign("M", "M-1")
	c.assign("A", "M") // point to the end
}

func (c *coder) pop(to string) {
	c.pointToStackEnd()
	c.decrement()
}

func (c *coder) pointToStackEnd() {
	c.address("SP")
	c.assign("A", "M")
}

func (c *coder) String() string {
	return c.builder.String()
}
