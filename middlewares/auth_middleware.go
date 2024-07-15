package middlewares

import (
	"go-gin-productManagerPro/services"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddlware(authService services.IAuthService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		header := ctx.GetHeader("Authorization")
		if header == "" {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Authorizationをキーに持つバリューチェック
		if !strings.HasPrefix(header, "Bearer ") {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// 認証に必要なトークン情報を取得
		tokenString := strings.TrimPrefix(header, "Bearer ")
		user, err := authService.GetUserFromToken(tokenString)
		if err != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		ctx.Set("user", user)

		ctx.Next()
	}
}

/*
1. リクエスト情報:
ctx.Request で *http.Request オブジェクトにアクセスできる
ctx.GetHeader("Header-Name") で特定のヘッダー情報を取得できる
ctx.PostForm("key") でPOSTフォームデータを取得できる
2. レスポンス情報:
ctx.Writer で http.ResponseWriter オブジェクトにアクセスできる
ctx.JSON(statusCode, jsonObj) でJSONレスポンスを送信できる
3. パラメータ:
ctx.Param("paramName") でURLパラメータを取得できる
ctx.Query("queryParam") でクエリパラメータを取得できる
4. コンテキストデータ:
ctx.Set("key", value) でデータを設定できる
ctx.Get("key") で設定したデータを取得できる
*/
