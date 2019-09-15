package constant_pool

import (
	"jvmgo/classfile/reader"
)

// 类描述符，包括类，超类，接口都用这个表示
type ConstantClassInfo struct {
	constantPool *ConstantPool
	nameIndex    uint16 // 类名
}

func (this *ConstantClassInfo) readInfo(reader *reader.ClassReader) {
	this.nameIndex = reader.ReadUint16()
}
