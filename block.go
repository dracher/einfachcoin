package main

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

// Block is
type Block struct {
	Timestamp     time.Time
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
	Nonce         int
}

// SetHash is
func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp.Unix(), 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)
	b.Hash = hash[:]
}

// NewBlock is
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now(), []byte(data), prevBlockHash, []byte{}, 0}
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}
