package constant_info

import "jvmgo/classfile"

// 类描述符和名称+描述符可以唯一确定一个字段或方法
type ConstantMemberDefInfo struct {
	classIndex       uint16 // 在哪个类中定义
	nameAndTypeIndex uint16
}

func (this *ConstantMemberDefInfo) readInfo(reader *classfile.ClassReader) {
	this.classIndex = reader.ReadUint16()
	this.nameAndTypeIndex = reader.ReadUint16()
}

// 字段
type ConstantFieldDefInfo struct {
	ConstantMemberDefInfo
}

// 方法
type ConstantMethodDefInfo struct {
	ConstantMemberDefInfo
}

// 接口方法
type ConstantInterfaceMethodDefInfo struct {
	ConstantMemberDefInfo
}

