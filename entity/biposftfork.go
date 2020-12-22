package entity

type Bipsoftfork struct {
	Bip Softfork`json:"bip"`
	Csv Softfork`json:"csv"`
	Srgwit Softfork`json:"srgwit"`
}
