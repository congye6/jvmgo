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

func (this *ClassMember) GetClass() *Class {
	return this.class
}
