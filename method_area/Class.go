package method_area

import (
	"jvmgo/classfile"
	"jvmgo/classfile/constant_info"
)

// 类信息
type Class struct {
	accessFlags        uint16
	name               string //字节码是存的索引,方法区存什么没有具体定义
	superClassName     string
	interfaceNames     []string
	constantPool       *constant_info.ConstantPool
	fields             []*Field
	methods            []*Method
	loader             *ClassLoader
	superClass         *Class
	interfaces         []*Class
	instancesSlotCount uint
	staticSlotCount    uint
	//staticVars *Slot
}

func newClass(classfileVO *classfile.ClassFile) *Class {
	class := &Class{
		accessFlags:        classfileVO.GetAccessFlags(),
		name:               classfileVO.GetClassName(),
		superClassName:     classfileVO.GetSuperClassName(),
		interfaceNames:     classfileVO.GetInterfacesName(),
		constantPool:       classfileVO.GetConstantPool(),
		superClass:         nil,
		interfaces:         nil,
		instancesSlotCount: 0,
		staticSlotCount:    0,
	}
	newFields(class, classfileVO.GetFields())
	newMethods(class, classfileVO.GetMethods())
	return class
}

func (this *Class) GetName() string {
	return this.name
}

func (this *Class) GetSuperClassName() string {
	return this.superClassName
}
