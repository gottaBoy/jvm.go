package main

import "flag"
import "fmt"
import "os"

//定义Cmd结构体
type Cmd struct{
    helpFlag     bool
    versionFlag     bool
    cpOption     string
    class     string
    args     []string
}

//解析命令行参数
func parseCmd() *Cmd {
    cmd:=&Cmd{}

    //将printUsage函数传给flag.Usage
    flag.Usage=printUsage
    //设置各种解析的选项
    flag.BoolVar(&cmd.helpFlag, "help", false, "print help message")
    flag.BoolVar(&cmd.helpFlag, "?", false, "print help message")
    flag.BoolVar(&cmd.versionFlag, "version", false, "print version and exit")
    flag.StringVar(&cmd.cpOption, "classpath", "", "classpath")
    flag.StringVar(&cmd.cpOption, "cp", "", "classpath")
    //所有选项设置完成后调用flag.Parse解析所有选项，如果Parse失败，则调用flag.Usage打印帮助信息
    flag.Parse()

    //调用flag.Args函数捕获未被解析的参数，第一个参数为主类名，后面的为传递给主类的参数
    args:=flag.Args()
    if len(args)>0{
        cmd.class=args[0]
        cmd.args=args[1:]
    }

    return cmd
}

func printUsage() {
    fmt.Printf("Usage:%s[-options] class [args...]\n",os.Args[0])
}