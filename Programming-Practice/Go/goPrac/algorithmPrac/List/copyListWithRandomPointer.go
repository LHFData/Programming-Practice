package main

type Node struct{
	Val int
	Next *Node
	Random *Node
}
func copyRandomList(head *Node) *Node {


	p:=head
	if head==nil{
		return nil
	}

	var res *Node
	var tail *Node
	res=new(Node)
	res.Val=head.Val
	tail=res
	for p!=nil{
		temp:=new(Node)
		temp.Val=p.Val
		tail.Next=temp
		tail=temp
		p=p.Next
	}
	//新链表就绪
	p=head
	pp:=res
	m:=make(map[*Node]int)
	r:=make(map[int]*Node)
	count:=0
	for p!=nil&&pp!=nil{
		m[p]=count
		r[count]=pp
		count++
		pp=pp.Next
		p=p.Next
	}
	pp=res
	p=head

	for p!=nil&&pp!=nil{
		if p.Random==nil{
			pp.Random=nil
		}else{
			index,_:=m[p]
			pp.Random=r[index]
		}
		p=p.Next
		pp=pp.Next
	}
	return res
}
