package classfile

import (
	"fmt"
	"jvmgo/classfile/constant_info"
	"jvmgo/classfile/reader"
)

// class文件属性
type ClassFile struct {
	classReader *reader.ClassReader

	// magic uint32
	minorVersion uint16 //版本号
	majorVersion uint16
	constantPool *constant_info.ConstantPool
	accessFlags  uint16 //访问标志
	thisClass    uint16 //类名的索引
	superClass   uint16 //超类索引
	interfaces   []uint16
	// TODO
}

func NewClassFile(data []byte) *ClassFile {
	return &ClassFile{
		classReader: &reader.ClassReader{Data: data},
	}
}

// 解析类文件
func (this *ClassFile) Init() {
	this.readAndCheckMagic()
	this.readAndCheckVersion()
	this.constantPool = constant_info.NewConstantPool(this.classReader)
	this.constantPool.Init()
	this.accessFlags = this.classReader.ReadUint16()
	this.thisClass = this.classReader.ReadUint16()
	this.superClass = this.classReader.ReadUint16()
	this.interfaces = this.classReader.ReadUint16s()
}

func (this *ClassFile) GetClassName() string {
	return this.constantPool.GetClassName(this.thisClass)
}

func (this *ClassFile) GetSuperClassName() string {
	return this.constantPool.GetClassName(this.superClass)
}

func (this *ClassFile) readAndCheckMagic() {
	if this.classReader == nil {
		fmt.Printf("【WARNING】reading magic,class reader id nil \n")
		return
	}

	magic := this.classReader.ReadUint32()
	if magic != 0xCAFEBABE {
		fmt.Printf("error magic:%d \n", magic)
		panic("java.lang.ClassFormatError,magic!")
	}
}

func (this *ClassFile) readAndCheckVersion() {
	this.minorVersion = this.classReader.ReadUint16()
	this.majorVersion = this.classReader.ReadUint16()
	fmt.Printf("reading version, minor:%d, major:%d \n", this.minorVersion, this.majorVersion)
}
