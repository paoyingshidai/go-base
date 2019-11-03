package main

import (
	"fmt"
	"reflect"
)

/*
	反射，

*/

type User struct {
	Id   int
	Name string
	Age  int
}

func (u User) Hello() {
	fmt.Println("hello world")
}

func Info(o interface{}) {

	t := reflect.TypeOf(o)
	fmt.Println("type: ", t.Name())

	v := reflect.ValueOf(o)
	fmt.Println("Fields: ")

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		val := v.Field(i).Interface()
		fmt.Println(f.Name, f.Type, val)
	}

	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Println(m.Name, m.Type)
	}

}

func main() {

	u := User{1, "OK", 12}
	Info(u)
}
