package router

import (
	"encoding/gob"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"

	"github.com/moritamori/go-oidc-sample/platform/authenticator"
	"github.com/moritamori/go-oidc-sample/web/app/callback"
	"github.com/moritamori/go-oidc-sample/web/app/login"
	"github.com/moritamori/go-oidc-sample/web/app/user"
)

// New はroutesを登録し、ルーターを返す
func New(auth *authenticator.Authenticator) *gin.Engine {
	router := gin.Default()

	// セッション設定
	gob.Register(map[string]interface{}{})
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("auth-session", store))

	// HTMLテンプレートパス設定
	router.LoadHTMLGlob("web/template/*")

	// 無名ハンドラ
	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "home.html", nil)
	})
	// ログインハンドラ
	router.GET("/login", login.Handler(auth))
	// コールバックハンドラ
	router.GET("/callback", callback.Handler(auth))
	// ユーザーハンドラ
	router.GET("/user", user.Handler)

	return router
}
