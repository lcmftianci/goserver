package main

import (
	"fmt"
	"io/ioutil"
	"os"
	// "path/filepath"
	// "strings"
)

var (
	files []string
)


//获取指定目录下的所有文件和目录
func ListDir(dirPth string) ( err error) {
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
	return  err
	}
	PthSep := string(os.PathSeparator)
	//suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写
	for _, fi := range dir {
		
		if fi.IsDir() { // 忽略目录
			//files1 = append(files1, dirPth+PthSep+fi.Name())
			ListDir(dirPth + PthSep + fi.Name())
			fmt.Println(dirPth + PthSep + fi.Name())
		}else{
			//files	d = append(files, dirPth + PthSep + fi.Name())
			//fmt.Println(files, dirPth + PthSep + fi.Name())
			files = append(files, dirPth+PthSep+fi.Name())
		}
	}
	return  nil
}

func main() {
	_ = ListDir("E:/file/data/")

	for _, table1 := range files {
		fmt.Println(table1)
	}
}