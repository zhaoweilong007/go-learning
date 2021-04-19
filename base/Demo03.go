package main

import "fmt"

func main() {
	//slice切片 动态数组，并不是真正意义上的动态数组，而是指向一个底层的array

	//和申明array一样只是少了长度
	var aslice = []int{1, 2, 3, 34}
	fmt.Println(aslice)

	//从array中申明slice,slice[i:j],i是开始位置，j是结束位置，不包含j
	var bt = []byte{'a', 'b', 'x', 'w', 'g', 's'}
	fmt.Println(bt)
	s1 := bt[0:3]
	s2 := bt[3:6]

	fmt.Println(s1, s2)

	//等价于
	s3 := bt[:3]
	s4 := bt[3:]
	s5 := bt[:] //包含所有元素
	fmt.Println(s3, s4, s5)

	//slice内置函数使用
	/**
	len:获取slice的长度
	cap:获取slice的最大容量
	append:向slice里面追加一个或者多个元素，然后返回一个和slice一样类型的slice
	*/

	fmt.Printf("s3 len%s,cap%s\n", len(s3), cap(s3))
	fmt.Printf("s4 len%s,cap%s\n", len(s4), cap(s4))

	arr1 := new([3]int)
	arr2 := [3]int{1, 2, 3}

	//拷贝
	arr1 = &arr2
	fmt.Printf("arr1:%v\n", arr1)
	fmt.Printf("arr2:%v\n", arr2)

	arr1[0] = 3
	fmt.Printf("arr2:%v\n", arr2)

	arr2[2] = 5
	fmt.Printf("arr1:%v\n", arr1)

	//指定索引位置赋值
	str := [...]string{1: "hello", 4: "my body"}
	fmt.Printf("str:%v\n", str)

	arr3 := make([]int, 10, 20)
	for i := 0; i < 10; i++ {
		arr3[i] = i
	}
	fmt.Println("arr3:", arr3)

}
