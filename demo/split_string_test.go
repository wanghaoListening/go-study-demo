package main

import (
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) {

	got := Split("abc,re,ew,q1", ",")
	want := []string{"bc", "re", "ew", "q1"}

	if !reflect.DeepEqual(got, want) {

		t.Errorf("want:%v,got:%v", want, got)
	}
}
