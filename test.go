package main

import (
	"fmt"
	// "reflect"
)
func f1() (result int) {
    defer func() {
        result++
    }()
    return 0
}
func f2() (r int) {
	t := 5
	defer func() {
	  t = t + 5
	}()
	return t
}
func f3() (r int) {
    defer func(r int) {
          r = r + 5
    }(r)
    return 1
}
func main() {
	// type User struct {
	// 	a string
	// 	n string
	// }
	// user := User{"aaaaa","nnnnn"}
	// rvalue := reflect.ValueOf(user)
	// for i := 0; i < rvalue.NumField(); i++ {
	// 	elevalue := rvalue.Field(i)
	// 	fmt.Println("element ", i, " its type is ", elevalue.Type())
	// 	fmt.Println("element ", i, " its value is ", elevalue)
	// }
	// ReflectStructElem(user)

	fmt.Println("f1:%d",f1()) //1
	fmt.Println("f2:%d",f2()) //10错了，应该是5 他是在return后操作的？
	fmt.Println("f3:%d",f3()) //6s


	// type MyStruct struct {
    //     N int
	// }
	// n := MyStruct{ 1 }
	// // get
	// immutable := reflect.ValueOf(n)
	// val := immutable.FieldByName("N")
	// fmt.Printf("N=%d\n", val) // prints 1
		// ReflectStructElem(user)
}


