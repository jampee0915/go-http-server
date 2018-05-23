package main

import (
	"fmt"
	"net/http"
)

type Handler struct{}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func main() {

	handler := Handler{}
	server := http.Server{
		Addr:    ":8080", // = 127.0.0.1:8080
		Handler: &handler,
	}

	/* 第1引数: ネットワークアドレス(デフォルト80)
	 * 第2引数: リクエストを処理するハンドラ(nilの場合はDefaultServeMux)
	 * http.ListenAndServe("", nil)
	 */
	server.ListenAndServe()
}
