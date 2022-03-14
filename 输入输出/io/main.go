package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	io.Copy(os.Stdout, strings.NewReader("Go语言中文网\n"))

	reader := strings.NewReader("Go语言中文网")
	reader.Seek(5, io.SeekStart)
	var str = make([]byte, 15)
	at, _ := reader.ReadAt(str, 5)
	r, _, _ := reader.ReadRune()
	fmt.Printf("%c  %s %d \n", r, str, at)
}
