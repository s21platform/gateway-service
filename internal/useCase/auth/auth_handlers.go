package auth

import (
	"fmt"
	"net/http"
)

func GetAuth(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET REQUEST")
	w.Write([]byte("Work"))
}
