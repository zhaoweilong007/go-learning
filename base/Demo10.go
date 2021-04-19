package main

import (
	"bytes"
	"fmt"
)

func main() {
	//接口与反射
	var v Valuable = stockPosition{
		ticker:     "G00G",
		sharePrice: 388.12,
		count:      5,
	}

	//类型断言
	if t, ok := v.(Valuable); ok {
		showValue(t)
	}

	v = &car{"QWE", "M@", 12312.12}
	if _, ok := v.(Valuable); ok {
		showValue(v)
	}

	var lst List

	if LongEnough(lst) {
		fmt.Printf("- lst is long enough\n")
	}

	plst := new(List)
	CountInto(plst, 1, 10)
	if LongEnough(plst) {
		fmt.Printf("- plst is long enough\n")
	}
}

type Valuable interface {
	getValue() float32
}

type stockPosition struct {
	ticker     string
	sharePrice float32
	count      float32
}

func (s stockPosition) getValue() float32 {
	return s.sharePrice * s.count
}

type car struct {
	make  string
	model string
	price float32
}

func (c car) getValue() float32 {
	return c.price
}

func showValue(asset Valuable) {
	fmt.Println("value of the asset is:", asset.getValue())
}

//接口嵌套

type ReadWrite interface {
	Read(b bytes.Buffer) bool
	Write(b bytes.Buffer) bool
}

type Lock interface {
	Lock()
	UnLock()
}

type File interface {
	ReadWrite
	Lock
	Close()
}

//使用方法集与接口
type List []int

func (l List) Len() int {
	return len(l)
}

func (l *List) Append(val int) {
	*l = append(*l, val)
}

type Appender interface {
	Append(int)
}

func CountInto(a Appender, start, end int) {
	for i := start; i <= end; i++ {
		a.Append(i)
	}
}

type Lener interface {
	Len() int
}

func LongEnough(l Lener) bool {
	return l.Len()*10 > 42
}
