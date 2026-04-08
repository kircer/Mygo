package main

import (
	"html/template"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Person struct {
	UserID   string `db:"id"`
	Username string `db:"name"`
	Age      int    `db:"age"`
	Address  string `db:"address"`
}

var db *sqlx.DB

var usersPageTpl = template.Must(template.New("users").Parse(`
<!doctype html>
<html lang="zh-CN">
<head>
  <meta charset="UTF-8" />
  <title>MySQL 用户列表</title>
  <style>
    body { font-family: "Microsoft YaHei", sans-serif; margin: 24px; }
    table { border-collapse: collapse; width: 100%; max-width: 900px; }
    th, td { border: 1px solid #ddd; padding: 10px; text-align: left; }
    th { background: #f5f7fa; }
    tr:nth-child(even) { background: #fafafa; }
  </style>
</head>
<body>
  <h2>用户列表</h2>
  <table>
    <thead>
      <tr>
        <th>ID</th>
        <th>姓名</th>
        <th>年龄</th>
        <th>地址</th>
      </tr>
    </thead>
    <tbody>
      {{range .}}
      <tr>
        <td>{{.UserID}}</td>
        <td>{{.Username}}</td>
        <td>{{.Age}}</td>
        <td>{{.Address}}</td>
      </tr>
      {{else}}
      <tr>
        <td colspan="4">暂无数据</td>
      </tr>
      {{end}}
    </tbody>
  </table>
</body>
</html>
`))

func initDB() error {
	conn, err := sqlx.Connect("mysql", "root:Zhaook123!@tcp(127.0.0.1:3306)/data?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		return err
	}
	db = conn
	return nil
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	var persons []Person
	err := db.Select(&persons, "SELECT id, name, age, address FROM user")
	if err != nil {
		http.Error(w, "查询数据库失败: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := usersPageTpl.Execute(w, persons); err != nil {
		http.Error(w, "页面渲染失败: "+err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	if err := initDB(); err != nil {
		log.Fatal("连接 MySQL 失败: ", err)
	}
	defer db.Close()

	http.HandleFunc("/", usersHandler)
	log.Println("服务已启动: http://127.0.0.1:9090")
	log.Fatal(http.ListenAndServe(":9090", nil))
}
