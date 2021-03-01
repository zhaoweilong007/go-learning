package main

import (
	"fmt"
	strconv "strconv"
	"strings"
	"time"
)

func main() {
	//strings和strconv包
	str := "abcd efgh ijkl mno pqrs tuvwx yz"

	//前缀匹配
	fmt.Println(strings.HasPrefix(str, "abc"))
	//后缀匹配
	fmt.Println(strings.HasSuffix(str, "jk"))
	//是否包含
	fmt.Println(strings.Contains(str, "e"))
	//字符串在索引的位置
	fmt.Println(strings.Index(str, "i"))
	//字符串最后索引的位置
	fmt.Println(strings.LastIndex(str, "i"))
	//字符串替换 将前n哥字符换old替换为new -1则替换所有
	fmt.Println(strings.Replace(str, "a", "1234", -1))
	//统计出现的次数
	fmt.Println(strings.Count(str, "abc"))
	//重复count次字符串
	fmt.Println(strings.Repeat(str, 1))
	//转换为大写
	fmt.Println(strings.ToUpper(str))
	//转化为小写
	fmt.Println(strings.ToLower(str))

	//修前字符串
	fmt.Println(strings.TrimLeft(str, "123"))
	//分割字符串
	fmt.Println(strings.Fields(str))
	//拼接字符串
	arr := []string{"qwe", "ad", "azcx"}
	fmt.Println(strings.Join(arr, ";"))

	count := "41412"

	s1, _ := strconv.Atoi(count)
	fmt.Println(s1)
	f1, _ := strconv.ParseFloat(count, 32)
	fmt.Println(f1)

	s2 := strconv.Itoa(s1)
	fmt.Println(s2)
	f2 := strconv.FormatFloat(f1, 'f', 2, 32)
	fmt.Println(f2)

	//时间哥日期
	t1 := time.Now()
	fmt.Println(t1)
	fmt.Println(t1.Format(time.ANSIC))
	fmt.Println(t1.Format("20210301"))

	fmt.Printf("day:%02d,week:%02d,month:%4d", t1.Day(), t1.Weekday(), t1.Month())

}
