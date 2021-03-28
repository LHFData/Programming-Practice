package output

type Queue struct{
	front int
	rear int
	queue []int
	capacity int
}

func (q Queue) init(cap int) bool{
	q.capacity=cap
	q.queue=make([]int,cap)
	q.front=0
	q.rear=0
	if len(q.queue)>0{
		return true
	}else{
		return false
	}
}

func (q Queue) GetLength() int{
	if q.front<q.rear{

	}
}


