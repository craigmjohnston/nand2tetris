package main

import "os"

// todo take filepath and output path as cmd arg

func main() {
	dat, err := os.ReadFile("res/test.vm")
	if err != nil {
		panic(err)
	}
	source := string(dat)
	t := Translator{}
	output, err := t.Translate(source)
	err = os.WriteFile("output/test.hack", []byte(output), os.ModePerm)
}
