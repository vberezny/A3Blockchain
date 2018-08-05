package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

type Block struct {
	Generation uint64
	Difficulty uint8
	Data       string
	PrevHash   []byte
	Hash       []byte
	Proof      uint64
}

// Create new initial (generation 0) block.
func Initial(difficulty uint8) Block {
	blk := new(Block)

	arr := make([]byte, 32)

	for i := 0; i < 32; i++ {
		arr[i] = '\x00'
	}

	blk.Difficulty = difficulty
	blk.Generation = 0
	blk.Data = ""
	blk.PrevHash = arr

	return *blk
}

// Create new block to follow this block, with provided data.
func (prev_block Block) Next(data string) Block {
	new_blk := new(Block)

	new_blk.Data = data
	new_blk.Generation = prev_block.Generation + 1
	new_blk.Difficulty = prev_block.Difficulty
	new_blk.PrevHash = prev_block.Hash

	return *new_blk
}

// Calculate the block's hash.
func (blk Block) CalcHash() []byte {
	//create hash string out of block fields
	str := hex.EncodeToString(blk.PrevHash)
	str += ":"
	str += string(blk.Generation)
	str += ":"
	str += string(blk.Difficulty)
	str += ":"
	str += blk.Data
	str += ":"
	str += string(blk.Proof)

	fmt.Println([]byte(str))

	//hash the string
	h := sha256.New()
	h.Write([]byte(str))
	fmt.Printf("%x\n", h.Sum(nil))

	//return hashed string
	return h.Sum(nil)
}

/*
// Is this block's hash valid?
func (blk Block) ValidHash() bool {
	// TODO
}

// Set the proof-of-work and calculate the block's "true" hash.
func (blk *Block) SetProof(proof uint64) {
	blk.Proof = proof
	blk.Hash = blk.CalcHash()
}
*/
