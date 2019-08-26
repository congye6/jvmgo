package attribute_info

import "jvmgo/classfile/reader"

// 方法抛出的异常
type ExceptionAttribute struct {
	exceptionIndexTable []uint16
}

func (this *ExceptionAttribute) readInfo(reader *reader.ClassReader) {
	this.exceptionIndexTable = reader.ReadUint16s()
}

func (this *ExceptionAttribute) GetName() string {
	return "Exception"
}

func (this *ExceptionAttribute) GetThrowExceptions() []uint16 {
	return this.exceptionIndexTable
}
