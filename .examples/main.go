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
	// Example exception using int value
	Try(func(ThrowNewException func(any)) {
		ThrowNewException(1)
	}).Catch(func(i int) {
		fmt.Println("This is an int exception:", i)
	})

	// Example exception using string value
	Try(func(ThrowNewException func(any)) {
		ThrowNewException("hello world")
	}).Catch(func(s string) {
		fmt.Println("This is a string exception:", s)
	})

	// Example exception using error value
	Try(func(ThrowNewException func(any)) {
		err := errors.New("this is an error exception")
		ThrowNewException(err)
	}).Catch(func(e error) {
		fmt.Println(e)
	})

	// Example exception using struct data
	Try(func(ThrowNewException func(any)) {
		test := Test{
			val: "This is a struct exception",
		}
		ThrowNewException(test)
	}).Catch(func(t Test, st StackTrace) {
		fmt.Println(t)
		fmt.Println("This is a stack trace result")
		st.Print()
	})

	// Example return error and stack trace
	exception := Try(func(ThrowNewException func(any)) {
		err := errors.New("this is an error exception")
		ThrowNewException(err)
	}).Catch(func(e error) {
		fmt.Println("This is an example for return exception")
	})
	fmt.Println(exception.Error())
	fmt.Println(exception.GetStackTrace())

	// Example success case without exception
	var test Test
	exception = Try(func(ThrowNewException func(any)) {
		test.val = "This is an example for success without error"
	}).Catch(func(e error) {
		fmt.Println("This is an example for return exception")
	})
	fmt.Println(exception.Error())
	fmt.Println(exception.GetStackTrace())
	fmt.Println(test)
}
