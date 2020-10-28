package main

import (
	"time"
	"strconv"
	"crypto/sha256"
	"math/rand"
	"fmt"
	"encoding/hex"
)

type Node struct {
	Tokens int
	Days int
	Address string
}
type Block struct {
	Index int
	Data string
	PreHash string
	Hash string
	Timestamp string
	Validator *Node
}
func genesisBlock() Block{
	var genesBlock=Block{0,"我是你爹","","",time.Now().String(),&Node{0,0,"dd"}}
	genesBlock.Hash=hex.EncodeToString(BlockHash(&genesBlock))
	return genesBlock
}
func BlockHash(block *Block) []byte{
	record:=strconv.Itoa(block.Index)+block.Data+block.PreHash+block.Timestamp+block.Validator.Address
	h:=sha256.New()
	h.Write([]byte(record))
	hashed:=h.Sum(nil)
	return hashed
}
var nodes=make([]Node,5)
var addr=make([]*Node,15)
func InitNodes(){
	for i:=0;i<5;i++{
		nodes[i]=Node{i,1,"0x1234"+strconv.Itoa(i)}
		fmt.Println(nodes[i])
		//这里有个坑，打印就不报内存错，不打印就报内存错。为什么
	}
	cnt:=0
	for i:=0;i<5;i++{
		//n个节点，每个节点按代币数*持有时长=地址池比重*地址池长度
		for j:=0;j<nodes[i].Tokens*nodes[i].Days;j++{
			addr[cnt]=&nodes[i]
			cnt++
		}
	}
}

func CreateNewBlock(lastBlock *Block,data string) Block{
	var newBlock Block
	newBlock.Index=lastBlock.Index+1
	newBlock.Timestamp=time.Now().String()
	newBlock.PreHash=lastBlock.Hash
	newBlock.Data=data
	//随机数生成
	rand.Seed(time.Now().Unix())
	var rd=rand.Intn(15)

	//选矿工，被选中成为矿工的几率与前面算的权益比重成正比
	node:=addr[rd]
	newBlock.Validator=node//当前区块验证者
	//简单挖矿奖励，具体用的什么还需要看
	node.Tokens+=1
	//算出本块哈希
	newBlock.Hash=hex.EncodeToString(BlockHash(&newBlock))
	//刨去激励过程，POS出块十分简单，省略计算挖矿的时间，出块很快
	return newBlock
}
func main(){
	InitNodes()

	var genesisBlock=genesisBlock()
	var newBlock=CreateNewBlock(&genesisBlock,"我是儿")

	fmt.Println(newBlock)
	fmt.Println(newBlock.Validator.Address)
	fmt.Println(newBlock.Validator.Tokens)
}