package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	//条件及控制语句
	if f := math.Ceil(123.23); f > 123 {
		fmt.Println(f)
	}

	switch s, err := qwe(43423, "dadasd"); {
	case !err:
		fmt.Println("error")
	case s == "dad":
		fmt.Println(s)
	default:
		fmt.Println("default")
	}

	for i := 0; i < 10; i++ {
		for j := 0; j < i; j++ {
			fmt.Print("G")
		}
		fmt.Println()
	}

	var i int = 5
	for i >= 0 {
		i = i - 1
		fmt.Println(i)
	}

	//for range 注意 对val始终为集中中对应索引的拷贝，对他所作的修改不会影响到集中中原有的值
	arr := []string{"qwe", "eqwe", "qe", "iop", "hj"}
	for index, val := range arr {
		fmt.Println("index=%d,val=%s", index, val)
		val = "123"
	}
	fmt.Printf("arr=>%v", arr)

	//标签与goto

	LABLE1:
		for i := 0; i < 5; i++ {
			for j := 0; i < 5; j++ {
				if j==4 {
					continue LABLE1
				}
				fmt.Printf("i:%d,j:%d\n",i,j)
			}
			
		}


}

func qwe(i int64, str string) (string, bool) {
	s1 := strconv.FormatInt(i, 32)
	return s1 + str, true
}
