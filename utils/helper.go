package utils

import (
	"fmt"
	"github.com/cheneylew/go-tools/util"
)

func MyLog() {
	fmt.Println("my log!")
	util.HTTPGet("http://www.baidu.com/")
}
