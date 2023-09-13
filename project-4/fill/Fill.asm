// This file is part of www.nand2tetris.org
// and the book "The Elements of Computing Systems"
// by Nisan and Schocken, MIT Press.
// File name: projects/04/Fill.asm

// Runs an infinite loop that listens to the keyboard input. 
// When a key is pressed (any key), the program blackens the screen,
// i.e. writes "black" in every pixel. When no key is pressed, the
// program clears the screen, i.e. writes "white" in every pixel.

// Put your code here.
@16384
D=A
@SCREENMAX
M=D
@prevColor
M=0
@color
M=0
@i
M=0
@target
M=0

(LOOP)
  @KBD
  D=M
  @NOKEYS
  D;JEQ

  // draw black
  @color
  M=-1

  @CONTINUE
  0;JMP

  (NOKEYS)
  // draw white
  @color
  M=0  

  (CONTINUE)
  @color
  D=M
  @prevColor
  D=D-M
  @DRAW
  D;JEQ // jump to DRAW if we're using the same color

  // reset loop if we're not
  @i
  M=0
  @color
  D=M
  @prevColor
  M=D

  // draw
  (DRAW)
  @i
  D=M
  @SCREEN
  D=A+D
  @target
  M=D
  @color
  D=M
  @target
  A=M
  M=D

  // check overflow
  @i
  D=M
  @SCREENMAX
  D=D-M
  @RESET
  D;JEQ

  // increment & loop
  @i
  M=M+1
  @LOOP
  0;JMP

  (RESET)
  @i
  M=0

  @LOOP
  0;JMP