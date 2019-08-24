package attribute_info

import (
	"jvmgo/classfile/constant_info"
	"jvmgo/classfile/reader"
)

// 源文件名称
type SourceFileAttribute struct {
	constantPool    *constant_info.ConstantPool
	sourceFileIndex uint16 //源文件名称索引
}

func (this *SourceFileAttribute) readInfo(reader *reader.ClassReader) {
	this.sourceFileIndex = reader.ReadUint16()
}

func (this *SourceFileAttribute) GetName() string {
	return "SourceFile"
}

func (this *SourceFileAttribute) FileName() string {
	return this.constantPool.GetUtf8(this.sourceFileIndex)
}
