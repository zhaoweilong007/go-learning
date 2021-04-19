package main

import (
	"fmt"
)

func main() {

	//map和struct
	m1 := map[string]string{}

	m1["name"] = "张三"
	m1["age"] = "16"

	m2 := make(map[int]string)

	m2[0] = "qwe"
	m2[1] = "asd"

	fmt.Println("m1:", m1)
	fmt.Println("m2:", m2)

	m3 := map[string]string{
		"ad": "卡莎",
		"ap": "阿卡丽",
	}

	fmt.Println("m3:", m3)

	key, ok := m3["ad"]
	if ok {
		fmt.Println("key:", key)
	}

	delete(m3, "ad")

	//遍历map
	for key, val := range m3 {
		fmt.Println("key:", key, ",val:", val)
	}

	items := make([]map[int]int, 5)

	for i := range items {
		items[i] = make(map[int]int, 1)
		items[i][i] = i + 1
	}
	fmt.Println("items:", items)

}
