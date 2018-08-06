package blockchain

import (
	"crypto/sha256"
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

	/*for i := 0; i < 32; i++ {
		arr[i] = 0
	}*/

	if arr[31] == '\x00' {
		fmt.Println("YES")
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
	//create string
	str := fmt.Sprintf("%x:%d:%d:%s:%d", blk.PrevHash, blk.Generation, blk.Difficulty, blk.Data, blk.Proof)
	fmt.Println(str)

	//hash the string
	h := sha256.New()
	h.Write([]byte(str))
	fmt.Printf("Hash: %x\n", h.Sum(nil))

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
