package control

import (
	"jvmgo/instruction/base"
	"jvmgo/interpreter"
	"jvmgo/stack"
)

// Branch always
type GOTO struct{ base.Branch }

func (self *GOTO) Execute(frame *stack.Frame) {
	base.BranchGoTo(frame, self.Offset)
}

// Branch always (wide index)
type GOTO_W struct {
	offset uint
}

func (self *GOTO_W) FetchOperands(reader *interpreter.BytecodeReader) {
	self.offset = uint(reader.ReadInt32())
}
func (self *GOTO_W) Execute(frame *stack.Frame) {
	base.BranchGoTo(frame, self.offset)
}
