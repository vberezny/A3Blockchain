package blockchain

import (
	"work_queue"
)

type miningWorker struct {
	work_queue.Worker
	Range []uint64
	blk   Block
}

type MiningResult struct {
	Proof uint64 // proof-of-work value, if found.
	Found bool   // true if valid proof-of-work was found.
}

func (m *miningWorker) Run() interface{} {
	res := new(MiningResult)
	res.Found = false
	// check if proof value is valid
	for i := range m.Range {
		m.blk.SetProof(m.Range[i])
		if m.blk.ValidHash() {
			res.Proof = m.Range[i]
			res.Found = true
			return *res
		}
	}
	return *res
}

// Mine the range of proof values, by breaking up into chunks and checking
// "workers" chunks concurrently in a work queue. Should return shortly after a result
// is found.
func (blk Block) MineRange(start uint64, end uint64, workers uint64, chunks uint64) MiningResult {
	q := work_queue.Create(uint(workers), uint(chunks))
	val := start
	interval := end / chunks
	for i := uint64(0); i < chunks; i++ {
		miner := new(miningWorker)
		miner.blk = blk
		arr := make([]uint64, interval)
		for i := uint64(0); i < interval; i++ {
			arr[i] = val
			val++
		}
		miner.Range = arr
		q.Enqueue(miner)
	}
	running := true
	for running {
		res := <-q.Results
		mr := res.(MiningResult)
		if mr.Found {
			q.Shutdown()
			running = false
			return mr
		}
	}
	return *new(MiningResult)
}

// Call .MineRange with some reasonable values that will probably find a result.
// Good enough for testing at least. Updates the block's .Proof and .Hash if successful.
func (blk *Block) Mine(workers uint64) bool {
	reasonableRangeEnd := uint64(4 * 1 << (8 * blk.Difficulty)) // 4 * 2^(bits that must be zero)
	mr := blk.MineRange(0, reasonableRangeEnd, workers, 4321)
	if mr.Found {
		blk.SetProof(mr.Proof)
	}
	return mr.Found
}
