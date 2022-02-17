package log

import (
	"fmt"
	"testing"
)

func TestLoger_Clone(t *testing.T) {
	WithPubKV("test","pub")
	l1 := WithKV("test","l1")
	l2 := WithKV("test","l2")
	l1.WithKV("test2","l1")
	fmt.Println(l1.args)
	fmt.Println(l2.args)
}
