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
	jrePath string
	args []string
}

//从命令行中解析参数
func readCmd() *Cmd {
	cmd:=new(Cmd)
	flag.BoolVar(&cmd.helpFlag,"help",false,"print help message")
	flag.BoolVar(&cmd.versionFlag,"version",  false,"print version message")
	flag.StringVar(&cmd.classpath,"cp","","classpath")
	flag.StringVar(&cmd.jrePath,"jre","","jre path")
	flag.Parse()
	args:=flag.Args()
	if len(args)>0{
		cmd.mainClass=args[0]
		cmd.args=args[1:]
	}
	return cmd
}

func printUsage(){
	fmt.Printf("Usage: %s [-options] class [args...]",os.Args[0])
}
