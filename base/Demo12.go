package main

import (
	"fmt"
	"reflect"
)

type NotKnowType struct {
	s1, s2, s3 string
}

func (n *NotKnowType) string() string {
	return n.s1 + "-" + n.s2 + "-" + n.s3
}

var secret interface{} = NotKnowType{"java", "go", "python"}

func main() {

	//反射
	var f float32 = 3.1415926
	fmt.Println("f type:", reflect.TypeOf(f))
	v := reflect.ValueOf(f)
	fmt.Println("f valueOf:", v)
	fmt.Println("f type", v.Type())
	fmt.Println("f kind", v.Kind())
	fmt.Println("f value", v.Float())
	fmt.Println("f interface", v.Interface().(float32))

	fmt.Println("======================")
	v2 := reflect.ValueOf(&f)

	fmt.Println("v2 canSet:", v2.CanSet())

	v2 = v2.Elem()

	fmt.Println("the elem v2 is:", v2)

	fmt.Println("v2 canSet:", v2.CanSet())
	v2.SetFloat(4.232131)
	fmt.Println("v2 interface :", v2.Interface())
	fmt.Println("v2 value :", v2)

	value := reflect.ValueOf(secret)

	for i := 0; i < value.NumField(); i++ {
		fmt.Printf("field %d : %v \n", i, value.Field(i))
	}

	//res := value.Method(0).Call(nil)
	//fmt.Println("resL", res)

}
