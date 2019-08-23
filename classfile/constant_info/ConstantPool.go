package constant_info

import (
	"jvmgo/classfile/reader"
)

type ConstantPool struct {
	reader    *reader.ClassReader
	constants []ConstantInfo
}

func NewConstantPool(reader *reader.ClassReader) *ConstantPool {
	return &ConstantPool{
		reader: reader,
	}
}

func (this *ConstantPool) Init() {
	constantLength := int(this.reader.ReadUint16()) // 读取常量的数量
	this.constants = make([]ConstantInfo, constantLength)

	constantFactory := ConstantInfoFactory{}
	for i := 1; i < constantLength; i++ { //从1开始，0为无效值
		tag := this.reader.ReadUint8()
		constantInfo := constantFactory.newConstantInfo(tag) // 创建info对象
		constantInfo.readInfo(this.reader)                   // 解析具体数据

		this.constants[i] = constantInfo // 添加到pool中

		//long和double占两个索引
		switch constantInfo.(type) {
		case *ConstantLongInfo, *ConstantDoubleInfo:
			i++
		}
	}
}

func (this *ConstantPool) getConstantInfo(index uint16) ConstantInfo {
	if constantInfo := this.constants[index]; constantInfo != nil {
		return constantInfo
	}
	panic("invalid constant pool index")
}

// 读取字符串字面量
func (this *ConstantPool) GetUtf8(index uint16) string {
	utf8Info := this.getConstantInfo(index).(*ConstantUtf8Info)
	return utf8Info.str
}

// 获取名称和描述符
func (this *ConstantPool) GetNameAndType(index uint16) (string, string) {
	nameTypeInfo := this.constants[index].(*ConstantNameAndTypeInfo)
	return this.GetUtf8(nameTypeInfo.nameIndex), this.GetUtf8(nameTypeInfo.descriptorIndex)
}

// 获取类描述符
func (this *ConstantPool) GetClassName(index uint16) string {
	classInfo := this.constants[index].(*ConstantClassInfo)
	return this.GetUtf8(classInfo.nameIndex)
}
