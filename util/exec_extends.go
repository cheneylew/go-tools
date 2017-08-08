package util

import (
	"os/exec"
	"strings"
)

func ExecComandWithDir(dir string,cmd string,params ...string) string  {
	c := exec.Command(cmd,params...)
	c.Dir = dir
	out, _ := c.Output()
	s := string(out)
	return Trim(s)
}

func ExecComand(cmd string,params ...string) string  {
	c := exec.Command(cmd,params...)
	out, _ := c.Output()
	s := string(out)
	s = strings.TrimSuffix(s,string('\n'))
	return s
}

func ExecShell(shell string) string  {
	cmds := strings.Split(shell," ")
	if len(cmds) >= 1 {
		params := cmds[1:len(cmds)]
		for i := 0; i < len(params); i++ {
			params[i] = strings.TrimSpace(params[i])
		}
		return ExecComand(cmds[0],params...)
	}

	return ""
}