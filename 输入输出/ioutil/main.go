package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

/**
io/ioutil 方便的 IO 操作函数集

ioutil.ReadDir()  读取当前路径下的文件名、目录名
ioutil.ReadAll()
ioutil.ReadFile()
ioutil.WriteFile()
*/
func main() {

	//------ioutil.ReadFile()---------
	content, err := ioutil.ReadFile("./一级目录/二级文件")
	if err != nil {
		fmt.Println(err)
	}
	// T(Value) T:变量类型  byte转为string类型
	fmt.Println(string(content), "\n")

	//------ioutil.WriteFile()---------
	message := []byte("weidada")
	writeFileErr := ioutil.WriteFile("./一级目录/二级文件", message, 0644)
	if writeFileErr != nil {
		fmt.Println(err)
	}
	// 追加
	fl, err := os.OpenFile("./一级目录/二级文件", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
	}
	defer fl.Close()
	n, err := fl.Write([]byte("\napped weidada"))
	if err != nil {
		fmt.Println(err)
	}
	if err == nil {
		fmt.Println(n)
	}

	//------ioutil.ReadAll()---------
	file, err := os.Open("./一级目录/二级文件")
	defer file.Close() // 关闭文件句柄
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(file) //返回的是一个文件地址 如：&{0x14000126180}
	// or
	r := strings.NewReader("Go is a general-purpose language designed with systems programming in mind.")

	content1, err := ioutil.ReadAll(file) //读取文件
	if err != nil {
		fmt.Println(err)
	}
	content2, err := ioutil.ReadAll(r)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(content1), "\n", string(content2), "\n")

	//------ioutil.ReadDir()---------
	dirInfo, err := ioutil.ReadDir("./")
	if err != nil {
		fmt.Println(err)
	}
	for _, v := range dirInfo {
		/**
		v.Name()	文件名称				go.mod
		v.Size()	文件大小				21
		v.IsDir()	是否是目录			false
		v.Mode()	文件或目录权限、时间	-rw-r--r-- 2021-12-22 17:01:51.26247689 +0800 CST
		*/
		fmt.Println(v.Name(), v.Size(), v.IsDir(), v.Mode(), v.ModTime())
	}

	fmt.Println("\n")
	listAll("./", 0)
}

func listAll(path string, curHier int) {
	fileInfos, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, info := range fileInfos {
		if info.IsDir() { //判断是否是目录
			for tmpHier := curHier; tmpHier > 0; tmpHier-- {
				fmt.Printf("|\t")
			}
			fmt.Println(info.Name(), "\\")
			listAll(path+"/"+info.Name(), curHier+1)
		} else {
			for tmpHier := curHier; tmpHier > 0; tmpHier-- {
				fmt.Printf("|\t")
			}
			fmt.Println(info.Name())
		}
	}
}
