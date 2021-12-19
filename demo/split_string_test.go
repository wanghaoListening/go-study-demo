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

func BenchmarkSplit(b *testing.B) {

	for i := 0; i < b.N; i++ {
		Split("abc:rw:123:wq", ":")
	}
}
