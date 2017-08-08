package main

import (
	"fmt"
	"github.com/cheneylew/go-tools/util"
)

func main() {
	fmt.Println("hello world2")
	util.HTTPGet("http://www.baidu.com/")
}
