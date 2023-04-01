package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	"strconv"
	"time"
)

func main() {

	dtstr := time.Now().Format("2006-01-02 15:04:05 Z0700")
	content := ""
	filedest := "C:/temp/hecallsme.txt"
	CRLF := "\r\n"
	if runtime.GOOS == "windows" {
	} else {
		filedest = "/tmp/hecallsme.txt"
		CRLF = "\n"
	}
	for _, val := range os.Environ() {
		content += val + CRLF
	}
	content += "==== " + dtstr + " ====" + CRLF
	for n, val := range os.Args {
		content += "Arg" + strconv.Itoa(n) + ": " + val + CRLF
	}
	fmt.Println(content)
	ioutil.WriteFile(filedest, []byte(content), os.ModePerm)
}
