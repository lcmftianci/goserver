// fileio project main.go
package main
 
import (
	"bufio"
	"io"
	"log"
	"os"
	"syscall"
	"path/filepath"
    "fmt"
    "flag"
)

// func copy(src, dst string) (int64, error) {
// 	sourceFileStat, err := os.Stat(src)
// 	if err != nil {         
// 		return 0, err
// 	}       
// 	if !sourceFileStat.Mode().IsRegular() {     
// 		return 0, fmt.Errorf("%s is not a regular file", src)
// 	}

// 	source, err := os.Open(src) 
// 		if err != nil {                
// 			return 0, err
// 	}
// 	defer source.Close()

// 	destination, err := os.Create(dst)        
// 		if err != nil {                
// 			return 0, err
// 	}
// 	defer destination.Close()
// 	nBytes, err := io.Copy(destination, source)        
// 	return nBytes, err
// }

// func mycopy(sourceFile, dst string){
// 		input, err := ioutil.ReadFile(sourceFile)        
// 		if err != nil {
// 			fmt.Println(err)                
// 			return
// 	}

// 	err = ioutil.WriteFile(destinationFile, input, 0644)        
// 		if err != nil {
// 			fmt.Println("Error creating", destinationFile)
// 			fmt.Println(err)                
// 			return
// 	}
// }

// func mycopy1(sourceFile, dst string){
// 	buf := make([]byte, BUFFERSIZE)        
// 	for {
// 		n, err := source.Read(buf)                
// 		if err != nil && err != io.EOF {                        
// 			return err
// 		}                
// 		if n == 0 {                        
// 			break
// 		}               
// 		if _, err := destination.Write(buf[:n]); err != nil {                        
// 			return err
// 		}
// 	}
// }

func getFilelist(path string) {
    err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
        if f == nil {
			return err
		}
        if f.IsDir() {
			return nil
		}
        println(path)
        return nil
        })
    if err != nil {
        fmt.Printf("filepath.Walk() returned %v\n", err)
    }
}

func main() {
	flag.Parse()
    root := flag.Arg(0)
    getFilelist(root)
	fr, err := os.OpenFile("E:/file/data/data.txt", syscall.O_RDONLY, 0)
	if err != nil {
		log.Fatal(err)
	}
	defer fr.Close()
 
	fw, err := os.OpenFile("E:\\file\\data\\result1.txt", syscall.O_CREAT|syscall.O_WRONLY, 0777)
	if err != nil {
		log.Fatal(err)
	}
	defer fw.Close()
 
	buf := bufio.NewReader(fr)
	for {
		line, err := buf.ReadBytes('\n')
		fw.Write(line)
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
	}
}
