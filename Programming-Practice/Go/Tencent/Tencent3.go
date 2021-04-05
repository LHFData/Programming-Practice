package main
import (
"strconv"
"strings"
"os"
"fmt"
"bufio"
)

func main(){
	input:=bufio.NewScanner(os.Stdin)
	count:=0

	var length int //测试数据组数
	var phoneNum,chargePersec int
	var phone [][]int
	phoneNo:=0
	for input.Scan(){
		if count==0{
			length,_=strconv.Atoi(strings.Split(input.Text()," ")[0])

			count++
			continue
		}else if count==1{
			phoneNum,_=strconv.Atoi(strings.Split(input.Text()," ")[0])
			chargePersec,_=strconv.Atoi(strings.Split(input.Text()," ")[1])
			for i:=0;i<phoneNum;i++{
				temp:=make([]int,2)
				phone=append(phone,temp)
			}
		}else{
			phone[phoneNo][0],_=strconv.Atoi(strings.Split(input.Text()," ")[0])
			phone[phoneNo][1],_=strconv.Atoi(strings.Split(input.Text()," ")[1])
			phoneNo++
		}
		if phoneNo==len(phone){
			flagForever:=true
			maxDamage:=phone[0][1]
			maxDamageRest:=phone[0][0]
			for i:=0;i<len(phone);i++{
				if phone[i][1]>maxDamage{
					maxDamage=phone[i][1]
					maxDamageRest=phone[i][0]
				}
				if chargePersec>=phone[i][1]{
					continue
				}else{
					flagForever:=false
				}
			}
			if flagForever{
				fmt.Println(-1)
			}else{
				fmt.Println(maxDamageRest/(maxDamage-chargePersec))
			}
			count=0
			phoneNum,chargePersec=0,0
			length=0
			phoneNo=0
			phone=make([][]int,0)
		}
	}

}
