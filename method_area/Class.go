package method_area

import (
	"jvmgo/classfile"
	"jvmgo/slot"
)

// 类信息
type Class struct {
	accessFlags        uint16
	name               string //字节码是存的索引,方法区存什么没有具体定义
	superClassName     string
	interfaceNames     []string
	constantPool       *ConstantPool
	fields             []*Field
	methods            []*Method
	loader             *ClassLoader
	superClass         *Class
	interfaces         []*Class
	instancesSlotCount uint
	staticSlotCount    uint
	staticVars         slot.Slots //类静态变量
}

func newClass(classfileVO *classfile.ClassFile) *Class {
	class := &Class{
		accessFlags:        classfileVO.GetAccessFlags(),
		name:               classfileVO.GetClassName(),
		superClassName:     classfileVO.GetSuperClassName(),
		interfaceNames:     classfileVO.GetInterfacesName(),
		constantPool:       NewConstantPool(classfileVO.GetConstantPool()),
		superClass:         nil,
		interfaces:         nil,
		instancesSlotCount: 0,
		staticSlotCount:    0,
	}
	class.constantPool.Init()
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

func (this *Class) searchField(name string) *Field {
	for _, field := range this.fields { //在本类中找
		if field.name == name {
			return field
		}
	}

	for _, iface := range this.interfaces { //接口中定义的字段
		if iface == nil {
			continue
		}
		if field := iface.searchField(name); field != nil {
			return field
		}
	}

	if this.superClass != nil { //根据继承关系，自下而上找
		return this.superClass.searchField(name)
	}
	return nil
}

func (this *Class) searchMethod(name string, descriptor string) *Method {
	for _, method := range this.methods { //在本类中找
		if method.name == name && method.descriptor == method.descriptor {
			return method
		}
	}

	if this.superClass != nil { //根据继承关系，自下而上找
		return this.superClass.searchMethod(name, descriptor)
	}

	for _, iface := range this.interfaces { //接口中定义的字段
		if method := iface.searchMethod(name, descriptor); method != nil {
			return method
		}
	}
	return nil
}
