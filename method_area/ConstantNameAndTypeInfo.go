package method_area

import (
	"jvmgo/classfile/reader"
)

// 字段或方法的描述符
// 方法名称相同，可以根据描述符来区别不同方法
type ConstantNameAndTypeInfo struct {
	constantPool *ConstantPool
	name         string // 字段名称或方法名称
	descriptor   string // 描述符
}

func (this *ConstantNameAndTypeInfo) readInfo(reader *reader.ClassReader) {
	this.name = this.constantPool.GetUtf8(reader.ReadUint16())
	this.descriptor = this.constantPool.GetUtf8(reader.ReadUint16())
}
