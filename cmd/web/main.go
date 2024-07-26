package main

import "fmt"

func main() {
	newblockchain := NewBlockchain() // Initialize the blockchain
	// create 2 blocks and add 2 transactions
	newblockchain.AddBlock("first transaction")  // first block containing one tx
	newblockchain.AddBlock("Second transaction") // second block containing one tx
	// Now print all the blocks and their contents
	for _, block := range newblockchain.Blocks { // iterate on each block
		fmt.Printf("Hash of the block %x\n", block.myBlockHash)                 // print the hash of the block
		fmt.Printf("Hash of the previous Block: %x\n", block.perviousBlockHash) // print the hash of the previous block
		fmt.Printf("All the transactions: %s\n", block.AllData)                 // print the transactions
	} // our blockchain will be printed
}
