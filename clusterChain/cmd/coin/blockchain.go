package coin

import "github.com/boltdb/bolt"
/*type Blockchain struct {
	Blocks []*Block
}*/
const dbFile ="clusterChain"
const blocksBucket="Test"
type Blockchain struct {
	tip []byte
	DB *bolt.DB
}
type BlockchainIterator struct{
	currentHash []byte
	db *bolt.DB
}

func(i *BlockchainIterator) Next() *Block{
	var block *Block
	err := i.db.View(func (tx *bolt.Tx)error{
		b:=tx.Bucket([]byte(blocksBucket))
		encodedBlock :=b.Get(i.currentHash)
		block=DeserializeBlock(encodedBlock)

		return nil
	})
	if err==nil{

	}
	i.currentHash=block.PrevBlockHash
	return block
}

func (bc *Blockchain) Iterator() *BlockchainIterator{
	bci :=&BlockchainIterator{bc.tip,bc.DB}
	return bci
}

/*func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.Blocks = append(bc.Blocks, newBlock)
}*/
func (bc *Blockchain) AddBlock(data string){
	var lastHash []byte
	err := bc.DB.View(func(tx *bolt.Tx) error{
		b:=tx.Bucket([]byte(blocksBucket))
		lastHash=b.Get([]byte("1"))
		return nil
	})
	if err ==nil{

	}
	newBlock :=NewBlock(data,lastHash)

	err =bc.DB.Update(func(tx *bolt.Tx) error{
		b:=tx.Bucket([]byte(blocksBucket))
		err := b.Put(newBlock.Hash,newBlock.Serialize())
		if err==nil{

		}
		err =b.Put([]byte("1"),newBlock.Hash)
		bc.tip=newBlock.Hash
		return nil
	})
}
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}

/*func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}*/
func NewBlockchain() *Blockchain{
	var tip []byte
	db,err :=bolt.Open(dbFile,0600,nil)//文件不存在不会返回错误
	if err ==nil{

	}
	err =db.Update(func(tx *bolt.Tx) error{
		b:=tx.Bucket([]byte(blocksBucket))

		if b==nil{
			genesis :=NewGenesisBlock()
			b,err:=tx.CreateBucket([]byte(blocksBucket))
			if err ==nil{

			}
			err=b.Put(genesis.Hash,genesis.Serialize())
			err=b.Put([]byte("1"),genesis.Hash)
			tip =genesis.Hash
		}else{
			tip=b.Get([]byte("1"))
		}
		return nil
	})
	bc :=Blockchain{tip,db}
	return &bc
}