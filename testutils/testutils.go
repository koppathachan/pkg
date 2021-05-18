package testutils

import (
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
)

//Asserter returns an assert(condition bool, msg string, v ...interface{}) function
//assert fails the test if the condition is false.
func Asserter(tb testing.TB) func(condition bool, msg string, v ...interface{}) {
	assert := func(condition bool, msg string, v ...interface{}) {
		if !condition {
			_, file, line, _ := runtime.Caller(1)
			fmt.Printf("\033[31m%s:%d: "+msg+"\033[39m\n\n", append([]interface{}{filepath.Base(file), line}, v...)...)
			tb.Fail()
		}
	}
	return assert
}

//Okayer returns an ok(err error) function
//ok fails the test if an err is not nil
func Okayer(tb testing.TB) func(err error) {
	ok := func(err error) {
		if err != nil {
			_, file, line, _ := runtime.Caller(1)
			fmt.Printf("\033[31m%s:%d: unexpected error: %s\033[39m\n\n", filepath.Base(file), line, err.Error())
			tb.Fail()
		}
	}
	return ok
}

//NotOkayer returns a notOk(err error) function
//notOk fails the test if an err is nil
func NotOkayer(tb testing.TB) func(err error) {
	notOk := func(err error) {
		if err == nil {
			_, file, line, _ := runtime.Caller(1)
			fmt.Printf("\033[31m%s:%d: expected error: %s\033[39m\n\n", filepath.Base(file), line, "")
			tb.Fail()
		}
	}
	return notOk
}

//Equater returns and equals(exp, act interface{}) function
//equals fails the test if exp is not equal to act.
func Equater(tb testing.TB) func(exp, act interface{}) {
	equals := func(exp, act interface{}) {
		if !reflect.DeepEqual(exp, act) {
			_, file, line, _ := runtime.Caller(1)
			fmt.Printf("\033[31m%s:%d:\n\n\texp: %#v\n\n\tgot: %#v\033[39m\n\n", filepath.Base(file), line, exp, act)
			tb.Fail()
		}
	}
	return equals
}
