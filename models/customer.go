package models

type Customer struct {
	Customer_Number string `gorm:"PrimaryKey" json:"Customer_Number"`
	Customer_Name   string `gorm:"type:varchar(300)" json:"Customer_Name"`
}
