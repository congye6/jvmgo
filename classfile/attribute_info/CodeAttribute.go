package attribute_info

import (
	"jvmgo/classfile/reader"
	"jvmgo/constant_pool"
)

type CodeAttribute struct {
	constantPool   *constant_pool.ConstantPool
	maxStack       uint16 // 最大操作数栈
	maxLocals      uint16 // 局部变量表大小
	code           []byte
	exceptionTable []*ExceptionTableEntry
	attributes     []AttributeInfo
}

func (this *CodeAttribute) readInfo(reader *reader.ClassReader) {
	this.maxStack = reader.ReadUint16()
	this.maxLocals = reader.ReadUint16()
	codeLength := reader.ReadUint32()
	this.code = reader.ReadBytes(codeLength)
	this.exceptionTable = readExceptionTable(reader)
	this.attributes = ReadAttributes(reader, this.constantPool)
}

func (this *CodeAttribute) GetMaxStack() uint16 {
	return this.maxStack
}

func (this *CodeAttribute) GetMaxLocals() uint16 {
	return this.maxLocals
}

func (this *CodeAttribute) GetCode() []byte {
	return this.code
}

func (this *CodeAttribute) GetName() string {
	return "Code"
}

func readExceptionTable(reader *reader.ClassReader) []*ExceptionTableEntry {
	exceptionTableLength := reader.ReadUint16()
	exceptionTable := make([]*ExceptionTableEntry, exceptionTableLength)

	for i := range exceptionTable {
		exceptionTable[i] = &ExceptionTableEntry{
			startPc:   reader.ReadUint16(),
			endPc:     reader.ReadUint16(),
			handlePc:  reader.ReadUint16(),
			catchType: reader.ReadUint16(),
		}
	}

	return exceptionTable
}

// 异常处理
type ExceptionTableEntry struct {
	startPc   uint16
	endPc     uint16
	handlePc  uint16
	catchType uint16
}
