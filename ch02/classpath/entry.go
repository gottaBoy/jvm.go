package classpath

import "os"
import "strings"

//存放路径分隔符
const pathListSeparator = string (os.PathListSeparator)

//定义Entry接口
type Entry interface{
    readClass(className string)([]byte,Entry,error)//查找和加载class文件
    String() string//类似于java中toString()函数
}

//类似于java的构造函数，根据参数创建不同类型的Entry
func newEntry(path string )Entry{
    if strings.Contains(path, pathListSeparator) {
        return newCompositeEntry(path)
    }
    if strings.HasSuffix(path, "*") {
        return newWildcardEntry(path)
    }
    if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") ||
        strings.HasSuffix(path, ".zip") || strings.HasSuffix(path, ".ZIP") {

        return newZipEntry(path)
    }

    return newDirEntry(path)
}