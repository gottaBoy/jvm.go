package classpath

import "io/ioutil"
import "path/filepath"

type DirEntry struct {
    absDir string    //存放绝对路径
}

//用path创建一个DirEntry实例并返回
func newDirEntry(path string) *DirEntry{
    //将path转换为绝对路径，如果出错则panic，无错则创建DirEntry实例并返回
    absDir,err:=filepath.Abs(path)
    if err!=nil{
        panic(err)
    }
    return &DirEntry{absDir}
}

//将指定class的内容读出
func (self *DirEntry) readClass(className string) ([]byte,Entry,error){
    //讲绝对路径和文件名拼接在一起，并使用ioutil包读取该指定文件内容，返回结果
    fileName :=filepath.Join(self.absDir,className)
    data,err:=ioutil.ReadFile(fileName)
    return data,self,err
}

func (self *DirEntry) String() string{
    return self.absDir
}