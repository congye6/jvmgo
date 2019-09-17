package method_area

import "jvmgo/classfile"

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

func (this *Method) isStaticLink() bool {
	return isFinalFlag(this.accessFlag) || isStaticFlag(this.accessFlag) || isPrivate(this.accessFlag)
}

func (this *Method) GetMaxLocals() uint16 {
	return this.maxLocals
}

func (this *Method) GetMaxStack() uint16 {
	return this.maxStack
}
