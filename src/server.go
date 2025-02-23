package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var db *sql.DB

func init() {
	var err error
	connStr := "user=postgres password=l2637962847 dbname=loginandregister sslmode=disable"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("无法打开数据库连接", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("无法连接到数据库", err)
	}
	fmt.Println("数据库连接成功!")
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "仅支持POST请求", http.StatusMethodNotAllowed)
		return
	}

	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "解析JSON失败", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var exists bool

	err = db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE username = $1)", user.Username).Scan(&exists)

	if err != nil {
		http.Error(w, "检查数据库时出错", http.StatusInternalServerError)
		return
	}
	if exists {
		json.NewEncoder(w).Encode(map[string]string{"message": "用户名已存在,请直接登录"})
		return
	}

	// 哈希加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "密码加密失败", http.StatusInternalServerError)
		return
	}

	_, err = db.Exec("INSERT INTO users (username, password) VALUES ($1, $2)", user.Username, hashedPassword)
	if err != nil {
		http.Error(w, "注册失败", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "注册成功"})
	fmt.Printf("注册用户: %s\n 密码； %s\n", user.Username, user.Password)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user User
	json.NewDecoder(r.Body).Decode(&user)

	//检查数据库是否存在该用户
	var storedPasswordHash string
	err := db.QueryRow("SELECT password FROM users WHERE username = $1", user.Username).Scan(&storedPasswordHash)
	if err == sql.ErrNoRows {
		json.NewEncoder(w).Encode(map[string]string{"message": "用户名不存在, 请先注册"})
		return
	} else if err != nil {
		query := "SELECT password FROM users WHERE username = ?"
		fmt.Println("Executing SQL:", query, "with username:", user.Username)
		log.Println("数据查询错误", err)
		http.Error(w, "数据库错误", http.StatusInternalServerError)
		return
	}

	//密码校验
	err = bcrypt.CompareHashAndPassword([]byte(storedPasswordHash), []byte(user.Password))
	if err != nil {
		log.Println(err)
		json.NewEncoder(w).Encode(map[string]string{"message": "密码错误"})
		return
	}

	// 登录成功，返回响应
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "登陆成功"})
}

func forgetHandle(w http.ResponseWriter, r *http.Request) {

}

func main() {
	fmt.Println("服务器启动中...")
	http.Handle("/register/", http.StripPrefix("/register", http.FileServer(http.Dir("../register"))))
	http.Handle("/login/", http.StripPrefix("/login", http.FileServer(http.Dir("../login"))))
	http.HandleFunc("/registerHandle", registerHandler)
	http.HandleFunc("/loginHandle", loginHandler)
	http.HandleFunc("/forgetHandle", forgetHandle)
	fmt.Println("服务器运行在 http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
