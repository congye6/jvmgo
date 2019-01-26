package main

import (
	"flag"
	"fmt"
	"os"
)

type Cmd struct {
	helpFlag bool
	versionFlag bool
	classpath string
	mainClass string
	args []string
}

//从命令行中解析参数
func readCmd() *Cmd {
	cmd:=new(Cmd)
	flag.BoolVar(&cmd.helpFlag,"help",false,"print help message")
	return cmd
}

func printUsage(){
	fmt.Println("Usage: %s [-options] class [args...]",os.Args[0])
}
