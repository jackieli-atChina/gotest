package o1

import (
	"fmt"
	"testing"
)
// 面向对象：封装
type Foo struct {
	baz string
}

func (f *Foo)  echo(){
	fmt.Println(f.baz)
}

func TestM(t *testing.T) {
	a := Foo{baz: "test"}
	a.echo()
}