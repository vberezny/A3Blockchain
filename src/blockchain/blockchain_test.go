package blockchain

import (
	"encoding/hex"
	"fmt"
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

func TestValid(t *testing.T) {
	testBlock := Initial(2)
	testBlock.SetProof(242278)
	if testBlock.ValidHash() == false {
		t.Error("ValidHash() returned false, should have been true")
	}
	testBlock.SetProof(242277)
	if testBlock.ValidHash() == true {
		t.Error("ValidHash() returned true, should have been false")
	}
}

// test Run()
func TestRun(t *testing.T) {
	testBlock := Initial(2)
	worker := new(miningWorker)
	worker.Range = []uint64{242276, 242277, 242278}
	worker.blk = testBlock

	var i interface{} = worker.Run()
	mr := i.(MiningResult)
	if mr.Found == false {
		t.Error("Run should have returned a found proof")
	}

	worker.Range = []uint64{242276, 242277, 242279}
	var x interface{} = worker.Run()
	mr2 := x.(MiningResult)
	if mr2.Found {
		t.Error("Run should not have returned a found proof")
	}
}

// test Mine()
func TestMine(t *testing.T) {
	b0 := Initial(2)
	b0.Mine(1)
	fmt.Println(b0.Proof, hex.EncodeToString(b0.Hash))
	b1 := b0.Next("this is an interesting message")
	b1.Mine(1)
	fmt.Println(b1.Proof, hex.EncodeToString(b1.Hash))
	b2 := b1.Next("this is not interesting")
	b2.Mine(1)
	fmt.Println(b2.Proof, hex.EncodeToString(b2.Hash))
}

func TestMine2(t *testing.T) {
	b0 := Initial(3)
	b0.Mine(3)
	fmt.Println(b0.Proof, hex.EncodeToString(b0.Hash))
	b1 := b0.Next("this is an interesting message")
	b1.Mine(3)
	fmt.Println(b1.Proof, hex.EncodeToString(b1.Hash))
}

func TestAdd(t *testing.T) {
	b0 := Initial(2)
	b0.SetProof(242278)
	chn := new(Blockchain)
	chn.Chain = make([]Block, 1)
	chn.Add(b0)
	fmt.Println(chn.Chain)
	b1 := b0.Next("this is an interesting message")
	b1.Mine(1)
	chn.Add(b1)
	fmt.Println(chn.Chain)
	b2 := b1.Next("this is not interesting")
	b2.Mine(1)
	chn.Add(b2)
	fmt.Println(chn.Chain)
}

func TestValidChain(t *testing.T) {
	b0 := Initial(2)
	b0.SetProof(242278)
	chn := new(Blockchain)
	chn.Chain = make([]Block, 0)
	chn.Add(b0)

	b1 := b0.Next("this is an interesting message")
	b1.Mine(1)
	chn.Add(b1)

	b2 := b1.Next("this is not interesting")
	b2.Mine(1)
	chn.Add(b2)

	if !chn.IsValid() {
		t.Error("Chain validation failed")
	}
}
