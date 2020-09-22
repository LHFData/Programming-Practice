package core
type BlockChain struct{
	Blocks []*Block
}
//添加区块链
func (b *BlockChain) AddBlock(data string){
	preB :=b.Blocks[len(b.Blocks)-1]
	newB :=Newblocks(data,preB.Hash)
	b.Blocks=append(b.Blocks,newB)
}
//新区块链生成，第一个应当是创世区块
func NewBlockChain() *BlockChain{
	Newb:=NewGenesisBlock()
	a:=[]*Block{Newb}
	return &BlockChain{a}
}
