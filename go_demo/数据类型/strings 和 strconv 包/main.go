package main

import (
	"fmt"
	"strings"
)

func main() {
	//HasPrefix 判断字符串 s 是否以 prefix 开头：
	flavor := "hw:nama_notes"

	if strings.HasPrefix(flavor, "hw") {
		fmt.Println("It's a 'hw' flavor.")
	} else {
		fmt.Println("Unknown flavor.")
	}
	//HasSuffix 判断字符串 s 是否以 suffix 结尾：
	rock := "hya:hello ha"
	if strings.HasSuffix(rock, "ha") {
		fmt.Println("rock is ha rock")
	} else {
		fmt.Println("rock ha ha!")
	}
	//下属都是后在前的索引
	//Contains 判断字符串 s 是否包含 string：
	str1 := "hello world"
	str2 := "hello"
	fmt.Println(strings.Contains(str1, str2))

	//Index 返回字符串 str 在字符串 s 中的索引（str 的第一个字符的索引），-1 表示字符串 s 不包含字符串 str：
	str3 := "somethings"
	str4 := "s"
	fmt.Println(strings.Index(str3, str4))

	//LastIndex 返回字符串 str 在字符串 s 中最后出现位置的索引（str 的第一个字符的索引），-1 表示字符串 s 不包含字符串 str：
	fmt.Println(strings.LastIndex(str3, str4))

	//定位总结：如果定位的不是Ascall中的，建议使用IndexRune定位

	//Replace 用于将字符串 str 中的前 n 个字符串 old 替换为字符串 new，并返回一个新的字符串，如果 n = -1 则替换所有字符串 old 为字符串 new：
	//strings.Replace(str, old, new string, n int) string
	fmt.Println(strings.Replace("hello", "oo", "string", -1))

	//Count 用于计算字符串 str 在字符串 s 中出现的非重叠次数：
	//strings.Count(s, str string) int
	fmt.Println(strings.Count(str3, str4))

	//Repeat 用于重复 count 次字符串 s 并返回一个新的字符串：
	//strings.Repeat(s, count int) string
	fmt.Println(str4, 5)

	//ToLower 将字符串中的 Unicode 字符全部转换为相应的小写字符：
	//strings.ToLower(s) string
	//ToUpper 将字符串中的 Unicode 字符全部转换为相应的大写字符：
	//strings.ToUpper(s) string
	fmt.Println(strings.ToLower(str3))
	fmt.Println(strings.ToUpper(str4))

	//针对从数字类型转换到字符串，Go 提供了以下函数：

	/* strconv.Itoa(i int) string 返回数字 i 所表示的字符串类型的十进制数。
	   strconv.FormatFloat(f float64, fmt byte, prec int, bitSize int) string 将 64 位浮点型的数字转换为字符串，其中 fmt 表示格式（其值可以是 'b'、'e'、'f' 或 'g'），prec 表示精度，bitSize 则使用 32 表示 float32，用 64 表示 float64。
	   将字符串转换为其它类型 tp 并不总是可能的，可能会在运行时抛出错误 parsing "…": invalid argument。

	   针对从字符串类型转换为数字类型，Go 提供了以下函数：

	   strconv.Atoi(s string) (i int, err error) 将字符串转换为 int 型。
	   strconv.ParseFloat(s string, bitSize int) (f float64, err error) 将字符串转换为 float64 型。
	   利用多返回值的特性，这些函数会返回 2 个值，第 1 个是转换后的结果（如果转换成功），第 2 个是可能出现的错误，因此，我们一般使用以下形式来进行从字符串到其它类型的转换
	*/

}
