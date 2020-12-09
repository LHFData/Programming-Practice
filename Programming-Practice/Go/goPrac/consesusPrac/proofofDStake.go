package main

import(
	"fmt"
	"math/rand"
	"time"
	"strconv"
	"crypto/sha256"
	"encoding/hex"
	)

type BlockDPOS struct{
	Index int
	Timestamp string
	Prehash string
	Hash string
	Data []byte
	delegate *NodeDPOS
}
type NodeDPOS struct {
	Name string
	Votes int
}
func GenesisBlock() BlockDPOS{
	gene :=BlockDPOS{0,time.Now().String(),"","",[]byte("创世区块"),nil}
	gene.Hash=string(blockHash(gene))

	var delegate *NodeDPOS=new(NodeDPOS)
	delegate.Name="创世区块"
	delegate.Votes=0
	gene.delegate=delegate

	return gene
}
func blockHash(block BlockDPOS) []byte{
	record :=strconv.Itoa(block.Index)+block.Timestamp+block.Prehash+hex.EncodeToString(block.Data)
	h:=sha256.New()
	h.Write([]byte(record))
	hashed:=h.Sum(nil)
	return hashed
}
func(node *NodeDPOS) GenerateNewBlock(lastBlock BlockDPOS,data []byte) BlockDPOS{
	var newBlock=BlockDPOS{lastBlock.Index+1,time.Now().String(),lastBlock.Hash,"",data,nil}
	newBlock.Hash=hex.EncodeToString(blockHash(newBlock))
	newBlock.delegate=node
	return newBlock
}
var NodeArr=make([]NodeDPOS,10)

func CreateNode()  {
	for i:=0;i<10;i++{
		name:=fmt.Sprintf("Node %d num",i+1)
		NodeArr[i]=NodeDPOS{name,0}
	}

}
func Vote(){
	for i:=0;i<10;i++{
		rand.Seed(time.Now().UnixNano())
		vote:=rand.Intn(10)+1
		NodeArr[i].Votes=vote
	}
}

//投票排序
func SortNodes() []NodeDPOS{
	n:=NodeArr
	for i:=0;i<len(n);i++{
		for j:=0;j<len(n)-1;j++{
			if n[j].Votes<n[j+1].Votes{
				n[j],n[j+1]=n[j+1],n[j]
			}
		}
	}
	return n[:3]
}

func main()  {
	CreateNode()//节点加入网络
	fmt.Println("节点生成")
	fmt.Println(NodeArr)
	fmt.Println("投票决定验证")
	Vote()//投票决定验证
	fmt.Println("前三投票")
	nodes:=SortNodes()//前三进行投票
	fmt.Println(nodes)

	gene:=GenesisBlock()//创世块生成

	lastBlock:=gene//上一块
	for i:=0;i<len(nodes);i++{
		node :=lastBlock.delegate
		fmt.Println("区块号",lastBlock.Index,"节点名称:",node.Name,"节点投票数：",node.Votes)
		lastBlock=nodes[i].GenerateNewBlock(lastBlock,[]byte(fmt.Sprintf("new block %d",i)))
	}
}