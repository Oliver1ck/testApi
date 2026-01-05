package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type responseUser struct {
	UserName string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var req responseUser
	if r.Method != http.MethodPost {
		fmt.Println("Ошибка метода")
		w.Header().Set("Allow", "POST, OPTIONS")
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		fmt.Println("ошибка декодирования %w", err)
	}
	fmt.Println(req)

	response, err := json.MarshalIndent(req, "", "	")
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Write(response)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	msg := "привет, я ответил"
	var req responseUser
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		fmt.Println("ошибка декодирования %w", err)
	}
	fmt.Println("Запрос", req)
	w.Write([]byte(msg))
}
