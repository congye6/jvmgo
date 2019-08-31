package method_area

import "jvmgo/classfile"

type Field struct {
	ClassMember
	constValueIndex uint16
	slotId          uint
}

func newFields(class *Class, fieldInfos []*classfile.FieldInfo) []*Field {
	count := len(fieldInfos)
	fields := make([]*Field, count)
	for i, fieldInfo := range fieldInfos {
		fields[i] = &Field{}
		fields[i].class = class
		fields[i].initClassMember(&fieldInfo.MemberInfo)
		constAttr := fieldInfo.GetConstAttribute()
		if constAttr != nil {
			fields[i].constValueIndex = constAttr.GetValueIndex()
		}
	}
	return fields
}

func (this *Field) isStatic() bool {
	return isStaticFlag(this.accessFlag)
}

func (this *Field) isFinal() bool {
	return isFinalFlag(this.accessFlag)
}

func (this *Field) isLongOrDouble() bool {
	return this.descriptor == "J" || this.descriptor == "D"
}
