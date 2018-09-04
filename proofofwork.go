package main

import (
	"bytes"
	"crypto/sha256"
	"math"
	"math/big"

	log "github.com/sirupsen/logrus"
)

var maxNonce = math.MaxInt64

const targetBits = 25

// ProofOfWork is
type ProofOfWork struct {
	block  *Block
	target *big.Int
}

// NewProofOfWork is
func NewProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetBits))

	pow := &ProofOfWork{b, target}
	return pow
}

func (pow *ProofOfWork) prepareData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.block.PrevBlockHash,
			pow.block.Data,
			IntToHex(pow.block.Timestamp.Unix()),
			IntToHex(int64(targetBits)),
			IntToHex(int64(nonce)),
		},
		[]byte{},
	)

	return data
}

// Run is
func (pow *ProofOfWork) Run() (int, []byte) {
	var hashInt big.Int
	var hash [32]byte
	nonce := 0

	log.Infof("Mining the block containing \"%s\"", pow.block.Data)

	for nonce < maxNonce {
		data := pow.prepareData(nonce)
		hash = sha256.Sum256(data)
		// log.Infof("\r%x", hash)
		hashInt.SetBytes(hash[:])

		if hashInt.Cmp(pow.target) == -1 {
			break
		} else {
			nonce++
		}
	}

	log.Info("")
	return nonce, hash[:]
}
