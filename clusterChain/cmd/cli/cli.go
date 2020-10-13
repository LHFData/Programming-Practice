package cli

import (
	"../coin"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)
type CLI struct{
	BC *coin.Blockchain
}
func (cli *CLI) printUsage(){
	fmt.Println("Usage")
	fmt.Println("addblock -data BLOCK_DATA - add a block to the blockchain")
	fmt.Println("printchain -print all the blocks of the blockchain")
}
func (cli *CLI) addBlock(data string){
	cli.BC.AddBlock(data)
	fmt.Println("Success!")
}
func (cli *CLI) validateArgs(){
	if len(os.Args) <2{
		cli.printUsage()
		os.Exit(1)
	}
}
func(cli *CLI) printChain(){
	bci :=cli.BC.Iterator()

	for {
		block :=bci.Next()

		fmt.Printf("Prev.hash:%x\n",block.PrevBlockHash)
		fmt.Printf("Data:%s\n",block.Data)
		fmt.Printf("Hash:%x\n",block.Hash)
		pow:=coin.NewProofOfWork(block)
		fmt.Printf("Pow:%s\n",strconv.FormatBool(pow.Validate()))
		fmt.Println()

		if len(block.PrevBlockHash)==0{
			break
		}

	}

}
func(cli *CLI) Run(){
	cli.validateArgs()
	for {
		addBlockCmd := flag.NewFlagSet("addblock", flag.ExitOnError)

		printChaincmd := flag.NewFlagSet("printchain", flag.ExitOnError)


		addBlockData := addBlockCmd.String("data", "", "Block data")

		switch os.Args[1] {
		case "addblock":
			err := addBlockCmd.Parse(os.Args[2:])
			if err != nil {
				log.Panic(err)
			}
		case "printChain":
			err := printChaincmd.Parse(os.Args[2:])
			if err != nil {
				log.Panic(err)
			}
		case "exit":
			os.Exit(1)
		default:
			cli.printUsage()
			os.Exit(1)

		}
		if addBlockCmd.Parsed() {
			if *addBlockData == "" {
				addBlockCmd.Usage()
				os.Exit(1)
			}
		}
		if printChaincmd.Parsed() {
			cli.printChain()
		}
	}
}