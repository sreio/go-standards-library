package main

import "fmt"

// 初始化变量
func main() {
	var a int
	var b string
	var c bool
	var d float64
	var e float32
	a = 1
	b = "bb"
	c = true
	d = 12.2
	e = 1.12
	fmt.Println(a, b, c, d, e)

	var arr = [1]int{1}
	var arrs = [2]string{"a", "b"}
	var arrsl = [2][3]int{{1, 2}, {1}} // 未指定的值使用类型的零值
	var arrsl2 = [2][]string{{"a1", "a2"}, {"a"}}

	// append copy
	var sl = []int{1}
	var sl2 = []string{"aa", "bb"}

	// delete
	var ma = map[int]string{1: "a1", 4: "fs"}
	var ma1 = map[string]string{"a": "aa", "b": "bb"}

	fmt.Println(arr, arrs, arrsl, arrsl2, sl, sl2, ma, ma1)
}
