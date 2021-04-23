package main

import (
	"fmt"
	"reflect"
)

func main() {
	type User struct {
		a string
		n string
	}
	user := User{"aaaaa","nnnnn"}
	rvalue := reflect.ValueOf(user)
	for i := 0; i < rvalue.NumField(); i++ {
		elevalue := rvalue.Field(i)
		fmt.Println("element ", i, " its type is ", elevalue.Type())
		fmt.Println("element ", i, " its value is ", elevalue)
	}
	// ReflectStructElem(user)



	type MyStruct struct {
        N int
	}
	n := MyStruct{ 1 }
	// get
	immutable := reflect.ValueOf(n)
	val := immutable.FieldByName("N")
	fmt.Printf("N=%d\n", val) // prints 1
		// ReflectStructElem(user)
}


