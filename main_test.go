package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	accountcontroller "github.com/hd4y2t/go_linkaja/controllers"
	"github.com/stretchr/testify/assert"
)

type TransferRequest struct {
	To_Account_Number string  `json:"to_account_number" binding:"required"`
	Amount            float64 `json:"amount" binding:"required"`
}

func TestTransfer(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	// Tentukan router yang akan digunakan dalam pengujian
	r.POST("/api/account/:account_number/transfer", accountcontroller.Update)
	destination_account_number := "555002"
	// Buat data JSON untuk request
	requestData := TransferRequest{
		To_Account_Number: destination_account_number,
		Amount:            5000,
	}
	jsonData, _ := json.Marshal(requestData)

	// Buat request HTTP palsu untuk pengujian
	req, _ := http.NewRequest(http.MethodPost, "/api/account/555001/transfer", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Periksa status code dan respons
	assert.Equal(t, http.StatusOK, w.Code)
	// Pastikan Anda melakukan pengujian yang lebih mendalam terhadap respons JSON jika diperlukan

}

func TestTransfer1(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	// Tentukan router yang akan digunakan dalam pengujian
	r.POST("/api/account/:account_number/transfer", accountcontroller.Update)
	destination_account_number := "555001"
	// Buat data JSON untuk request
	requestData := TransferRequest{
		To_Account_Number: destination_account_number,
		Amount:            5000,
	}
	jsonData, _ := json.Marshal(requestData)

	// Buat request HTTP palsu untuk pengujian
	req, _ := http.NewRequest(http.MethodPost, "/api/account/555001/transfer", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Periksa status code dan respons
	assert.Equal(t, http.StatusOK, w.Code)
	// Pastikan Anda melakukan pengujian yang lebih mendalam terhadap respons JSON jika diperlukan

}

func TestTransfer3(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	// Tentukan router yang akan digunakan dalam pengujian
	r.POST("/api/account/:account_number/transfer", accountcontroller.Update)
	destination_account_number := "555002"
	// Buat data JSON untuk request
	requestData := TransferRequest{
		To_Account_Number: destination_account_number,
		Amount:            5000000,
	}
	jsonData, _ := json.Marshal(requestData)

	// Buat request HTTP palsu untuk pengujian
	req, _ := http.NewRequest(http.MethodPost, "/api/account/555001/transfer", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Periksa status code dan respons
	assert.Equal(t, http.StatusOK, w.Code)
	// Pastikan Anda melakukan pengujian yang lebih mendalam terhadap respons JSON jika diperlukan

}

func TestShow(t *testing.T) {
	// Transfer := accountcontroller.Transfer
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	r.GET("/api/account/:account_number", accountcontroller.Show)

	req, _ := http.NewRequest(http.MethodGet, "/api/account/555001", nil)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestShow1(t *testing.T) {
	// Transfer := accountcontroller.Transfer
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	r.GET("/api/account/:account_number", accountcontroller.Show)

	req, _ := http.NewRequest(http.MethodGet, "/api/account/555004", nil)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
