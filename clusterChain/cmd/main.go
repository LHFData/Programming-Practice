package main

import (
	"./coin"
	"./cli"
)


/*func main() {
	bc := coin.NewBlockchain()
	bc.AddBlock("send 1 BTC To me")
	bc.AddBlock("send 2 more to WY")
	for _, block := range bc.Blocks {
		fmt.Printf("prev.hash:%x\n", block.PrevBlockHash)
		fmt.Printf("Data:%s\n", block.Data)
		fmt.Printf("Hash:%x\n", block.Hash)
		pow := coin.NewProofOfWork(block)
		fmt.Printf("Pow:%s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()

	}

}*/

func main(){
	bc :=coin.NewBlockchain()
	defer bc.DB.Close()

	cli :=cli.CLI{bc}
	cli.Run()
}
