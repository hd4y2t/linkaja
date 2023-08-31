package models

type Account struct {
	Account_Number  string  `gorm:"PrimaryKey" json:"Account_Number"`
	Customer_Number string  `gorm:"ForeignKey:Customer_Number;association_foreignkey:Customer_Number" json:"Customer_Number"`
	Balance         float64 `gorm:"type:int" json:"Balance"`
}
