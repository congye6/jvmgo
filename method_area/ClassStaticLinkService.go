package method_area

// 静态链接，将常量池中的符号引用转成直接引用，如field.name -> *Field
// 这里静态链接在加载类时做了，也可以在运行时做
// 类：根据类名加载类，将类名替换成类的指针
// 字段：先解析定义字段的类，再在该类中根据字段名查找字段，找不到就根据继承关系往上找
// 方法：和字段一样，不过只有static，final，private方法可以在加载类时静态链接，其他方法需要根据动态类型链接
func link(class *Class) {
	constantPool := class.constantPool
	//fmt.Printf("[DEBUG] linking ....... len:%d name:%s\n", len(constantPool.constants), class.name)
	for _, constant := range constantPool.constants {
		switch constant.(type) {
		case *ConstantClassInfo:
			class := constant.(*ConstantClassInfo).link(class)
			if class != nil {
				//fmt.Printf("[DEBUG] link class,name:%s\n", constant.(*ConstantClassInfo).name)
			}
		case *ConstantFieldDefInfo:
			constant.(*ConstantFieldDefInfo).link(class)
		case *ConstantMethodDefInfo:
			method := constant.(*ConstantMethodDefInfo).staticLink(class)
			if method != nil {
				//fmt.Printf("[DEBUG] link method,name:%s, class name:%s \n", method.name, method.class.name)
			}
		}
	}
}
