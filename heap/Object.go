package heap

import "jvmgo/method_area"

type Object struct {
	// todo
	class  *method_area.Class
	fields []method_area.Slots //实例变量
}
