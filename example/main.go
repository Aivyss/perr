package main

import (
	"fmt"
	"github.com/aivyss/perr"
	"github.com/aivyss/perr/example/errs"
)

func main() {
	err := errs.GetFirstError()
	if !errs.MyFirstError.Same(err) {
		panic("unexpected result1")
	}
	if errs.MySecondError.Same(err) {
		panic("unexpected result2")
	}
	if !perr.ErrorGroup(errs.MyFirstError, errs.MySecondError).Contains(err) {
		panic("unexpected result3")
	}
	if perr.ErrorGroup(errs.MySecondError).Contains(err) {
		panic("unexpected result4")
	}
	if perr.ErrorGroup().Contains(err) {
		panic("unexpected result5")
	}
	fmt.Println("raw = ", err.Error())

	err = errs.GetSecondError()
	if errs.MyFirstError.Same(err) {
		panic("unexpected result6")
	}
	if !errs.MySecondError.Same(err) {
		panic("unexpected result7")
	}
	if !perr.ErrorGroup(errs.MyFirstError, errs.MySecondError).Contains(err) {
		panic("unexpected result8")
	}
	if perr.ErrorGroup(errs.MyFirstError).Contains(err) {
		panic("unexpected result9")
	}
	if perr.ErrorGroup().Contains(err) {
		panic("unexpected result10")
	}
	fmt.Println("raw = ", err.Error())

	err = errs.GetFirstErrorWithMessage("my-message")
	if !errs.MyFirstError.Same(err) {
		panic("unexpected result11")
	}
	if errs.MySecondError.Same(err) {
		panic("unexpected result12")
	}
	if !perr.ErrorGroup(errs.MyFirstError, errs.MySecondError).Contains(err) {
		panic("unexpected result13")
	}
	if perr.ErrorGroup(errs.MySecondError).Contains(err) {
		panic("unexpected result14")
	}
	if perr.ErrorGroup().Contains(err) {
		panic("unexpected result15")
	}
	fmt.Println("modified = ", err.Error())

	err = errs.GetSecondErrorWithMessage("my-msg")
	if errs.MyFirstError.Same(err) {
		panic("unexpected result16")
	}
	if !errs.MySecondError.Same(err) {
		panic("unexpected result17")
	}
	if !perr.ErrorGroup(errs.MyFirstError, errs.MySecondError).Contains(err) {
		panic("unexpected result18")
	}
	if perr.ErrorGroup(errs.MyFirstError).Contains(err) {
		panic("unexpected result19")
	}
	if perr.ErrorGroup().Contains(err) {
		panic("unexpected result20")
	}
	fmt.Println("modified = ", err.Error())

	fmt.Println("success!")
}
