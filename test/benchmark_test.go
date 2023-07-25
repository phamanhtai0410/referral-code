package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/google/uuid"
)

func BenchmarkAPISAVE(b *testing.B) {
	for i := 0; i < b.N; i++ {
		address := uuid.New().String()
		requestBody := map[string]interface{}{
			"refcode": 111182,
			"domain":  "kawasaki.meme",
			"price":   1.2,
			"address": address,
		}

		jsonData, _ := json.Marshal(requestBody)
		url := "http://127.0.0.1:5000/save"
		resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
		if err != nil {
			fmt.Println("Lỗi khi gửi yêu cầu:", err)
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			b.Errorf("Expected status code %d, but got %d", http.StatusOK, resp.StatusCode)
		}
	}
}
