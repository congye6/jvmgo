package base

import (
	"jvmgo/interpreter"
	"jvmgo/stack"
)

type Instruction interface {
	FetchOperands(reader *interpreter.BytecodeReader)
	Execute(frame *stack.Frame)
}

// 访问常量池，索引由两字节（16位）给出
type Index16 struct {
	Index uint16
}

func (this *Index16) FetchOperands(reader *interpreter.BytecodeReader) {
	this.Index = uint16(reader.ReadUint16())
}

func (this *Index16) Execute(frame *stack.Frame) {

}

// 存储和加载指令需要根据索引存取局部变量表，索引由单字节（8位）给出
type Index8 struct {
	Index uint
}

func (this *Index8) FetchOperands(reader *interpreter.BytecodeReader) {
	this.Index = uint(reader.ReadUint8())
}

func (this *Index8) Execute(frame *stack.Frame) {

}

// 无操作数指令
type NoOperand struct {
}

func (this *NoOperand) FetchOperands(reader *interpreter.BytecodeReader) {
	// nothing to do
}

func (this *NoOperand) Execute(frame *stack.Frame) {

}

// 跳转指令
type Branch struct {
	Offset uint
}

func (this *Branch) FetchOperands(reader *interpreter.BytecodeReader) {
	this.Offset = uint(reader.ReadUint16())
}

func (this *Branch) Execute(frame *stack.Frame) {

}
