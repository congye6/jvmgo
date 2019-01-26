package main

import "fmt"

const DEFAUL_CLASS_PATH  = ""

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
	if(cmd.classpath==""){
		cmd.classpath=DEFAUL_CLASS_PATH
	}
	startJVM(cmd)
}

func startJVM(cmd *Cmd)  {
	fmt.Println("start jvm")
	fmt.Println(cmd.mainClass+" "+cmd.classpath)
}
