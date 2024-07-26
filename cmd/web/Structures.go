package main

type Block struct {
	timestamp         int64  // the time when the block was created
	perviousBlockHash []byte // the hash of the previous block
	myBlockHash       []byte // the hash of the current block
	AllData           []byte // the data or transactions (body info)
}

type Blockchain struct {
	Blocks []*Block
}
