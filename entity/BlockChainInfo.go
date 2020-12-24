package entity

import "math/big"

type BlockChainInfo struct {
	Chain string
	Blocks int64
	Headers int64
	BestBlockHash string
	Difficulty float64
	Mediantime big.Int
	Verificationprogress float64
	Initialblockdownload bool
	Chainblock string
	Size_on_desk big.Int
	Pruned bool

	Softfork SortFork
	Warnings string

}
