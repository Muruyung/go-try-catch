package main

import (
	"errors"
	"fmt"

	gotry "github.com/Muruyung/go-try-catch"
)

type Test struct {
	val interface{}
}

func main() {
	// Example exception using int value
	gotry.Try(func(ThrowNewException func(any)) {
		ThrowNewException(1)
	}).Catch(func(i int) {
		fmt.Println("This is an int exception:", i)
	})

	// Example exception using string value
	gotry.Try(func(ThrowNewException func(any)) {
		ThrowNewException("hello world")
	}).Catch(func(s string) {
		fmt.Println("This is a string exception:", s)
	})

	// Example exception using error value
	gotry.Try(func(ThrowNewException func(any)) {
		err := errors.New("this is an error exception")
		ThrowNewException(err)
	}).Catch(func(e error) {
		fmt.Println(e)
	})

	// Example exception using struct data
	gotry.Try(func(ThrowNewException func(any)) {
		test := Test{
			val: "This is a struct exception",
		}
		ThrowNewException(test)
	}).Catch(func(t Test, st gotry.StackTrace) {
		fmt.Println(t)
		fmt.Println("This is a stack trace result")
		st.Print()
	})

	// Example return error and stack trace
	exception := gotry.Try(func(ThrowNewException func(any)) {
		err := errors.New("this is an error exception")
		ThrowNewException(err)
	}).Catch(func(e error) {
		fmt.Println("This is an example for return exception")
	})
	fmt.Println(exception.Error())
	fmt.Println(exception.GetStackTrace())

	// Example success case without exception
	var test Test
	exception = gotry.Try(func(ThrowNewException func(any)) {
		test.val = "This is an example for success without error"
	}).Catch(func(e error) {
		fmt.Println("This is an example for return exception")
	})
	fmt.Println(exception.Error())
	fmt.Println(exception.GetStackTrace())
	fmt.Println(test)

	// Example using block
	exception = gotry.Block{
		Try: func(ThrowNewException func(any)) {
			err := errors.New("this is an error using block exception")
			ThrowNewException(err)
		},
		Catch: func(e error, st gotry.StackTrace) {
			fmt.Println(e)
			for _, val := range st.GetList() {
				fmt.Print(val)
			}
		},
	}.Do()
	fmt.Println(exception.Error())
	fmt.Println(exception.GetStackTrace())
}
