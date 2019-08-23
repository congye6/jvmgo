package classfile

import "fmt"

// class文件属性
type ClassFile struct {
	classReader *ClassReader

	// magic uint32
	minorVersion uint16
	majorVersion uint16
	// TODO
}

func NewClassFile(data []byte) *ClassFile {
	return &ClassFile{
		classReader: &ClassReader{data:data},
	}
}

// 解析类文件
func (this *ClassFile) Init() {
	this.readAndCheckMagic()
	this.readAndCheckVersion()
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
