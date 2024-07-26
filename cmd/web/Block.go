package main

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

func (block *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(block.timestamp, 10))
	headers := bytes.Join([][]byte{timestamp, block.perviousBlockHash, block.AllData}, []byte{})
	hash := sha256.Sum256(headers)
	block.myBlockHash = hash[:]
}

func NewBlock(data string, pervBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), pervBlockHash, []byte{}, []byte(data)}
	block.SetHash()
	return block
}

func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}
