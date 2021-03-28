package main

import (
	"errors"
	fmt "fmt"
)
type doubleLinkListNode struct {
	Val int
	prev *doubleLinkListNode
	next *doubleLinkListNode
}

type linkQueue struct {
	head *doubleLinkListNode
	tail *doubleLinkListNode
	length int
}

func (q *linkQueue) GetFront() (int,error){
	if q.length==0{
		return 0,errors.New("Empty queue,Get fail")
	}
	return q.head.Val,nil
}
func (q *linkQueue) GetRear() (int,error){
	if q.length==0{
		return 0,errors.New("Empty queue,Get fail")
	}
	return q.tail.Val,nil
}
func (q *linkQueue) EnqueueRear(n int)  {
	var temp doubleLinkListNode
	if q.tail!=nil{
		temp=doubleLinkListNode{n,q.tail,nil}
	}else{
		temp=doubleLinkListNode{n,nil,nil}
	}
	q.tail=&temp
}
func (q *linkQueue) EnqueueFront(n int)  {
	var temp doubleLinkListNode
	if q.head!=nil {
		temp = doubleLinkListNode{n, nil, q.head}

	}else{
		temp = doubleLinkListNode{n, nil, nil}
	}
	q.length++
	q.head = &temp
}
func (q *linkQueue) DequeueFront()(int,error){
	n := q.head.Val
	if q.head!=nil {

		q.head = q.head.next
		return n,nil
	}
	return 0,errors.New("Empty Queue")
}
func (q *linkQueue) DequeueRear()(int,error){
	if q.tail!=nil {
		n := q.tail.Val
		if q.tail.prev != nil {
			q.tail.prev.next = nil
			q.tail = q.tail.prev
			return n, nil
		} else {
			q.tail = nil
			return n, nil
		}
	}
	return 0,errors.New("Empty Queue")

}
func maxSlidingWindow(nums []int, k int) []int {
	que:=make([]int,len(nums)-k+2)
	result:=make([]int,k)
	length:=0
	head:=0
	tail:=0
	queFront:=0 //指向队头
	queRear:=0 //指向末端空位
	queLength:=0
	for i:=0;i<len(nums);i++{
		//初始化过程开始
		if length<k-1{
			length++
			if tail>0{
				if que[queRear-1]>nums[tail]{//开始滑动窗口位置
					que[queRear]=nums[tail]//若下一位小于队尾元素，乖乖入队
					queLength++
					queRear++//队尾开始移动
				}else{
					que[queRear-1]=nums[tail]//若下一位比队尾要大，直接覆盖掉前值，等于双端中的前值出队后值入队
					//这种情况下队伍长度并无变化
				}
			}else{//初始位置，直接入队
				que[queRear]=nums[tail]
				queLength++
				queRear++
			}
			tail++
		}else{//初始化过程结束
			if que[queRear-1] > nums[tail]{
				que[queRear] = nums[tail]
				queLength++
				queRear++
			} else {
				que[queRear-1] = nums[tail] //直接覆盖前值，无长度变化
			}
			//出队时间有问题
			result[head]=que[queFront]//head的值是对的，直接重复用了
			queLength--
			queFront++
			//窗口滑动
			head++
			tail++
		}
	}
	return result
}
func main(){
	n:=[]int{2,3,1,4,5,6,1,3,5}
	fmt.Println(maxSlidingWindow(n,3))
}