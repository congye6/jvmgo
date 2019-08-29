package method_area

import "jvmgo/classfile"

type ClassMember struct {
	accessFlag uint16
	name       string
	descriptor string
	class      *Class
}

func (this *ClassMember) initClassMember(memberInfo *classfile.MemberInfo) {
	this.name = memberInfo.GetName()
	this.accessFlag = memberInfo.GetAccessFlag()
	this.descriptor = memberInfo.GetDescriptor()
}

type Field struct {
	ClassMember
}

func newFields(class *Class, fieldInfos []*classfile.FieldInfo) []*Field {
	count := len(fieldInfos)
	fields := make([]*Field, count)
	for i, fieldInfo := range fieldInfos {
		fields[i] = &Field{}
		fields[i].class = class
		fields[i].initClassMember(&fieldInfo.MemberInfo)
	}
	return fields
}

type Method struct {
	ClassMember
	maxStack  uint16
	maxLocals uint16
	code      []byte
}

func newMethods(class *Class, methodInfos []*classfile.MethodInfo) []*Method {
	count := len(methodInfos)
	methods := make([]*Method, count)
	for i, methodInfo := range methodInfos {
		methods[i] = &Method{}
		methods[i].class = class
		methods[i].initClassMember(&methodInfo.MemberInfo)

		codeAttribute := methodInfo.GetCodeAttribute()
		if codeAttribute == nil {
			continue
		}
		methods[i].maxStack = codeAttribute.GetMaxStack()
		methods[i].maxLocals = codeAttribute.GetMaxLocals()
		methods[i].code = codeAttribute.GetCode()
	}
	return methods
}
