package core

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"log"
	"math"
	"math/big"
)
var MaxNonce=math.MaxInt64
const Targetbits=20
//Pow 由新区块与目标值生成
type ProofofWork struct {
	Block *Block
	Target *big.Int
}
func NewProofofWork(b *Block) *ProofofWork{
	target :=big.NewInt(1)
	//Lsh位置敏感哈希
	target.Lsh(target,uint(256-Targetbits))
	pow :=&ProofofWork{b,target}
	return pow
}
func (p *ProofofWork)prepareDate(n int) []byte{
	data :=bytes.Join([][]byte{
		p.Block.PrevBlockHash,
		p.Block.Data,
		IntToHex(p.Block.Timestamp),
		IntToHex(int64(Targetbits)),
		IntToHex(int64(n)),
	},[]byte{},)
	return data
}
func (p *ProofofWork) Run() (int,[]byte){
	var Hashbig big.Int
	var hash [32]byte
	var nonce=0
	for nonce <MaxNonce{
		data :=p.prepareDate(nonce)
		hash=sha256.Sum256(data)
		Hashbig.SetBytes(hash[:])
		if Hashbig.Cmp(p.Target)==-1{
			break
		} else{
			nonce++
		}
	}
	return nonce,hash[:]
}
func (p *ProofofWork) Isvalidate() bool{
	var Hashbig big.Int
	data :=p.prepareDate(p.Block.Nonce)
	hash := sha256.Sum256(data)
	Hashbig.SetBytes(hash[:])
	is := Hashbig.Cmp(p.Target)==-1
	return is
}
func IntToHex(num int64) []byte{
	buff:=new(bytes.Buffer)
	err:=binary.Write(buff,binary.BigEndian,num)
	if err !=nil{
		log.Panic(err)
	}
	return buff.Bytes()
}