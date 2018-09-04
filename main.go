package main

import (
	log "github.com/sirupsen/logrus"
)

func main() {
	bc := NewBlockchain()

	bc.AddBlock("Send 1 BTC to dracher")
	bc.AddBlock("Send 2 more BTC to dracher")

	for _, block := range bc.blocks {
		log.Infof("Prev.hash: %x", block.PrevBlockHash)
		log.Infof("Data: %s", block.Data)
		log.Infof("Hash: %x", block.Hash)
		log.Infof("Nonce: %d", block.Nonce)
		log.Info("")
	}
}
