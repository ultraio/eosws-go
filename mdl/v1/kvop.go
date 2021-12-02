package mdl

type KVOp struct {
	Op          string `json:"op,omitempty"`
	ActionIndex int    `json:"action_idx"`
	Account     string `json:"account,omitempty"`
	Key         string `json:"key,omitempty"`
	Old         *KVRow `json:"old,omitempty"`
	New         *KVRow `json:"new,omitempty"`
}

type KVRow struct {
	Payer string      `json:"payer,omitempty"`
	Hex   string      `json:"hex,omitempty"`
	JSON  interface{} `json:"json,omitempty"`
	Error string      `json:"error,omitempty"`
}
