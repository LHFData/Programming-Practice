package packageExample

type No int
type Id int
type Name string

//iota常量生成器，批量初始化常量，它会按照顺序依次将iota从0开始对后续声明的常量进行赋值
type Weekday int

const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

var LHF = 1

//初始化函数，不可被调用或引用，每个文件可包含多个init初始化函数
func init() {
	LHF = 2
}
