package main

import (
	"fmt"

	"github.com/debugtitan/vibes-blockchain/blockchain"
)


func main(){
	chain := blockchain.InitBlockChain()
	chain.AddBlock("First Block after Genesis")
	chain.AddBlock("First Block after Genesis")
	chain.AddBlock("First Block after Genesis")
	for _, block := range chain.Blocks{
		fmt.Printf("Previous Hash: %x\n",block.PrevHash)
		fmt.Printf("Data in Block: %x\n",block.Data)
		fmt.Printf("Hash: %x\n",block.Hash)
	}

}