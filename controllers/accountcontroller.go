package accountcontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hd4y2t/go_linkaja/models"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {
	var accounts []models.Account

	models.DB.Find(&accounts)
	c.JSON(http.StatusOK, gin.H{"account": accounts})
}

func Show(c *gin.Context) {
	var account models.Account
	account_number := c.Param("account_number")

	if err := models.DB.First(&account, account_number).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Account tidak ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err})
			return
		}
	}

	var customer models.Customer
	if err := models.DB.First(&customer, account.Customer_Number).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	accountWithCustomer := struct {
		Account_Number string
		Customer_Name  string
		Balance        float64
	}{
		Account_Number: account.Account_Number, Customer_Name: customer.Customer_Name, Balance: account.Balance}

	c.JSON(http.StatusOK, gin.H{"data": accountWithCustomer})

}

type TransferRequest struct {
	To_Account_Number string  `json:"to_account_number" binding:"required"`
	Amount            float64 `json:"amount" binding:"required"`
}

func Update(c *gin.Context) {
	account_Number := c.Param("account_number")

	var sourceAccount models.Account
	if err := models.DB.First(&sourceAccount, "account_number = ?", account_Number).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"message": "Account tidak ditemukan"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		}
		return
	}

	var transferData TransferRequest
	if err := c.ShouldBindJSON(&transferData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Request data salah"})
		return
	}

	to_account_number := transferData.To_Account_Number
	if account_Number == to_account_number {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "Tidak dapat melakukan transaksi diaccount yang sama"})
		return
	}

	amount := transferData.Amount

	if sourceAccount.Balance < amount {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "Saldo tidak mencukupi"})
		return
	}

	tx := models.DB.Begin()

	// Mengurangi saldo dari akun sumber
	sourceAccount.Balance -= amount
	if err := tx.Save(&sourceAccount).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal Update Saldo"})
		return
	}

	var destinationAccount models.Account
	if err := tx.First(&destinationAccount, "account_number = ?", transferData.To_Account_Number).Error; err != nil {
		tx.Rollback()
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"message": "Account tujuan salah"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		}
		return
	}

	destinationAccount.Balance += amount
	if err := tx.Save(&destinationAccount).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal update saldo"})
		return
	}

	tx.Commit()

	c.JSON(http.StatusOK, gin.H{"message": "Transfer Sukses"})

}
