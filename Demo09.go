package main

import (
	"fmt"
	"reflect"
)

//struct
type person struct {
	name   string "姓名"
	age    int    "年龄"
	string        //anonymous 每一种数据类型只能有一个匿名字段
	personInfo
}

type personInfo struct {
	desc string
}

type p0 person

func (p *person) callPerson() {
	p.name = "hello " + p.name
	p.age = p.age - 1
}

func (p *personInfo) callPersonInfo() {
	fmt.Println("调用callPersonInfo")
}

//重写内嵌类型的方法
func (p *person) callPersonInfo() {
	fmt.Println("重写callPersonInfo")
}

func main() {

	//使用new创建自定义结构体，new会分配内存并初始化零值，返回已分配内存的指针
	p1 := new(person)
	p1.name = "张三"
	p1.age = 23

	//等价于new
	p2 := &person{
		name: "李四",
		age:  12,
	}

	fmt.Println("p1:", p1)
	fmt.Println("p2:", p2)

	p3 := person{
		name: "王五",
		age:  23,
	}
	fmt.Println("p3：", p3)
	fmt.Println("p3：", &p3)

	p := p0{name: "赵六", age: 18}
	fmt.Println("p:", p)

	p_ := person(p)
	fmt.Println("p_", p_)
	fmt.Println("&p:", &p)
	fmt.Println("&p_:", &p_)

	p_.name = "瞎子"
	fmt.Println("p_", p_)
	fmt.Println("p:", p)

	typeOf := reflect.TypeOf(p)
	for i := 0; i < 2; i++ {
		field := typeOf.Field(i)
		fmt.Printf("index:%v,tag:%v,type:%v,anonymous:%v,path:%v\n", field.Index, field.Tag, field.Type, field.Anonymous, field.PkgPath)
	}

	p.string = "this is anonymous field"
	p.personInfo = personInfo{"死亡如风，常伴吾身"}
	p.callPersonInfo()
	fmt.Println("p:", p)

	akl := &person{
		name:   "阿卡丽",
		age:    23,
		string: "ap",
		personInfo: personInfo{
			desc: "有时候刀比魔法更好用",
		},
	}

	akl.callPersonInfo()

	v := new(Voodoo)
	v.magic()
	v.moreMagic()
}

type Base struct {
}

func (Base) magic() {
	fmt.Println("base magic")
}

func (self Base) moreMagic() {
	self.magic()
	self.magic()
}

type Voodoo struct {
	Base
}

func (Voodoo) magic() {
	fmt.Println("voodoo magic")
}
