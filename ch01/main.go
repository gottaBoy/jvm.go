package main

import "fmt"

func main() {
    //调用parseCmd解析命令行参数
    cmd:=parseCmd()

    if cmd.versionFlag{
        //输入了-version选项
        fmt.Println("version 0.0.1")
    }else if cmd.helpFlag||cmd.class==""{
        //输入了-help选项
        printUsage()
    }else{
        //启动jvm
        stratJVM(cmd)
    }
}

func stratJVM(cmd *Cmd){
    fmt.Printf("classpath:%s class:%s args:%v\n",
        cmd.cpOption,cmd.class,cmd.args)
}