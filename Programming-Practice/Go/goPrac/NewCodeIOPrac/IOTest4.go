package main
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){
	var length int
	input:=bufio.NewScanner(os.Stdin)
	for input.Scan(){
		a:=0
		length,_=strconv.Atoi(strings.Split(input.Text()," ")[0])
		if length==0{
			return
		}
		for i:=1;i<length+1;i++{
			temp,_:=strconv.Atoi(strings.Split(input.Text()," ")[i])
			a=a+temp
		}
		fmt.Println(a)
	}
}
