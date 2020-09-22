package core
import(
	"time"
)
type Block struct{
	Timestamp int64
	Data []byte
	PrevBlockHash []byte
	Hash []byte
	Nonce int
}
func Newblocks(Data string,PrevBlockHash []byte) *Block{
	block:=&Block{time.Now().Unix(),[]byte(Data),PrevBlockHash,[]byte{},0} //新区块基于时间戳，新数据与前置区块Hash生成
	pow:= NewProofofWork(block)
	n,h :=pow.Run()
	block.Hash=h
	block.Nonce=n
	return block
}
//新创世区块
func NewGenesisBlock() (*Block){
	return Newblocks("Genesis Block",[]byte{})
}