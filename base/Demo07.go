package main

import (
	"fmt"
	"log"
)

func main() {
	//函数

	/**
	- 普通的带有名字的函数
	- 匿名函数或者lambda
	- 方法

	ps:go中不支持重载，不支持泛型，实现泛型可以通过interface或者反射的方式实现，但不推荐这么做，性能比较低
	*/

	fun2(1, 2, "da", "da")

	//匿名函数 闭包
	fun3 := func(x, y int) int { return x * y }

	fun3(123, 123)

	//函数作为返回值
	f6 := func6()
	fmt.Println("10:", f6(10))
	fmt.Println("200:", f6(200))
	fmt.Println("300:", f6(300))

	persision(123, func(i int) {
		fmt.Println("sucess", i)
	}, func() {
		fmt.Println("error")
	})

	//defer 延迟函数，在return之前执行某个语句或函数，类似于java中的finally语句块
	fmt.Println(f())

	//函数参数传递默认是按值传递，传递的是参数的副本，使用&参数传递参数的引用
	//切片（slice）、字典（map）、接口（interface）、通道（chaannel）默认使用引用传递
	a, b, c := 123, 341, 0
	func4(a, b, &c)
	fmt.Println(c)

	fmt.Println("==============")
	funcb()

	//内置函数
	//new、make函数，用于分配内存，new用于值类型和用户类型的定义，如自定义类型，make用于内置引用类型，

	//panic、recover用于错误处理机制

	var where = log.Println

	where("qeqweqweqw")

}

//以申明的形式创建函数
type fun1 func(s1, s2 string) string

func fun2(a, b int, s0, s1 string) (int, string) {
	return a + b, s0 + s1
}

//定义一个没有形参名的函数
func func3(int, int) {

}

func func4(a, b int, c *int) {
	*c = a + b
}

func func5(args ...string) {

}

func func6() func(i int) int {
	var x int
	return func(i int) int {
		x += i
		return x
	}
}

func persision(i int, sucess func(i int), error func()) {
	if i%2 == 2 {
		sucess(i)
	} else {
		error()
	}
}

func f() (ret int) {
	defer func() {
		ret++
		fmt.Println("execute...", ret)
	}()
	fmt.Println("return...", ret)
	return 1
}

func trace(s string) string {
	fmt.Println("entering:", s)
	return s
}

func un(s string) {
	fmt.Println("leaving:", s)
}

func q() {
	defer un(trace("a"))
	fmt.Println("in a")
}

func funcb() {
	defer un(trace("b"))
	fmt.Println("in b")
	q()
}
