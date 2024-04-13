package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAddHandler(t *testing.T) {
	// 创建一个模拟的请求数据
	requestBody, _ := json.Marshal(map[string]int{
		"num1": 10,
		"num2": 15,
	})
	request := httptest.NewRequest("POST", "/add", bytes.NewBuffer(requestBody))
	request.Header.Set("Content-Type", "application/json")

	// 创建一个响应记录器来捕获响应
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(addHandler)

	// 调用 addHandler 函数
	handler.ServeHTTP(responseRecorder, request)

	// 检查状态码是否为200
	if status := responseRecorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// 检查返回的内容
	var resp ResponseData
	if err := json.Unmarshal(responseRecorder.Body.Bytes(), &resp); err != nil {
		t.Fatal("could not unmarshal response:", err)
	}

	expectedSum := 25
	if resp.Sum != expectedSum {
		t.Errorf("handler returned unexpected sum: got %v want %v", resp.Sum, expectedSum)
	}
}
