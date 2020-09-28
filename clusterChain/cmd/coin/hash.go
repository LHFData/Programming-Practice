package coin

import (
	"bytes"
	"strconv"
)

func(b *Block) SetHash(){
	timestamp:=[]byte(strconv.FormatInt(b.Timestamp,10))
	headers:=bytes.Join([][]byte{b.PrevBlockHash,b.Data,timestamp},[]byte{})
}