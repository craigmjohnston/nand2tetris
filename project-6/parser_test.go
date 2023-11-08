package main

import (
	"os"
	"testing"
)
import "github.com/stretchr/testify/assert"

// this is a pretty basic comparison test, but it's more so that I
// have an entry point for the profiler
func TestAssembler(t *testing.T) {
	testFiles := []string{
		"res/Add",
		"res/Max",
		"res/MaxL",
		"res/Pong",
		"res/PongL",
		"res/Rect",
		"res/RectL",
	}

	for _, testFile := range testFiles {
		source := readSource(t, testFile+".asm")
		assembler := Assembler{}
		got := assembler.assemble(source)
		want := readSource(t, testFile+".hack")
		assert.Equal(t, want, got)
	}
}

func readSource(t *testing.T, path string) string {
	dat, err := os.ReadFile(path)
	assert.NoError(t, err)
	return string(dat)
}
