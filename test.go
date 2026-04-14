package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

type Person struct {
	UserID   string `db:"id" json:"id"`
	Username string `db:"name" json:"name"`
	Age      int    `db:"age" json:"age"`
	Address  string `db:"address" json:"address"`
}

type usersResponse struct {
	Data  []Person `json:"data"`
	Error string   `json:"error,omitempty"`
}

func initDB() error {
	dsn := os.Getenv("MYSQL_DSN")
	if dsn == "" {
		dsn = "root:Zhaook123!@tcp(127.0.0.1:3306)/data?parseTime=true"
	}

	conn, err := sqlx.Open("mysql", dsn)
	if err != nil {
		return err
	}

	conn.SetMaxOpenConns(10)
	conn.SetMaxIdleConns(5)
	conn.SetConnMaxLifetime(30 * time.Minute)

	if err := conn.Ping(); err != nil {
		_ = conn.Close()
		return err
	}

	db = conn
	return nil
}

func listUsers() ([]Person, error) {
	users := make([]Person, 0)
	if err := db.Select(&users, "SELECT id, name, age, address FROM user"); err != nil {
		return nil, err
	}
	return users, nil
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	users, err := listUsers()
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, usersResponse{Error: err.Error()})
		return
	}

	writeJSON(w, http.StatusOK, usersResponse{Data: users})
}

func writeJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(payload); err != nil {
		log.Printf("encode response failed: %v", err)
	}
}

func withCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	if err := initDB(); err != nil {
		log.Fatalf("connect mysql failed: %v", err)
	}
	defer db.Close()

	mux := http.NewServeMux()
	mux.HandleFunc("/api/users", usersHandler)

	addr := ":8080"
	log.Printf("API server listening on %s", addr)
	if err := http.ListenAndServe(addr, withCORS(mux)); err != nil && err != http.ErrServerClosed {
		log.Fatalf("server stopped with error: %v", err)
	}
}
