package run

import (
	"github.com/cheneylew/go-tools/util"
	"strings"
	"os"
	"io"
	"path"
	"bufio"
	"github.com/fatih/color"
	"fmt"
	"flag"
	"regexp"
)

var ReplaceFound = false
var Exts []string = []string{}

func Run() {
	f := ""
	fr := ""
	r := ""
	ext := ""
	flag.StringVar(&f, "f", "", "word search")
	flag.StringVar(&fr, "fr", "", `regexp search:cmd -fr '\d+'`)
	flag.StringVar(&r, "r", "", `replace key work`)
	flag.StringVar(&ext, "ext", "", `file extension:'.go|.h|.m'`)
	flag.Parse()

	if len(ext) > 0 {
		Exts = strings.Split(ext,"|")
	}

	if len(f) > 0 {
		search([]string{util.ExeDir()}, f)
	} else if len(fr) > 0 {
		searchRegexp([]string{util.ExeDir()}, fr)
	} else if len(r) > 0 {
		dirs := []string{util.ExeDir()}
		search(dirs, r)
		if !ReplaceFound {
			util.JJKPrintln("replace not found!")
			return
		}
		newStr := util.InputStringWithMessage("\nnew string:")
		replace(dirs,r,newStr)
	} else {
		flag.Usage()
	}
}

func replace(dirtmp []string, old string, new string) {
	dirs := dirtmp
	oldStr := old
	newStr := new
	code := util.InputIntWithMessage("[1]confirm;\nEnter Code 1:")
	if code != 1 {
		util.JJKPrintln("\nexit!")
		return
	}

	for _, dir := range dirs {
		files := util.FilesAtDir(dir)
		for _, file := range files {
			//指定后缀文件
			if len(Exts) != 0 {
				okType := false
				for _, value := range Exts {
					if file.Ext == value {
						okType = true
					}
				}
				if !okType {
					continue
				}
			}

			//隐藏目录忽略
			pathPart := strings.Split(file.Dir,"/")
			last := ""
			for _, value := range pathPart {
				if len(value) > 0 {
					last = value
				}
			}
			if strings.HasPrefix(last,".") {
				continue;
			}

			ext := path.Ext(file.Path)
			if !strings.Contains(ext, ".") {
				continue;
			}

			text := util.FileReadAllString(file.Path)
			if strings.Contains(text, oldStr) {
				util.JJKPrintln("write file",file.Path)
				newText := strings.Replace(text, oldStr, newStr, -1)
				util.FileWriteString(file.Path, newText)
			}
		}
	}
}

func search(dirs []string, word string) {
	util.JJKPrintln("\n=========================== search ===========================")
	for _, dir := range dirs {
		files := util.FilesAtDir(dir)
		for _, file := range files {
			//指定后缀文件
			if len(Exts) != 0 {
				okType := false
				for _, value := range Exts {
					if file.Ext == value {
						okType = true
					}
				}
				if !okType {
					continue
				}
			}

			//隐藏目录忽略
			pathPart := strings.Split(file.Dir,"/")
			last := ""
			for _, value := range pathPart {
				if len(value) > 0 {
					last = value
				}
			}
			if strings.HasPrefix(last,".") {
				continue;
			}

			//后缀
			ext := path.Ext(file.Path)
			if !strings.Contains(ext, ".") {
				continue;
			}
			fi, err := os.Open(file.Path)
			if err != nil {
				util.JJKPrintln(err)
			}
			defer fi.Close()
			reader := bufio.NewReader(fi)

			lineNum := 0
			flag := 0
			for {
				lineNum += 1
				line, _, err1 := reader.ReadLine()
				if err1 == io.EOF {
					break
				} else if err1 != nil {
					util.JJKPrintln(err1)
					break
				}
				if len(line) > 0 {
					lineStr := string(line)
					iscontain := strings.Contains(lineStr, word)
					if iscontain {
						ReplaceFound = true

						if flag == 0 {
							fmt.Printf("\nfile: %v\n", file.Path)
						}
						flag = 1;

						charCount := 20
						arr := make([]string, 0)
						strs := strings.Split(lineStr, word)
						for _, str := range strs {
							str = util.Trim(str)
							strLength := len(str)
							if strLength < charCount {
								arr = append(arr, str)
							} else {
								chars := []byte(str)
								arr = append(arr, string(chars[(strLength - charCount):]))
							}
						}

						//高亮关键字
						yellow := color.New(color.FgYellow).SprintFunc()
						red := color.New(color.FgRed).SprintFunc()
						fmt.Printf("line:%v:", red(lineNum))
						for key, str := range arr {
							if key == len(arr) - 1 {
								fmt.Printf("%s", str)
							} else {
								fmt.Printf("%s%s", str, yellow(word))
							}
						}
						fmt.Println("")
					}
				}

			}
		}
	}
}

func searchRegexp(dirs []string, regStr string) {
	util.JJKPrintln("\n=========================== search regexp ===========================")
	for _, dir := range dirs {
		files := util.FilesAtDir(dir)
		for _, file := range files {
			//指定后缀文件
			if len(Exts) != 0 {
				okType := false
				for _, value := range Exts {
					if file.Ext == value {
						okType = true
					}
				}
				if !okType {
					continue
				}
			}

			//隐藏目录忽略
			pathPart := strings.Split(file.Dir,"/")
			last := ""
			for _, value := range pathPart {
				if len(value) > 0 {
					last = value
				}
			}
			if strings.HasPrefix(last,".") {
				continue;
			}

			ext := path.Ext(file.Path)
			if !strings.Contains(ext, ".") {
				continue;
			}
			if strings.Contains(ext, "DS_Store") {
				continue;
			}

			fi, err := os.Open(file.Path)
			if err != nil {
				util.JJKPrintln(err)
			}
			defer fi.Close()
			reader := bufio.NewReader(fi)

			lineNum := 0
			flag := 0
			for {
				lineNum += 1
				line, _, err1 := reader.ReadLine()
				if err1 == io.EOF {
					break
				} else if err1 != nil {
					util.JJKPrintln(err1)
					break
				}
				if len(line) > 0 {
					lineStr := string(line)
					result := util.JKRegFindAll(lineStr, regStr)
					iscontain := len(result) > 0
					if iscontain {
						if flag == 0 {
							fmt.Printf("\nfile: %v\n", file.Path)
						}
						flag = 1;

						//劈开一行
						matchStr := strings.Join(result, "|")
						util.JJKPrintln(matchStr)
						re := regexp.MustCompile(matchStr)
						result := re.Split(lineStr, -1)

						//关键字前后的文字
						charCount := 20
						arr := make([]string, 0)
						strs := result
						for _, str := range strs {
							str = util.Trim(str)
							strLength := len(str)
							if strLength < charCount {
								arr = append(arr, str)
							} else {
								chars := []byte(str)
								arr = append(arr, string(chars[(strLength - charCount):]))
							}
						}

						//高亮关键字
						yellow := color.New(color.FgYellow).SprintFunc()
						red := color.New(color.FgRed).SprintFunc()
						fmt.Printf("line:%v:", red(lineNum))
						for key, str := range arr {
							if key == len(arr) - 1 {
								fmt.Printf("%s", str)
							} else {
								fmt.Printf("%s%s", str, yellow(matchStr))
							}
						}
						fmt.Println("")
					}
				}

			}
		}
	}
}