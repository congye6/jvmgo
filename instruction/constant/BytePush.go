package constant

import (
	"jvmgo/interpreter"
	"jvmgo/stack"
)

//字节流中读取一个或两个字节，扩展成int，压入操作数栈

// 读取一个byte
type BIPUSH struct {
	val int8
}

func (this *BIPUSH) FetchOperands(reader *interpreter.BytecodeReader) {
	this.val = int8(reader.ReadUint8())
}

func (this *BIPUSH) Execute(frame *stack.Frame) {
	frame.GetOperandStack().PushInt(int32(this.val))
}

// 读取一个short，两字节
type SIPUSH struct {
	val int16
}

func (this *SIPUSH) FetchOperands(reader *interpreter.BytecodeReader) {
	this.val = reader.ReadInt16()
}

func (this *SIPUSH) Execute(frame *stack.Frame) {
	frame.GetOperandStack().PushInt(int32(this.val))
}
