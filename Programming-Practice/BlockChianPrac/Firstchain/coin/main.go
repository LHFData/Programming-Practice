package main

import (
	"../core"
	"fmt"
)
func main(){
	bc:=core.NewBlockChain()
	bc.AddBlock("send 1 to wangying")
	bc.AddBlock("send 12 to lhf")
	for _,block := range bc.Blocks{
		fmt.Printf("Prev Hash:%x\n",block.PrevBlockHash)
		fmt.Printf("Data:%s\n",block.Data)
		fmt.Printf("Hash:%x\n",block.Hash)
		pow:=core.NewProofofWork(block)
		fmt.Printf("pow:%t\n",pow.Isvalidate())
	}
}