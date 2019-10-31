package packageExample

type No int
type Id int
type Name string

var LHF = 1

//初始化函数，不可被调用或引用，每个文件可包含多个init初始化函数
func init() {
	LHF = 2
}
