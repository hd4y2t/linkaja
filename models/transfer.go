package models

type Transfer struct {
	To_Account_Number string  `json:"to_account_number" binding:"required"`
	Amount            float64 `json:"amount" binding:"required"`
}
