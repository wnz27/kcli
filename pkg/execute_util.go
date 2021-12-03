/**
 * @project ads-serving
 * @Author 27
 * @Description //TODO
 * @Date 2021/11/20 11:24 11月
 **/
package pkg

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
)

// LogFatalChangeDirForRun 执行cmd 把cmd执行的命令的输出定向到标准输出 错误输出到标准错误,
// 这个适用于执行的命令后续有持续输出
func LogFatalChangeDirForRun(cmd *exec.Cmd, execPath string) {
	cmd.Dir = execPath
	var stdoutBuf, stderrBuf bytes.Buffer
	cmd.Stdout = io.MultiWriter(os.Stdout, &stdoutBuf)
	cmd.Stderr = io.MultiWriter(os.Stderr, &stderrBuf)

	err := cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	outStr, errStr := stdoutBuf.String(), stderrBuf.String()
	fmt.Printf("\nout:\n%s\nerr:\n%s\n", outStr, errStr)
}

// LogFatal cmd 有问题直接报错
func LogFatal(cmd *exec.Cmd) {
	stdoutStderr, err := cmd.Output()

	if err != nil {
		log.Fatal(err)
	}
	if len(stdoutStderr) != 0 {
		fmt.Println(stdoutStderr)
	}
}
