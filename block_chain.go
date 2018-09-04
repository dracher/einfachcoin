package main

import (
	"fmt"

	bolt "github.com/etcd-io/bbolt"
)

// Blockchain is
type Blockchain struct {
	blocks []*Block
}

// AddBlock is
func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}

// NewGenesisBlock is
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}

// NewBlockchain is
func NewBlockchain() *Blockchain {
	var tip []byte
	_, err := bolt.Open(dbFile, 0600, nil)

	fmt.Println(err)

	return &Blockchain{[]*Block{NewGenesisBlock()}}
}
