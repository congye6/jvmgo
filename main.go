package main

import (
	"fmt"
	"jvmgo/classfile"
	"jvmgo/classpath"
	"os"
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
		cmd.jrePath = os.Getenv(JAVA_HOME) + "\\jre"
	}
	startJVM(cmd)
}

func startJVM(cmd *Cmd) {
	fmt.Println("【DEBUG】starting jvm...")
	fmt.Printf("【DEBUG】main class:%s \n, classpath:%s \n", cmd.mainClass, cmd.classpath)
	classpathService := classpath.Parse(cmd.jrePath, cmd.classpath) // classpath就是应用类加载器的路径
	data, _ := classpathService.ReadClass(cmd.mainClass)

	classFile := classfile.NewClassFile(data)
	classFile.Init()
	debug(classFile)
}

func debug(classFile *classfile.ClassFile) {
	fmt.Printf("【DEBUG】class name:%s\n", classFile.GetClassName())
	fmt.Printf("【DEBUG】super class name:%s\n", classFile.GetSuperClassName())
	for _, attribute := range classFile.GetAttributes() {
		fmt.Printf("【DEBUG】attribute name:%s \n", attribute.GetName())
	}
	for _, field := range classFile.GetFields() {
		fmt.Printf("【DEBUG】field name:%s , descriptor:%s \n", field.MemberInfo.GetName(), field.GetDescriptor())
		for _, attribute := range field.GetAttributes() {
			fmt.Printf("【DEBUG】attribute name:%s \n", attribute.GetName())
		}
	}
	for _, method := range classFile.GetMethods() {
		fmt.Printf("【DEBUG】method name:%s , descriptor:%s \n", method.MemberInfo.GetName(), method.GetDescriptor())
		for _, attribute := range method.GetAttributes() {
			fmt.Printf("【DEBUG】attribute name:%s \n", attribute.GetName())
		}
	}
}

//
