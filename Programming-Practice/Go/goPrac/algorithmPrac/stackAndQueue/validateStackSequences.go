package main

func validateStackSequences(pushed []int, popped []int) bool {
	temp:=[]int{}
	push:=0
	pop:=0
	for pop<len(popped){

		if push>=len(pushed){
			if temp[len(temp)-1]==popped[pop]{
				temp=temp[:len(temp)-1]

				push++
			}else{
				return false
			}
		}else{
			if pushed[push]!=popped[pop]{
				if len(temp)==0{
					temp=append(temp,pushed[push])
					push++
				}else{
					if temp[len(temp)-1]==popped[pop]{
						temp=temp[:len(temp)-1]
						pop++
						push++
					}else{
						temp=append(temp,pushed[push])
						push++
					}
				}

			}else if pushed[push]==popped[pop]{
				push++
				pop++
			}
		}
	}

	return true
}