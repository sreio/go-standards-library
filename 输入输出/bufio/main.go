package main

import (
	"bufio"
	"fmt"
	"strings"
)

/**
bufio.NewReader()
bufio.NewReaderSize()
bufio.NewReadWriter()
bufio.NewWriter()
bufio.NewWriterSize()
bufio.NewScanner()
bufio.ScanBytes()
bufio.ScanLines()
bufio.ScanRunes()
bufio.ScanWords()
*/
func main() {
	fmt.Println("bufio \n")

	reader := bufio.NewReader(strings.NewReader("http://studygolang.com. \nIt is the home of gophers"))
	line, _, _ := reader.ReadLine()
	line1, _ := reader.ReadSlice('\n')
	fmt.Println(string(line), string(line1))
}
