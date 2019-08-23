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
	fmt.Printf("【DEBUG】class name:%s\n", classFile.GetClassName())
	fmt.Printf("【DEBUG】super class name:%s\n", classFile.GetSuperClassName())
}

//./jvmgo -cp /Users/zhoucong/workspace/practice/out/production/practice /Users/zhoucong/workspace/practice/out/production/practice/Solution.class
