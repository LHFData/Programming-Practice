package main

import (
	"bytes"
	"fmt"
)

func main(){
	s:="Hello 世界"
	for _,c:= range s{
		buf:=bytes.NewBuffer([]byte{})
		buf.writeRune(c)

	}
}
