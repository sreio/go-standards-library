package main

import (
	"fmt"
	"unicode"
)

/**
func IsControl(r rune) bool  // 是否控制字符
func IsDigit(r rune) bool  // 是否阿拉伯数字字符，即 0-9
func IsGraphic(r rune) bool // 是否图形字符
func IsLetter(r rune) bool // 是否字母
func IsLower(r rune) bool // 是否小写字符
func IsMark(r rune) bool // 是否符号字符
func IsNumber(r rune) bool // 是否数字字符，比如罗马数字Ⅷ也是数字字符
func IsOneOf(ranges []*RangeTable, r rune) bool // 是否是 RangeTable 中的一个
func IsPrint(r rune) bool // 是否可打印字符
func IsPunct(r rune) bool // 是否标点符号
func IsSpace(r rune) bool // 是否空格
func IsSymbol(r rune) bool // 是否符号字符
func IsTitle(r rune) bool // 是否 title case
func IsUpper(r rune) bool // 是否大写字符
func Is(rangeTab *RangeTable, r rune) bool // r 是否为 rangeTab 类型的字符
func In(r rune, ranges ...*RangeTable) bool  // r 是否为 ranges 中任意一个类型的字符
*/

func main() {
	single := '\u0015'
	fmt.Println(string(single), unicode.IsControl(single))
	single = '\ufe35'
	fmt.Println(string(single), unicode.IsControl(single))

	digit := '1'
	fmt.Println(digit, unicode.IsDigit(digit))
	fmt.Println(digit, unicode.IsNumber(digit))

	letter := 'Ⅷ'
	fmt.Println(string(letter), unicode.IsDigit(letter))
	fmt.Println(string(letter), unicode.IsNumber(letter))

	han := '你'
	fmt.Println(string(han), unicode.IsDigit(han))
	fmt.Println(string(han), unicode.Is(unicode.Han, han))
	fmt.Println(string(han), unicode.In(han, unicode.Gujarati, unicode.White_Space))
}
