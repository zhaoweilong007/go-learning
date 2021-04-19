package main

import "fmt"

func main() {
	//指针
	//取值符&,返回变量的内存地址,反引用*,或者叫内容/间接引用，指针转移
	var s1 string = "dasdasda"
	s2 := &s1
	fmt.Println("s1 &=>", s2)
	fmt.Println("s2 *=>", *s2)
	test(s1)
	fmt.Println("s1=>%s", s1)
	test2(&s1)
	fmt.Println("s1=>%s", s1)

	s := "see you"
	var p *string = &s
	*p = "good bye"
	fmt.Printf("p=>%s\n", p)
	fmt.Printf("*p=>%s\n", *p)
	fmt.Printf("s=>%s\n", s)

}

func test(str string) {
	str = "qeqweqw"
}

func test2(str *string) {
	s1 := "xkhjkhejkqe"
	str = &s1
}
