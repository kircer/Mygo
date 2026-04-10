package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// 1. 我们手动定义一个结构体来实现 Handler 接口
type UserHandler struct{}

// 2. 必须叫 ServeHTTP，这是契约
func (u *UserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 解析 JSON
	var data map[string]interface{}
	json.NewDecoder(r.Body).Decode(&data)
	data["message"] = "获取用户成功"

	// 返回 JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func main() {
	// 3. 原生注册路由的方式：指定路径 + Handler
	http.Handle("/user", &UserHandler{})

	// 启动服务
	fmt.Println("Server running on :8080")
	http.ListenAndServe(":9090", nil)
}
