package main

import (
    "github.com/buger/jsonparser"
	"fmt"
	"log"
	"os"
	"io"
	"io/ioutil"
	"path"
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

func ParserIWant(filePath string)(ret int64) {
	//filePath := "E:\\mycode\\KafkaManager\\192.168.4.65众创空间朝阳RTSP0a\\非机动车\\0_1337433360_396.json"
	f, err := os.Open(filePath)
 	if err != nil {
		fmt.Println("File Open Error")
	 }else{
		 fmt.Println("File Open OK")
	 }
	line, err := ioutil.ReadAll(f)
	//fmt.Print(line)
	if err == io.EOF {
		fmt.Println("File Is Commplete")
	} else if err != nil {
		log.Fatal(err)
	}

	result, err := jsonparser.GetInt(line, "otherVehicleInfo", "totalMotobike", "motobikeCount")
    if err != nil {
        fmt.Println(err)
	}
	if result == 1{
		fmt.Println(result)
		return result
	}else{
		result, err := jsonparser.GetInt(line, "otherVehicleInfo", "totalTribike", "tribikeCount")
		if err != nil {
			fmt.Println(err)
		}
		if result == 1{
			fmt.Println(result)
			return result
		}else{
			result, err := jsonparser.GetInt(line, "otherVehicleInfo", "totalBike", "bikeCount")
			if err != nil {
				fmt.Println(err)
			}
			if result == 1{
				fmt.Println(result)
				return result
			}
		}
	}
	return 0
}

func main() {
	_ = ListDir("E:\\mycode\\KafkaManager\\192.168.4.65众创空间朝阳RTSP0a\\非机动车\\")
	var nVehicle int64
	for _, table1 := range files {
		//fmt.Println(table1)
		fileSuffix := path.Ext(table1)
		fmt.Println(fileSuffix)
		if fileSuffix==".json"{
			ret := ParserIWant(table1)
			if ret == 1{
				nVehicle++
			}
		}
	}
	//nVehicle = 6
	fmt.Printf("Has Vehicle:%d", nVehicle)
}