package classpath

import "os"
import "path/filepath"

type Classpath struct {
    bootClasspath Entry    //启动类路径，默认为jre/lib目录
    extClasspath Entry//扩展类路径，默认为jre/lib/ext
    userClasspath Entry//用户类路径
}

//解析启动类路径和扩展类路径
func Parse(jreOption, cpOption string) *Classpath {
    cp := &Classpath{}
    cp.parseBootAndExtClasspath(jreOption)//找启动类路径和扩展类路径
    cp.parseUserClasspath(cpOption)//找用户类路径
    return cp
}


func (self *Classpath) parseBootAndExtClasspath(jreOption string) {
    jreDir := getJreDir(jreOption)

    // jre/lib/*
    jreLibPath := filepath.Join(jreDir, "lib", "*")
    self.bootClasspath = newWildcardEntry(jreLibPath)

    // jre/lib/ext/*
    jreExtPath := filepath.Join(jreDir, "lib", "ext", "*")
    self.extClasspath = newWildcardEntry(jreExtPath)
}

//找jre目录
func getJreDir(jreOption string) string {
    //找用户输入的路径jre
    if jreOption != "" && exists(jreOption) {
        return jreOption
    }
    //找当前目录下jre
    if exists("./jre") {
        return "./jre"
    }
    //找JAVA_HOME中jre
    if jh := os.Getenv("JAVA_HOME"); jh != "" {
        return filepath.Join(jh, "jre")
    }
    //都找不到，返回panic
    panic("Can not find jre folder!")
}

//判断目录是否存在
func exists(path string) bool {
    if _, err := os.Stat(path); err != nil {
        if os.IsNotExist(err) {
            return false
        }
    }
    return true
}

//用户未输入-classpath/-cp参数，默认使用当前目录作为用户类路径
func (self *Classpath) parseUserClasspath(cpOption string) {
    if cpOption == "" {
        cpOption = "."
    }
    self.userClasspath = newEntry(cpOption)
}

//依次从启动类路径、扩展类路径和用户类路径中搜索class文件
func (self *Classpath) ReadClass(className string) ([]byte, Entry, error) {
    className = className + ".class"
    if data, entry, err := self.bootClasspath.readClass(className); err == nil {
        return data, entry, err
    }
    if data, entry, err := self.extClasspath.readClass(className); err == nil {
        return data, entry, err
    }
    return self.userClasspath.readClass(className)
}

//返回用户路径的字符串表示
func (self *Classpath) String() string {
    return self.userClasspath.String()
}