package method_area

import "jvmgo/slot"

type Object struct {
	// todo
	class  *Class
	fields []slot.Slots //实例变量
}
