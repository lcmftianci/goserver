package main
import (
    "fmt"
    "io/ioutil"
    "os"
    //    "path/filepath"
    //    "strings"
)
//获取指定目录下的所有文件和目录
func ListDir(dirPth string) (files []string,files1 []string, err error) {
    //fmt.Println(dirPth)
    dir, err := ioutil.ReadDir(dirPth)
    if err != nil {
        return nil,nil, err
    }
    PthSep := string(os.PathSeparator)
    //    suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写
    
    for _, fi := range dir {
        
        if fi.IsDir() { // 忽略目录
            files1 = append(files1, dirPth+PthSep+fi.Name())
            ListDir(dirPth + PthSep + fi.Name())
        }else{
            //fmt.Println("s")
            files = append(files, dirPth+PthSep+fi.Name())
        }
    }
    return files,files1, nil
}
func main() {
    files,files1, _ := ListDir("E:/file/data/")
    
    for _, table := range files1 {
        temp,_,_:=ListDir(table)
        for _,temp1 :=range temp{
            files = append(files, temp1)
        }
        
    }
    
    for _, table1 := range files {
        fmt.Println(table1)
    }
}