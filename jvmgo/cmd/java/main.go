package main

import (
	"fmt"
	"os"
	"runtime/pprof"

	"github.com/gottaBoy/jvm.go/jvmgo/classpath"
	"github.com/gottaBoy/jvm.go/jvmgo/interpreter"
	"github.com/gottaBoy/jvm.go/jvmgo/jutil"
	_ "github.com/gottaBoy/jvm.go/jvmgo/native"
	"github.com/gottaBoy/jvm.go/jvmgo/options"
	"github.com/gottaBoy/jvm.go/jvmgo/rtda"
	"github.com/gottaBoy/jvm.go/jvmgo/rtda/heap"
)

func main() {
	cmd, err := parseCommand(os.Args)
	if err != nil {
		fmt.Println(err)
		printUsage()
	} else {
		startJVM(cmd)
	}
}

func printUsage() {
	fmt.Println("usage: jvmgo [-options] class [args...]")
}

func startJVM(cmd Command) {
	Xcpuprofile := cmd.Options.Xcpuprofile
	if Xcpuprofile != "" {
		f, err := os.Create(Xcpuprofile)
		if err != nil {
			panic(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	options.InitOptions(cmd.Options.VerboseClass, cmd.Options.Xss, cmd.Options.XuseJavaHome)

	cp := classpath.Parse(cmd.Options.Classpath)
	heap.InitBootLoader(cp)

	mainClassName := jutil.ReplaceAll(cmd.Class, ".", "/")
	mainThread := createMainThread(mainClassName, cmd.Args)
	interpreter.Loop(mainThread)
	interpreter.KeepAlive()
}

func createMainThread(className string, args []string) *rtda.Thread {
	mainThread := rtda.NewThread(nil)
	bootMethod := heap.BootstrapMethod()
	bootArgs := []interface{}{className, args}
	mainThread.InvokeMethodWithShim(bootMethod, bootArgs)
	return mainThread
}
