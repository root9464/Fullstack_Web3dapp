package response

type Data struct {
	Transaction Transaction `json:"transaction"`
	Interfaces  []string    `json:"interfaces"`
	Children    []Child     `json:"children"`
}

type Transaction struct {
	Hash            string       `json:"hash"`
	Lt              int64        `json:"lt"`
	Account         Account      `json:"account"`
	Success         bool         `json:"success"`
	Utime           int64        `json:"utime"`
	OrigStatus      string       `json:"orig_status"`
	EndStatus       string       `json:"end_status"`
	TotalFees       int64        `json:"total_fees"`
	EndBalance      int64        `json:"end_balance"`
	TransactionType string       `json:"transaction_type"`
	StateUpdateOld  string       `json:"state_update_old"`
	StateUpdateNew  string       `json:"state_update_new"`
	InMsg           InMessage    `json:"in_msg"`
	OutMsgs         []OutMessage `json:"out_msgs"`
	Block           string       `json:"block"`
	PrevTransHash   string       `json:"prev_trans_hash"`
	PrevTransLt     int64        `json:"prev_trans_lt"`
	ComputePhase    ComputePhase `json:"compute_phase"`
	StoragePhase    StoragePhase `json:"storage_phase"`
	ActionPhase     ActionPhase  `json:"action_phase"`
	Aborted         bool         `json:"aborted"`
	Destroyed       bool         `json:"destroyed"`
	Raw             string       `json:"raw"`
}

type Account struct {
	Address  string `json:"address"`
	IsScam   bool   `json:"is_scam"`
	IsWallet bool   `json:"is_wallet"`
}

type InMessage struct {
	MsgType       string      `json:"msg_type"`
	CreatedLt     int64       `json:"created_lt"`
	IhrDisabled   bool        `json:"ihr_disabled"`
	Bounce        bool        `json:"bounce"`
	Bounced       bool        `json:"bounced"`
	Value         int64       `json:"value"`
	FwdFee        int64       `json:"fwd_fee"`
	IhrFee        int64       `json:"ihr_fee"`
	Destination   Account     `json:"destination"`
	Source        Account     `json:"source"`
	ImportFee     int64       `json:"import_fee"`
	CreatedAt     int64       `json:"created_at"`
	Hash          string      `json:"hash"`
	RawBody       string      `json:"raw_body"`
	DecodedOpName string      `json:"decoded_op_name"`
	DecodedBody   DecodedBody `json:"decoded_body"`
}

type DecodedBody struct {
	Signature   string    `json:"signature"`
	SubwalletID int64     `json:"subwallet_id"`
	ValidUntil  int64     `json:"valid_until"`
	Seqno       int       `json:"seqno"`
	Op          int       `json:"op"`
	Payload     []Payload `json:"payload"`
}

type Payload struct {
	Mode    int            `json:"mode"`
	Message PayloadMessage `json:"message"`
}

type PayloadMessage struct {
	SumType         string          `json:"sum_type"`
	MessageInternal MessageInternal `json:"message_internal"`
}

type MessageInternal struct {
	IhrDisabled bool      `json:"ihr_disabled"`
	Bounce      bool      `json:"bounce"`
	Bounced     bool      `json:"bounced"`
	Src         string    `json:"src"`
	Dest        string    `json:"dest"`
	Value       Value     `json:"value"`
	IhrFee      string    `json:"ihr_fee"`
	FwdFee      string    `json:"fwd_fee"`
	CreatedLt   int64     `json:"created_lt"`
	CreatedAt   int64     `json:"created_at"`
	Init        *InitData `json:"init"`
	Body        Body      `json:"body"`
}

type Value struct {
	Grams string            `json:"grams"`
	Other map[string]string `json:"other"`
}

type InitData struct {
	Boc        string   `json:"boc"`
	Interfaces []string `json:"interfaces"`
}

type Body struct {
	IsRight bool        `json:"is_right"`
	Value   interface{} `json:"value"`
}

type OutMessage struct {
	// Add fields if necessary, similar to InMessage
}

type ComputePhase struct {
	Skipped      bool   `json:"skipped"`
	SkipReason   string `json:"skip_reason"`
	Success      bool   `json:"success"`
	GasFees      int64  `json:"gas_fees"`
	GasUsed      int64  `json:"gas_used"`
	VmSteps      int    `json:"vm_steps"`
	ExitCode     int    `json:"exit_code"`
	ExitCodeDesc string `json:"exit_code_description"`
}

type StoragePhase struct {
	FeesCollected int64  `json:"fees_collected"`
	StatusChange  string `json:"status_change"`
}

type ActionPhase struct {
	Success        bool   `json:"success"`
	ResultCode     int    `json:"result_code"`
	TotalActions   int    `json:"total_actions"`
	SkippedActions int    `json:"skipped_actions"`
	FwdFees        int64  `json:"fwd_fees"`
	TotalFees      int64  `json:"total_fees"`
	ResultCodeDesc string `json:"result_code_description"`
}

type Child struct {
	Transaction Transaction `json:"transaction"`
	Interfaces  []string    `json:"interfaces"`
}
