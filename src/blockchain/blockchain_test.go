package blockchain

import (
	"encoding/hex"
	"testing"
)

// TODO: some useful tests of Blocks

// Tests Initial()
func TestInitial(t *testing.T) {
	testBlock := Initial(2)

	arr := make([]byte, 32)

	for i := 0; i < 32; i++ {
		arr[i] = '\x00'
	}

	if testBlock.Data != "" {
		t.Error("Data:", testBlock.Data)
	}
	if testBlock.Difficulty != 2 {
		t.Error("Difficulty:", testBlock.Difficulty)
	}
	if testBlock.Generation != 0 {
		t.Error("Generation:", testBlock.Generation)
	}
}

func TestCalcHash(t *testing.T) {
	testBlock := Initial(2)
	testBlock.Proof = 242278
	hash := hex.EncodeToString(testBlock.CalcHash())
	//test first string
	if hash != "29528aaf90e167b2dc248587718caab237a81fd25619a5b18be4986f75f30000" {
		t.Error("Initial hash is incorrect\nRecieved Hash: ", hash)
	}
}
