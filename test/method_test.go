package main

import "testing"

type T struct {
	a int
}

func (t T) Get() int {
	return t.a
}

func (t *T) Set(a int) {
	t.a = a
}


func TestStruct(t *testing.T) {
	var tt T
	T.Get(tt)
	(*T).Set(&tt, 1)
	
	t.Logf("success")
}
