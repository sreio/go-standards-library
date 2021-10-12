package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"unicode/utf8"
)

func main() {
	fmt.Println(Utf8Index("伟大2a大", "伟"))
	fmt.Println(Utf8Index("伟大2a大", "大"))
	fmt.Println(Utf8Index("伟大2a大", "a"))
	fmt.Println(Utf8Index("伟大2a大", "2"))
	fmt.Println(Utf8Index("1a大大", "大大"))

	dirInfo, err := ioutil.ReadDir("./")
	if err != nil {
		fmt.Println(err)
	}
	for _, v := range dirInfo {
		fmt.Println(v.Name())
	}
	
	listAll("./", 0)
}


// strings.Index 的 UTF-8 版本
// 即 Utf8Index("Go语言中文网", "中文") 返回 4，而不是 strings.Index 的 8
func Utf8Index(str, substr string) int {
    index := strings.Index(str, substr)
    if index < 0{
        return -1
    }
    return utf8.RuneCountInString(str[:index])
}

func listAll(path string, curHier int){
    fileInfos, err := ioutil.ReadDir(path)
    if err != nil{fmt.Println(err); return}

    for _, info := range fileInfos{
        if info.IsDir(){
            for tmpHier := curHier; tmpHier > 0; tmpHier--{
                fmt.Printf("|\t")
            }
            fmt.Println(info.Name(),"\\")
            listAll(path + "/" + info.Name(),curHier + 1)
        }else{
            for tmpHier := curHier; tmpHier > 0; tmpHier--{
                fmt.Printf("|\t")
            }
            fmt.Println(info.Name())
        }
    }
}
