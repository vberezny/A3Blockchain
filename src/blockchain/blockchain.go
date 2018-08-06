package blockchain

import (
	"encoding/hex"
	"reflect"
)

type Blockchain struct {
	Chain []Block
}

func (chain *Blockchain) Add(blk Block) {
	if !blk.ValidHash() {
		panic("adding block with invalid hash")
	}
	chain.Chain = append(chain.Chain, blk)
}

func (chain Blockchain) IsValid() bool {
	if chain.Chain[0].Generation != 0 {
		return false
	}
	for i := range chain.Chain[0].PrevHash {
		if chain.Chain[0].PrevHash[i] != 0 {
			return false
		}
	}
	diff := chain.Chain[0].Difficulty
	gen := uint64(0)
	for a := range chain.Chain {
		if chain.Chain[a].Difficulty != diff {
			return false
		}
		if chain.Chain[a].Generation != gen {
			return false
		}
		if a > 0 && !reflect.DeepEqual(chain.Chain[a].PrevHash, chain.Chain[a-1].Hash) {
			return false
		}
		if hex.EncodeToString(chain.Chain[a].Hash) != hex.EncodeToString(chain.Chain[a].CalcHash()) {
			return false
		}
		if !chain.Chain[a].ValidHash() {
			return false
		}
		gen++
	}

	return true
}
