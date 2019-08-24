package attribute_info

import "jvmgo/classfile/reader"

// 只出现在filed_info中，表示常量表达式的值
type ConstValueAttribute struct {
	valueIndex uint16
}

func (this *ConstValueAttribute) readInfo(reader *reader.ClassReader) {
	this.valueIndex = reader.ReadUint16()
}

func (this *ConstValueAttribute) GetName() string {
	return "ConstantValue"
}

// 不知道具体类型，只返回索引
func (this *ConstValueAttribute) GetValueIndex() uint16 {
	return this.valueIndex
}
