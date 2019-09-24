package constant

import (
	"jvmgo/instruction/base"
	"jvmgo/method_area"
	"jvmgo/stack"
)

// TODO

// Push item from run-time constant pool
type LDC struct{ base.Index8 }

func (self *LDC) Execute(frame *stack.Frame) {
	_ldc(frame, uint16(self.Index))
}

// Push item from run-time constant pool (wide index)
type LDC_W struct{ base.Index16 }

func (self *LDC_W) Execute(frame *stack.Frame) {
	_ldc(frame, self.Index)
}

func _ldc(frame *stack.Frame, index uint16) {
	stack := frame.GetOperandStack()
	cp := frame.GetConstantPool()
	c := cp.GetConstantInfo(index)

	switch c.(type) {
	case *method_area.ConstantIntegerInfo:
		stack.PushInt(c.(*method_area.ConstantIntegerInfo).GetVal())
	case float32:
		stack.PushFloat(c.(float32))
	case string:

	case *method_area.ConstantClassInfo:

	default:
		// todo
		// ref to MethodType or MethodHandle
		panic("todo: ldc! %v")
	}
}

// Push long or double from run-time constant pool (wide index)
type LDC2_W struct{ base.Index16 }

func (self *LDC2_W) Execute(frame *stack.Frame) {
	stack := frame.GetOperandStack()
	cp := frame.GetConstantPool()
	c := cp.GetConstantInfo(self.Index)

	switch c.(type) {
	case int64:
		stack.PushLong(c.(int64))
	case float64:
		stack.PushDouble(c.(float64))
	default:
		panic("ldc2_w! %v")
	}
}
