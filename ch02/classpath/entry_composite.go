package classpath

import "errors"
import "strings"

type CompositeEntry []Entry

//创建一个CompositeEntry实例
func newCompositeEntry(pathList string) CompositeEntry {
    compositeEntry := []Entry{}
    //将传入的pathList按分隔符分成小路径
    for _, path := range strings.Split(pathList, pathListSeparator) {
        entry := newEntry(path)
        compositeEntry = append(compositeEntry, entry)
    }
    return compositeEntry
}

//遍历并调用每个子路径的readClass方法，读取class数据并返回
func (self CompositeEntry) readClass(className string) ([]byte, Entry, error) {
    for _, entry := range self {
        data, from, err := entry.readClass(className)
        if err == nil {
            return data, from, nil
        }
    }
    return nil, nil, errors.New("class not found: " + className)
}

//调用每个子路径的String，再拼接返回
func (self CompositeEntry) String() string {
    strs := make([]string, len(self))
    for i, entry := range self {
        strs[i] = entry.String()
    }

    return strings.Join(strs, pathListSeparator)
}