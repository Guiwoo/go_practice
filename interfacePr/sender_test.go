package main

import (
	"reflect"
	"testing"
)

func Test_Sender(t *testing.T) {
	var a Sender = &ExpressEmail{
		&Email{"a", "b", "13472 Mortgatan6 Saltjobaden"},
		10000000,
	}

	if v, ok := a.(*Email); ok {
		v.send()
	} else {
		t.Errorf("Test fail Can't type assert from %v\n", reflect.TypeOf(a))
	}
	if v, ok := a.(*ExpressEmail); ok {
		v.superFast()
		v.send()
	}
}
