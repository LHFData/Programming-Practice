package main

import "fmt"

func partition(arr []int,i int,j int) int{
	pivot:=arr[i]
	for i<j{
		for i<j&&arr[j]>pivot{
			j--
		}
		if i<j{
			arr[i]=arr[j]
			i++
		}
		for i<j&&arr[i]<pivot{
			i++
		}
		if i<j{
			arr[j]=arr[i]
			j--
		}
	}
	arr[i]=pivot
	return i
}
func QuickSort(arr []int,low int,high int){
	var pivotPos int
	if low<high{
		pivotPos=partition(arr,low,high)
		QuickSort(arr,low,pivotPos-1)
		QuickSort(arr,pivotPos+1,high)
	}
}
func partition(arr []int,low int,high int)int{
	pivot:=arr[low]
	for low<high{
		for low<high&&arr[high]>pivot{
			high--
		}
		if low<high{
			arr[low]=arr[high]
			low++
		}
		for low<high&&arr[low]<pivot{
			low++
		}
		if low<high{
			arr[high]=arr[low]
			high--
		}
	}
	arr[low]=pivot
	return low
}

func QuickSort(arr []int,low int,high int){
	var pivotPos int
	if low<high{
		pivotPos=partition(arr,low,high)
		QuickSort(arr,low,pivotPos-1)
		QuickSort(arr,pivotPos+1,high)
	}
}



//func partition(array []int,i int,j int) int{
//	pivot:=array[i]
//	for i<j{
//		for j>i && array[j]>pivot{
//			j--
//		}
//		if j>i{
//			array[i]=array[j]
//			i++
//		}
//		for i<j&&array[i]<pivot{
//			i++
//		}
//		if i<j{
//			array[j]=array[i]
//			j--
//		}
//	}
//	array[i]=pivot
//	return i
//}
//func QuickSort(array []int,low int,high int){
//	var pivotPos int
//	if low <high{
//		pivotPos=partition(array,low,high)
//		QuickSort(array,low,pivotPos-1)
//		QuickSort(array,pivotPos+1,high)
//	}
//}
func main(){
	a:=[]int{1,3,5,2,4,6}
	QuickSort(a,0,5)
	fmt.Println(a)
}