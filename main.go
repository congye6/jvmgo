package main

import (
	"fmt"
	"jvmgo/method_area"
	"os"
	"time"
)

const JAVA_HOME = "JAVA_HOME"

const CLASS_PATH = "CLASS_PATH"

func main() {
	cmd := readCmd()
	if cmd.helpFlag {
		printUsage()
		return
	}

	if cmd.versionFlag {
		fmt.Println("jvmgo version 1.0")
		return
	}

	if cmd.mainClass == "" {
		fmt.Println("Please input main class")
		return
	}
	if cmd.classpath == "" {
		cmd.classpath = os.Getenv(CLASS_PATH)
	}
	if cmd.jrePath == "" {
		cmd.jrePath = os.Getenv(JAVA_HOME)
	}
	startJVM(cmd)
}

func startJVM(cmd *Cmd) {
	fmt.Println("【DEBUG】starting jvm...")
	fmt.Printf("【DEBUG】main class:%s \n, classpath:%s \n", cmd.mainClass, cmd.classpath)
	before := time.Now().Unix()
	applicationLoader := method_area.GetApplicationLoader(cmd.jrePath, cmd.classpath) // 应用类加载器
	class := applicationLoader.LoadClass(cmd.mainClass)
	after := time.Now().Unix()
	fmt.Printf("[DEBUG] loads class time:%d \n", after-before)
	debug(class)
}

func debug(class *method_area.Class) {
	fmt.Printf("【DEBUG】class name:%s\n", class.GetName())
	fmt.Printf("【DEBUG】super class name:%s\n", class.GetSuperClassName())
	//for _, attribute := range classFile.GetAttributes() {
	//	fmt.Printf("【DEBUG】attribute name:%s \n", attribute.GetName())
	//}
	//for _, field := range classFile.GetFields() {
	//	fmt.Printf("【DEBUG】field name:%s , descriptor:%s \n", field.MemberInfo.GetName(), field.GetDescriptor())
	//	for _, attribute := range field.GetAttributes() {
	//		fmt.Printf("【DEBUG】attribute name:%s \n", attribute.GetName())
	//	}
	//}
	//for _, method := range classFile.GetMethods() {
	//	fmt.Printf("【DEBUG】method name:%s , descriptor:%s \n", method.MemberInfo.GetName(), method.GetDescriptor())
	//	for _, attribute := range method.GetAttributes() {
	//		fmt.Printf("【DEBUG】attribute name:%s \n", attribute.GetName())
	//	}
	//}
}

//
