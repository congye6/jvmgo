package method_area

import (
	"fmt"
	"jvmgo/classfile"
	"jvmgo/classpath"
	"strings"
)

const OBJECT_CLASS_NAME = "java/lang/Object"

// 类加载器
type ClassLoader struct {
	basePath     classpath.Entry // 类加载器的路径
	classMap     map[string]*Class
	parentLoader *ClassLoader //父加载器
	name         string       //名称
}

// 创建类加载器，不做预加载
func NewClassLoader(basePath classpath.Entry, parentLoader *ClassLoader, name string) *ClassLoader {
	return &ClassLoader{
		basePath:     basePath,
		classMap:     make(map[string]*Class),
		parentLoader: parentLoader,
		name:         name,
	}
}

// 加载类，父加载器->classMap->加载类
func (this *ClassLoader) LoadClass(name string) *Class {
	if this.parentLoader != nil { //先交给父加载器加载
		class := this.parentLoader.LoadClass(name)
		if class != nil {
			return class
		}
	}
	class := this.classMap[name] //从缓存中找
	if class != nil {
		return class
	}

	return this.loadClass(name)
}

// 加载类，暂不考虑数组
func (this *ClassLoader) loadClass(name string) *Class {
	fmt.Printf("[DEBUG] loading class,loader:%s,class name:%s \n", this.name, name)
	data, _ := this.readClass(name)
	if data == nil { //找不到类
		return nil
	}
	class := this.defineClass(data)
	fmt.Printf("[DEBUG] load class success,loader:%s,class name:%s \n", this.name, name)
	prepare(class)
	return class
}

func (this *ClassLoader) readClass(name string) ([]byte, classpath.Entry) {
	if !strings.Contains(name,".class"){
		name=name+".class"
	}
	data, entry, err := this.basePath.ReadClass(name)
	if err != nil {
		fmt.Printf("[ERROR] java.lang.ClassNotFoundException:%s", name)
		return nil, nil
	}
	return data, entry
}

func (this *ClassLoader) defineClass(data []byte) *Class {
	class := this.parseClass(data)
	class.loader = this
	this.LoadSuperClass(class)
	this.LoadInterfaces(class)
	this.classMap[class.name] = class
	return class
}

func (this *ClassLoader) parseClass(data []byte) *Class {
	classfileVO := classfile.NewClassFile(data)
	classfileVO.Init()
	return newClass(classfileVO)
}

func (this *ClassLoader) LoadSuperClass(class *Class) {
	if class.name == OBJECT_CLASS_NAME { //最终基类不用再加载父类
		return
	}
	class.superClass = class.loader.LoadClass(class.superClassName)
}

func (this *ClassLoader) LoadInterfaces(class *Class) {
	interfaceCount := len(class.interfaceNames)
	if interfaceCount == 0 {
		return
	}
	class.interfaces = make([]*Class, interfaceCount)
	for i, interfaceName := range class.interfaceNames {
		class.interfaces[i] = class.loader.LoadClass(interfaceName)
	}
}