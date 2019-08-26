package attribute_info

import "jvmgo/classfile/reader"

// 未解析的属性
type UnparsedAttribute struct {
	name   string //属性名
	length uint32 //info长度
	info   []byte
}

func (this *UnparsedAttribute) readInfo(reader *reader.ClassReader) {
	this.info = reader.ReadBytes(this.length)
}

func (this *UnparsedAttribute) GetName() string {
	return this.name
}
