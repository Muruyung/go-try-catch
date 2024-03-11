package main

import (
	"errors"
	"fmt"

	. "github.com/Muruyung/go-try-catch"
)

type Test struct {
	val interface{}
}

func main() {
	Try(func() {
		ThrowNewException(1)
	}).Catch(func(i int) {
		fmt.Println("This is an int exception:", i)
	})

	Try(func() {
		ThrowNewException("hello world")
	}).Catch(func(s string) {
		fmt.Println("This is a string exception:", s)
	})

	Try(func() {
		err := errors.New("this is an error exception")
		ThrowNewException(err)
	}).Catch(func(e error) {
		fmt.Println(e)
	})

	Try(func() {
		test := Test{
			val: "This is a struct exception",
		}
		ThrowNewException(test)
	}).Catch(func(t Test, st StackTrace) {
		fmt.Println(t)
		fmt.Println("This is a stack trace result")
		st.Print()
	})

	exception := Try(func() {
		err := errors.New("this is an error exception")
		ThrowNewException(err)
	}).Catch(func(e error) {
		fmt.Println("This is an example for return exception")
	})

	fmt.Println(exception.Error())
	fmt.Println(exception.GetStackTrace())
}
