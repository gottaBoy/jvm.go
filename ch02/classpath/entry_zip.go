package classpath

import "archive/zip"
import "errors"
import "io/ioutil"
import "path/filepath"

type ZipEntry struct {
    //存放绝对路径
    absPath string
}

//创建一个ZipEntry实例
func newZipEntry(path string) *ZipEntry {
    absPath, err := filepath.Abs(path)
    if err != nil {
        panic(err)
    }
    return &ZipEntry{absPath}
}

//从zip文件中提取class文件
func (self *ZipEntry) readClass(className string) ([]byte, Entry, error) {
    //利用archive/zip包打开zip文件，出错则返回错误
    r, err := zip.OpenReader(self.absPath)
    if err != nil {
        return nil, nil, err
    }

    //defer保证在return前执行
    defer r.Close()
    //遍历指定路径中的File
    for _, f := range r.File {
        //如果找到与className相同的文件，读出rc（为ReadCloser接口提供读取文件内容的方法），如果出错，返回错误信息
        if f.Name == className {
            rc, err := f.Open()
            if err != nil {
                return nil, nil, err
            }
            //defer保证在return前执行，即保证关闭
            defer rc.Close()
            //通过rc读出其中内容为data，返回
            data, err := ioutil.ReadAll(rc)
            if err != nil {
                return nil, nil, err
            }
            return data, self, nil
        }
    }
    //遍历完成，没有找到对应的文件，返回class not found信息
    return nil, nil, errors.New("class not found: " + className)
}

func (self *ZipEntry) String() string {
    return self.absPath
}