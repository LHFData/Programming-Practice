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


package main
import (
"fmt"
"strconv"
"strings"
"bufio"
"os"
)
//写这玩意儿是因为用Go刷牛客的人特别少，很多题解都没看到go语言的
//然后面对机试的ACM核心模式，足以让常年用Go刷LC的人很不适应
func main(){
	input:=bufio.NewScanner(os.Stdin)
	for input.Scan(){
		temp,_:=strconv.Atoi(strings.Split(input.Text()," ")[0])
		//下标为字符串按分隔符分割后索引的字符位置
		//每次循环读入一行，并按行将一行数据按分隔进行读入，十分方便
		//可能会带来一些额外的开销，但作为一种通用的方式也是懒人福音了
		//后续有新的效率较高的方式进行输入输出处理，再更新

	}
}
————————————————
版权声明：本文为CSDN博主「西土城想念伦敦」的原创文章，遵循CC 4.0 BY-SA版权协议，转载请附上原文出处链接及本声明。
原文链接：https://blog.csdn.net/LHF_debug/article/details/115443052