package classpath

import "os"
import "path/filepath"
import "strings"

//类似CompositeEntry，不定义新的类型

//创建一个WildcardEntry实例
func newWildcardEntry(path string) CompositeEntry {
    //删除末尾*
    baseDir := path[:len(path)-1]
    compositeEntry := []Entry{}
    //根据后缀名选出jar文件，并跳过子目录
    walkFn := func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }
        if info.IsDir() && path != baseDir {
            return filepath.SkipDir
        }
        if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") {
            jarEntry := newZipEntry(path)
            compositeEntry = append(compositeEntry, jarEntry)
        }
        return nil
    }
    //遍历baseDir创建ZipEntry
    filepath.Walk(baseDir, walkFn)

    return compositeEntry
}
