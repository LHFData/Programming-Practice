package main

//代码通过包来组织，包下有.go源码文件
//类似python模块，或C的库
import (
	"./packageExample"
	"fmt" //导入包，注意，没用到的包引入了的话会过不了编译，go编译是不告警的，所以依赖IDE的静态代码检查
	"math"
	//每个包在解决依赖的前提下，以导入声明的顺序初始化，每个包只会被初始化一次，如果一个包包含了别的包，那么被包含的在包含的初始化过程中即被初始化
	//因此main包最后进行初始化
)

//常量的批量声明
const (
	pi = 3.1415926
	e  = 2.71828
)

//语法方面，go的语句末不需要分号，除非一行有多条
//在编译时，编译器会把
//标识符，整数，浮点数，虚数，字符，字符串文字，关键字，运算符和分隔符这些内容之后的换行符换成分号
//因此，在表达式的变量中间不应该换行，编译换进去的分号会把表达式截断

//推荐驼峰式命名

//函数声明由 函数名 参数列表 可选的返回值列表 包含函数定义的函数体 组成
// func 函数名 输入参数列表 返回参数类型
func main() {
	fmt.Println("Hello,world") //万年不变开头第一句HelloWorld

	//变量声明语法: var 变量名 （类型）|（= 表达式） 有类型了可以不用表达式初始化，有初始化表达式了不用声明类型
	//未初始化变量将被零值初始化，聚合类型的零值是对应各字段的零值，接口或引用类型的零值是nil，字符类型对应的零值是空字符串，布尔类型对应零值是nil

	var l, h = 1, 2
	add(l, h)

	//简短变量声明
	lh := 1

	//指针,可寻址的变量都能够被&取址
	var p = &lh
	fmt.Println("指针：", p, *p, &p, addP(l, lh))
	//new函数，创建一个对应类型的指针变量
	pp := new(int)
	fmt.Println("new的指针", pp)
	//Go的GC通过包级变量和每个当前函数的局部变量的指针或者引用路径，判断是否变量可达，无访问路径即不会影响程序后续计算结果，则可以回收
	//但在函数里，把短生命周期的变量保存到长生命周期的变量内，会导致短生命周期变量从函数内部逃逸，也就是函数返回了之后，该地址空间不会被回收。
	//externalDemo 1
	//局部变量赋值函数
	gcc()
	fmt.Println(global) //这个时候这个x是分配在堆上的，因为必须要在包外可见并基于global生命周期决定是否释放，如果是一般的局部变量是会在栈上的，这由编译器决定
	//元组赋值方式
	l, h, lh = 1, 2, 3
	fmt.Println("元组赋值方式", l, h, lh)
	//多返回值函数的接收，及对不需要值的废弃（使用_接收）
	_, h = multiReturn()
	fmt.Println("多返回值接收", l, h)
	//类型声明语句 type 新类型 底层类型名 用这种方式来分隔不同概念的类型，具有相同底层类型的新类型相互之间也互不兼容
	//类型声明一般出现在包一级，因此若首字母大写则对包外可见
	type stuNum int
	type classNum int
	var sn stuNum = 12
	var cn classNum = 1 //sn直接赋值给cn将无法通过编译
	fmt.Println("类型声明", sn, cn)
	cn = classNum(sn) //但显式的类型转换是可以的，当且仅当底层类型相同时，可以实现这种语义上的转换
	fmt.Println(sn, cn)
	//包
	fmt.Println("包调用", packageExample.Add(l, lh))
	//位运算
	//& AND | OR ^ XOR &^ AND NOT << 左移 >> 右移
	fmt.Println("位运算")
	packageExample.BitOutput(1)
	packageExample.BitOutput(packageExample.BitAND(1, 2))
	packageExample.BitOutput(packageExample.BitOR(1, 2))
	packageExample.BitOutput(packageExample.BitLeftMove(1, 2))
	packageExample.BitOutput(packageExample.BitRightMove(1, 2))
	packageExample.BitOutput(packageExample.BitAndNot(1, 2)) //a中位为1，则结果对应位为0 否则按b对应的位值来
	packageExample.BitOutput(packageExample.BitXOR(1, 2))
	//浮点数
	var f float32 = 16777216 //1<<24 float32有效位为23位，其他位用于指数与符号，当整数大于23位时表示出现误差
	fmt.Println("浮点数精确表示范围", f == f+1)
	//一般来说优先使用float64
	//用下面两个值看两个数据类型的表示最大值
	fmt.Println(math.MaxFloat32, math.MaxFloat64)
	//浮点数表示
	for x := 0; x < 4; x++ {
		y := math.Exp(float64(x))
		fmt.Printf("x=%d e^x=%10.6f %g %e\n", x, y, y, y) //%g打印完整，精度足够，%f字符宽度与精度，%e用科学计数法表示
	}
	//复数
	var cx complex128 = complex(1, 2) //r为实部，i为虚部
	var cy complex128 = complex(3, 4)
	fmt.Println("复数运算", cx*cy, real(cx*cy), imag(cx*cy))
	//字符串
	str := "what the fuck?"
	fmt.Println("字符串", str, str[0]) //超出范围则panic异常，索引获得的是字节值
	fmt.Println(str[5:7])           //子串操作s[i:j]从i到j-1个字节组成新串
	fmt.Println("字符串拼接", "concat"+str)
	//字符串可进行比较，按字符顺序比
	//go标准库字符串处理四大天王 bytes、strings、strconv、unicode 四个包
	//常量测试
	fmt.Println(packageExample.Friday)

	//数据类型，数组
	var arr1 [3]int            //声明
	var arr2 = [...]int{1, 2}  //不定长初始化
	var arr3 = [3]int{1, 2, 3} //定长初始化
	fmt.Println(arr1[1], arr2[1], arr3[1])
	//指定下标的初始化
	arr4 := [...]string{1: "123", 2: "124"}
	fmt.Println(arr4[1])
	arr5 := [...]string{10: "lhf"}
	fmt.Println(len(arr5), arr5[10])

}

//externalDemo 1
var global *int

// 局部变量 全局保存实验
func gcc() {
	var x int
	x = 1
	global = &x
	fmt.Println(&x)
}

//大写字母开头的函数是导出函数，是包外可见的，小写的则是内部函数

func add(a int, b int) int {
	return a + b
}

// go里面同一个包似乎不支持函数重载

func addP(a int, b int) *int {
	var c = a + b
	return &c
}

//多返回值函数
func multiReturn() (int, int) {
	return 88, 99
}
