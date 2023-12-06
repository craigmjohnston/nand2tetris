// push constant 111
@111
D=A
@SP
A=M
@SP
M=M-1
A=M

// push constant 333
@333
D=A
@SP
A=M
@SP
M=M-1
A=M

// push constant 888
@888
D=A
@SP
A=M
@SP
M=M-1
A=M

// pop static 8
@SP
A=M
M=D
@SP
M=M+1
A=M
@Static.8
M=D

// pop static 3
@SP
A=M
M=D
@SP
M=M+1
A=M
@Static.3
M=D

// pop static 1
@SP
A=M
M=D
@SP
M=M+1
A=M
@Static.1
M=D

// push static 3
@Static.3
D=M
@SP
A=M
@SP
M=M-1
A=M

// push static 1
@Static.1
D=M
@SP
A=M
@SP
M=M-1
A=M

// sub
@SP
A=M
@SP
M=M-1
A=M
@SP
M=M-1
A=M
M=D-M
D=M
@SP
A=M
M=D
@SP
M=M+1
A=M

// push static 8
@Static.8
D=M
@SP
A=M
@SP
M=M-1
A=M

// add
@SP
A=M
@SP
M=M-1
A=M
@SP
M=M-1
A=M
M=D+M
D=M
@SP
A=M
M=D
@SP
M=M+1
A=M

