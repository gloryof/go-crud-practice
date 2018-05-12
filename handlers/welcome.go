package handlers

import (
	"fmt"
	"net/http"
)

// WelcomeHandler Welcomeページのハンドラ
type WelcomeHandler struct{}

// ServeHTTP 表示処理
func (h *WelcomeHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "welcome")
}
