package main

import "fmt"

func main()  {
	cmd:=readCmd()
	if cmd.helpFlag{
		printUsage()
	}else{
		fmt.Println("start jvm")
	}
}
