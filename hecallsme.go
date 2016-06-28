package main
import (
"fmt"
"os"
"io/ioutil"
"strconv"
)

const CRLF = "\r\n"

func main(){
content := ""

for _,val := range os.Environ(){
content += val + CRLF
}
content += "======" + CRLF
for n, val := range os.Args {
content += "Arg"+ strconv.Itoa(n)+ ": "+ val + CRLF
}
fmt.Println(content)
ioutil.WriteFile("C:/temp/howcallsme.txt",[]byte(content),os.ModePerm)
}
