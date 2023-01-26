package response

type TransactionResponse struct {
	Got_amount         string `json:"got_amount"`
	Got_reference      string `string:"got_reference"`
	Got_date           string `json:"got_date"`
	Got_time           string `json:"got_time"`
	Got_iid            string `json:"got_iid"`
	Got_cid            string `json:"got_cid"`
	Got_particular     string `jaon:"got_particular"`
	Got_source_account string `json:"got_source_account"`
	Got_target         string `json:"got_target"`
	Got_total_fee      string `json:"got_total_fee"`
	Got_agent_fee      string `json:"got_agent_fee"`
	Got_bank_fee       string `json:"got_bank_fee"`
	Got_standing       string `json:"got_standing"`
}

type FormulaFromDB struct {
	Formula string `json:"formula"`
}

type GetCellValuePath struct {
	Path string `json:"path"`
}

type RowCount struct {
	Count int `json:"count"`
}

type GetFilePath struct {
	MyPath string `json:"mypath"`
}
