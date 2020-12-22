package entity

type Softfork struct {
	Id string`json:"type"`
	Version bool`json:"active"`
	Reject status`json:"height"`
}

type status struct {
	Status bool `json:"status"`
}