package constant_pool

import (
	"jvmgo/classfile/reader"
)

// 类描述符和名称+描述符可以唯一确定一个字段或方法
type ConstantMemberDefInfo struct {
	constantPool     *ConstantPool
	classIndex       uint16 // 在哪个类中定义
	nameAndTypeIndex uint16
}

func (this *ConstantMemberDefInfo) readInfo(reader *reader.ClassReader) {
	this.classIndex = reader.ReadUint16()
	this.nameAndTypeIndex = reader.ReadUint16()
}

func (this *ConstantMemberDefInfo) GetClassIndex() uint16 {
	return this.classIndex
}

func (this *ConstantMemberDefInfo) GetNameAndType() (string, string) {
	return this.constantPool.GetNameAndType(this.nameAndTypeIndex)
}

// 接口方法
type ConstantInterfaceMethodDefInfo struct {
	ConstantMemberDefInfo
}
