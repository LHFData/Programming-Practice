package main

import "fmt"

func isNumber(s string) bool {

	var state string
	state="spaceStart"
	//spaceStart,PN,start,point,pointEnd,spaceEnd,E,EPN,Eend,
	start:=0
	end:=len(s)-1
	for i:=0;i<len(s);i++{
		if s[i]==' '{
			start++
		}else{
			break
		}
	}
	for i:=len(s)-1;i>=0;i--{
		if s[i]==' '{
			end--
		}else{
			break
		}
	}
	for i:=start;i<=end;i++{

		switch s[i]{
		case '.':{
			if state=="start"||state=="spaceStart"{
				state="point"
				if start==end{
					return false
				}
			}else{
				return false
			}
		}
		case '+':{
			if state=="spaceStart"{
				state="start"
				if start==end{
					return false
				}
			}else if state=="E"{
				state="PN"
			}else{
				return false
			}
		}
		case '-':{
			if state=="spaceStart"{
				state="PN"
				if start==end{
					return false
				}
			}else if state=="E"{
				state="EPN"
			}else{
				return false
			}
		}
		case 'E':{
			if state=="start"||state=="pointEnd"{
				state="E"
			}else{
				return false
			}
		}
		case 'e':{
			if state=="start"||state=="pointEnd"{
				state="E"
			}else{
				return false
			}
		}
		case ' ':{
			return false
		}
		case '0','1','2','3','4','5','6','7','8','9':{

			if state=="E"{
				state="Eend"
			}else if state=="spaceStart"{
				state="start"
			}else if state=="PN"{
				state="start"
			}else if state=="EPN"{
				state="Eend"
			}else if state=="point"{
				state="pointEnd"
			}else if state=="pointEnd"{

			}else if state=="start"{

			}else if state=="Eend"{

			}else{
				return false
			}
		}
		default:
			return false
		}
	}
	if state=="spaceStart"||state=="E"||state=="EPN"||state=="PN"{
		return false
	}

	return true
}
func main(){
	fmt.Println(isNumber(" . "))
}
