// 多肽
package o3

import (
	"fmt"
	"testing"
)

type Foo interface {
	qux()
}
type Bar struct {}
type Baz struct {}
// 非侵入式接口
func (o Bar) qux(){}
func (o Baz) qux(){}
func TestMap(t *testing.T)  {
	var f Foo
	f = Bar{}
	f = Baz{}
	fmt.Println(f)
}