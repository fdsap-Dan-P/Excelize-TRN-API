package request

type TransactionRequest struct {
	Start_Date string `json:"start_date"`
	End_Date   string `json:"end_date"`
	File_Name  string `json:"file_name"`
}
