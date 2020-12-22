package entity

type ChainInfo struct {
	Chain string`json:"chain"`
	Blocks int`json:"blocks"`
	Headers int`json:"headers"`
	Bestblockhash string`json:"bestblockhash"`
	Difficulty int`json:"difficulty"`
	Mediantime int64`json:"mediantime"`
	Verificationprogress float64`json:"verificationprogress"`
	Initialblockdownload bool`json:"initialblockdownload"`
	Chainword string`json:"chainword"`
	Size_on_disk int`json:"size_on_disk"`
	Pruned bool`json:"pruned"`
	//Pruneheight int`json:"pruneheight"`
	//Automatic_pruning bool`json:"automatic_pruning"`
	//Prune_target_size int64`json:"prune_target_size"`
	Softforks Bipsoftfork`json:"softforks"`
	Warnings string`json:"warnings"`
}
