package classfile

import (
	"fmt"
	"jvmgo/classfile/attribute_info"
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
	fields       []*FieldInfo  //字段
	methods      []*MethodInfo //方法
	attributes   []attribute_info.AttributeInfo
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
	this.fields = readFields(this.classReader, this.constantPool)
	this.methods = readMethods(this.classReader, this.constantPool)
	this.attributes = attribute_info.ReadAttributes(this.classReader, this.constantPool)
}

func (this *ClassFile) GetAccessFlags() uint16 {
	return this.accessFlags
}

func (this *ClassFile) GetClassName() string {
	return this.constantPool.GetClassName(this.thisClass)
}

func (this *ClassFile) GetSuperClassName() string {
	if this.GetClassName() == "java/lang/Object"{
		return ""
	}
	return this.constantPool.GetClassName(this.superClass)
}

func (this *ClassFile) GetInterfacesName() []string {
	count := len(this.interfaces)
	interfaceNames := make([]string, count)
	for i, interfaceName := range this.interfaces {
		interfaceNames[i] = this.constantPool.GetUtf8(interfaceName)
	}
	return interfaceNames
}

func (this *ClassFile) GetConstantPool() *constant_info.ConstantPool{
	return this.constantPool
}

func (this *ClassFile) GetAttributes() []attribute_info.AttributeInfo {
	return this.attributes
}

func (this *ClassFile) GetFields() []*FieldInfo {
	return this.fields
}

func (this *ClassFile) GetMethods() []*MethodInfo {
	return this.methods
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
}
