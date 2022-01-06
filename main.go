package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"

	"github.com/moritamori/go-oidc-sample/platform/authenticator"
	"github.com/moritamori/go-oidc-sample/platform/router"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("envファイルの読み込みに失敗しました: %v", err)
	}

	auth, err := authenticator.New()
	if err != nil {
		log.Fatalf("認証器の生成に失敗しました: %v", err)
	}

	rtr := router.New(auth)

	log.Print("Server listening on http://localhost:3000/")
	if err := http.ListenAndServe("0.0.0.0:3000", rtr); err != nil {
		log.Fatalf("HTTPサーバーの起動時にエラーが発生しました: %v", err)
	}
}
