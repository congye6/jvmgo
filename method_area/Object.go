package method_area

import "jvmgo/slot"

type Object struct {
	// todo
	class  *Class
	fields slot.Slots //实例变量
}

func (this *Object) GetFields() slot.Slots {
	return this.fields
}

func (this *Object) IsInstanceOf(class *Class) bool {
	return this.class.IsAssignableFrom(class)
}
