package control

import (
	"jvmgo/instruction/base"
	"jvmgo/interpreter"
	"jvmgo/stack"
)

// Jump subroutine
// Java6之前用于finally子句的调用
type JSR struct{ base.Branch }

func (self *JSR) Execute(frame *stack.Frame) {
	panic("todo")
}

// Jump subroutine (wide index)
type JSR_W struct {
	offset int
}

func (self *JSR_W) FetchOperands(reader *interpreter.BytecodeReader) {
	self.offset = int(reader.ReadInt32())
}
func (self *JSR_W) Execute(frame *stack.Frame) {
	panic("todo")
}

// Return from subroutine
type RET struct{ base.Index8 }

func (self *RET) Execute(frame *stack.Frame) {
	panic("todo")
}
