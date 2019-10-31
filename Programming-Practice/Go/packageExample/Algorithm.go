package packageExample

func Add(a int, b int) int {
	return a + b
}

//同一个包内，每个文件的包级名字可被包内其他成员直接访问，就跟在同一个文件里一样
func ConvNumToId(id int) No {
	return No(id)
}
