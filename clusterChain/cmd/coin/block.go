package coin

type Block struct{
	Timestamp int64
	Data []byte
	PrevBlockHash []byte
	hash []byte
}

