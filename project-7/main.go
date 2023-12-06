package main

import "os"

// todo take filepath and output path as cmd arg

func main() {
	dat, err := os.ReadFile("res/MemoryAccess/StaticTest/StaticTest.vm")
	if err != nil {
		panic(err)
	}
	source := string(dat)
	t := Translator{}
	output, err := t.Translate(source)
	err = os.WriteFile("output/SimpleAdd.hack", []byte(output), os.ModePerm)
}
