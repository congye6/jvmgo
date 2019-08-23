package constant_info

import (
	"jvmgo/classfile/reader"
)

// 字段或方法的描述符
// 方法名称相同，可以根据描述符来区别不同方法
type ConstantNameAndTypeInfo struct {
	nameIndex       uint16 // 字段名称或方法名称
	descriptorIndex uint16 // 描述符
}

func (this *ConstantNameAndTypeInfo) readInfo(reader *reader.ClassReader) {
	this.nameIndex = reader.ReadUint16()
	this.descriptorIndex = reader.ReadUint16()
}
