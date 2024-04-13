package webapp

import (
	"encoding/json"
	"log"
	"net/http"
)

// RequestData 定义了请求的结构，包括两个加数
type RequestData struct {
	Num1 int `json:"num1"` // 第一个加数
	Num2 int `json:"num2"` // 第二个加数
}

// ResponseData 定义了响应的结构，包括和
type ResponseData struct {
	Sum int `json:"sum"` // 和
}

func main() {
	http.HandleFunc("/add", addHandler) // 设置处理函数
	log.Println("Server starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil)) // 启动服务器
}

// addHandler 处理/add路由的POST请求
func addHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	var data RequestData
	// 解析请求中的JSON体
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// 计算结果
	result := ResponseData{
		Sum: data.Num1 + data.Num2,
	}

	// 发送响应
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(result); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	func max(a, b int) int {
    if a > b {
        return a
    }
    return b
  }
}
