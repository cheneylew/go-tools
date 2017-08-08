package run

import (
	"github.com/cheneylew/go-tools/util"
	"strings"
	"os"
	"github.com/cyfdecyf/bufio"
	"io"
	"path"
	"fmt"
)

func Run()  {
	dirs := []string{}
	for {
		code := util.InputIntWithMessage("[0]go on;[1]input dir;number:")
		flag := 0
		switch code {
		default:
			flag = 1
			break
		case 1:
			dir := util.InputStringWithMessage("target dir:")
			if len(dir) > 0 {
				dirs = append(dirs, util.Trim(dir))
			}
		}

		if flag == 1 {
			break
		}
	}

	if len(dirs) == 0 {
		util.JJKPrintln("dir is empty! use current dir:",util.ExeDir())
		dirs = append(dirs, util.ExeDir())
	}


	oldStr := ""
	for {
		old := util.InputStringWithMessage("old string:")
		if len(old) > 0 {
			break
		}
	}
	//search(dirs,oldStr)

	newStr := util.InputStringWithMessage("new string:")
	code := util.InputIntWithMessage("[1]replace all;Enter Code:")
	if code != 1 {
		util.JJKPrintln("\nexit!")
		return
	}
	util.JJKPrintln("start replace!")

	for _, mydir := range dirs {
		files := util.FilesAtDir(mydir)
		for _, file := range files {
			ext := path.Ext(file.Path)
			if !strings.Contains(ext,".") {
				continue;
			}

			util.JJKPrintln(fmt.Sprintf("%v %v",oldStr, newStr))

			text := util.FileReadAllString(file.Path)
			newText := strings.Replace(text, oldStr, newStr, -1)
			util.FileWriteString(file.Path, newText)
		}
	}
}

func search(dirs []string, word string)  {
	for _, dir := range dirs {
		files := util.FilesAtDir(dir)
		for _, file := range files {
			ext := path.Ext(file.Path)
			if !strings.Contains(ext,".") {
				continue;
			}
			fi, err := os.Open(file.Path)
			if err != nil {
				util.JJKPrintln(err)
			}
			defer fi.Close()
			reader := bufio.NewReader(fi)
			for {
				line,_,err1 := reader.ReadLine()
				if err1 == io.EOF {
					break
				} else if err1 != nil {
					break
				}
				if len(line) > 0 {
					lineStr := string(line)
					iscontain := strings.Contains(lineStr, word)
					if iscontain {
						util.JJKPrintln(lineStr)
					}
				}

			}
		}
	}
}