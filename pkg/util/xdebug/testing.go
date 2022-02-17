package xdebug

import (
	"flag"
	"sync"
)

var(
	isTestingMode     bool
)

// IsTestingMode 判断是否在测试模式下
var onceTest = sync.Once{}

func IsTestingMode() bool {
	onceTest.Do(func() {
		isTestingMode = flag.Lookup("test.v") != nil
	})

	return isTestingMode
}

