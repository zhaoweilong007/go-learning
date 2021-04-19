package main

import (
	"errors"
	"fmt"
)

var v1 int

//初始化函数
func init() {
	v1 = 123
}

func main() {
	//变量申明
	var str string = "zhangsan"
	fmt.Println(str)

	var a, b, c = 1, 2, 3
	fmt.Println(a, b, c)

	name := "lisi"
	fmt.Println(name)

	//常量申明
	const val = "const value"

	//数据类型申明
	flag := true
	fmt.Println("flag:{}", flag)

	var i1 int = 12
	var i2 int8 = 123
	var i3 rune = 1234556

	fmt.Println(i1, i2, i3)
	var f1 float32 = 23.123
	var f2 float64 = 123.123123
	fmt.Println(f1, f2)

	var str1 string
	str2 := "str2"
	fmt.Println(str1, str2)

	//申明多行字符串
	str3 := `hello
			go`

	fmt.Println(str3)

	err := errors.New("NullPrintException")
	if err != nil {
		fmt.Println(err)
	}

	//分组创建
	const (
		qwe = 123
		asd = "dada"
	)

	var (
		i int     = 111
		s string  = "qddad"
		f float64 = 123.123
	)
	fmt.Println(i, s, f)

	//数组 数组不可动态变化，一旦申明，其长度就是固定的
	//一维数组
	var arr_1 [3]int
	fmt.Println(arr_1)

	arr2 := []string{"a", "b", "c"}
	fmt.Println(arr2)

	// 使用...来自动计算数组长度
	arr3 := [...]int{1, 2, 34, 5, 6, 2, 2}
	fmt.Println(arr3)

	arr4 := [8]int{0: 3, 4: 5, 7: 9}
	fmt.Println(arr4)

	//二维数组
	arr5 := [2][3]int{{1, 2, 3}, {3, 2, 5}}
	fmt.Println(arr5)

	arr6 := [...][2]int{{12, 2}, {213, 23}}

	fmt.Println(arr6)

}
