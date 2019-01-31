package main

import (
	"fmt"
	"jvmgo/classpath"
	"os"
)

const JAVA_HOME  = "JAVA_HOME"

const CLASS_PATH  = "CLASS_PATH"

func main()  {
	cmd:=readCmd()
	if cmd.helpFlag{
		printUsage()
		return
	}

	if cmd.versionFlag{
		fmt.Println("jvmgo version 1.0")
		return
	}

	if cmd.mainClass==""{
		fmt.Println("Please input main class")
		return
	}
	if cmd.classpath==""{
		cmd.classpath=os.Getenv(CLASS_PATH)
	}
	if cmd.jrePath==""{
		cmd.jrePath=os.Getenv(JAVA_HOME)+"\\jre"
	}
	startJVM(cmd)
}

func startJVM(cmd *Cmd)  {
	fmt.Println("start jvm")
	fmt.Println(cmd.mainClass+" "+cmd.classpath)
	classpath:=classpath.Parse(cmd.jrePath,cmd.classpath)
	data,_:=classpath.ReadClass(cmd.mainClass)
	fmt.Println(string(data))
}

//jvmgo.exe -cp E:\workspace\practice\bin E:\workspace\practice\bin\Solution.class
