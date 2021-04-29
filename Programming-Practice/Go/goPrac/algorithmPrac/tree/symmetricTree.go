package main

func Mirrors(t *TreeNode){
	t.Right,t.Left=t.Left,t.Right
	if t.Left!=nil{
		Mirrors(t.Left)
	}
	if t.Right!=nil{
		Mirrors(t.Right)
	}
}
func DFS(t *TreeNode,old []int){
	old=append(old, t.Val)
	if t.Left!=nil{
		DFS(t.Left,old)
	}else{
		old=append(old, -1)
	}
	if t.Right!=nil{
		DFS(t.Right,old)
	}else{
		old=append(old, -1)
	}
}
func isSymmetric(root *TreeNode) bool {
	if root==nil{
		return true
	}
	var old []int
	DFS(root,old)
	var new []int
	Mirrors(root)
	DFS(root,new)
	flag:=true
	for i:=0;i<len(old);i++{
		if old[i]!=new[i]{
			flag=false
		}
	}
	return flag
}

func main()  {
	
}
