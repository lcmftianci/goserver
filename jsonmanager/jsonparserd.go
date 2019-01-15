package main
import (
    "github.com/buger/jsonparser"
	"fmt"
	"log"
	"os"
	"io"
	"io/ioutil"
)

// "syscall"
// "bufio"
 
func main() {
//     data := []byte(`{
//   "person": {
//     "name":{
//       "first": "Leonid",
//       "last": "Bugaev",
//       "fullName": "Leonid Bugaev"
//     },
//     "github": {
//       "handle": "buger",
//       "followers": 109
//     },
//     "avatars": [
//       { "url": "https://avatars1.githubusercontent.com/u/14009?v=3&s=460", "type": "thumbnail" }
//     ]
//   },
//   "company": {
//     "name": "Acme"
//   }
// }`)

	// fr, err := os.OpenFile("E:\\mycode\\KafkaManager\\192.168.4.65众创空间朝阳RTSP0a\\非机动车\\13.txt", syscall.O_RDONLY, 0666)
	// if err != nil {
	// 	log.Fatal(err)
	// 	panic(err)
	// }
	// defer fr.Close()
	
	// buf := bufio.NewReader(fr)
	// fmt.Println(buf)
	// for{
	// 	//line, err := buf.ReadString('\n')
	// 	line, err := ioutil.ReadAll(buf)
	// 	fmt.Println(line)
	// 	if err == io.EOF {
	// 		break
	// 	} else if err != nil {
	// 		log.Fatal(err)
	// 	}
	// }

	filePath := "E:\\mycode\\KafkaManager\\192.168.4.65众创空间朝阳RTSP0a\\非机动车\\0_1337433360_396.json"
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

	//fmt.Println("File Is Commplete 2")

	result, err := jsonparser.GetInt(line, "otherVehicleInfo", "totalMotobike", "motobikeCount")
    if err != nil {
        fmt.Println(err)
	}
	if result == 1{
		fmt.Println(result)
	}else{
		result, err := jsonparser.GetInt(line, "otherVehicleInfo", "totalTribike", "tribikeCount")
		if err != nil {
			fmt.Println(err)
		}
		if result == 1{
			fmt.Println(result)
		}else{
			result, err := jsonparser.GetInt(line, "otherVehicleInfo", "totalBike", "bikeCount")
			if err != nil {
				fmt.Println(err)
			}
			if result == 1{
				fmt.Println(result)
			}
		}
	}
	
    // result, err := jsonparser.GetString(data, "person", "name", "fullName")
    // if err != nil {
    //     fmt.Println(err)
    // }
    // fmt.Println(result)
 
    // content, valueType, offset, err := jsonparser.Get(data, "person", "name", "fullName")
    // if err != nil {
    //     fmt.Println(err)
    // }
    // fmt.Println(content, valueType, offset)
    // //jsonparser提供了解析bool、string、float64以及int64类型的方法，至于其他类型，我们可以通过valueType类型来自己进行转化
    // result1, err := jsonparser.ParseString(content)
    // if err != nil {
    //     fmt.Println(err)
    // }
    // fmt.Println(result1)
 
    // err = jsonparser.ObjectEach(data, func(key []byte, value []byte, dataType jsonparser.ValueType, offset int) error {
    //     fmt.Printf("key:%s\n value:%s\n Type:%s\n", string(key), string(value), dataType)
    //     return nil
    // }, "person", "name")
}
