// 继承
package o2

import (
	"fmt"
	"testing"
)
type Foo struct {
	baz string
}
type Bar struct {
	Foo
}
func(f *Foo) echo(){
	fmt.Print(f.baz)
}
func TestObjOne(t *testing.T)  {
	a := Bar{Foo{baz: "test"}}
	a.echo()
}

