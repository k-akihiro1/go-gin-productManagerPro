# go-gin-productManagerPro

This repository is implemented to refresh my memory of Go/Gin.

# Development Environment

Language: Go  
Framework: Gin  
ORM: Gorm  
Database: MySQL  
Authentication: JWT (gin-jwt)

# Adopted Architecture

Layered Architecture  
Controller

- Handling request data
- Setting response
  Service
- Functionality (implementation of business logic)
- Implemented functionalities:
  - Products: search all, search by ID, register, update, delete
  - Users: login, authentication features (applicable to: search by ID, register, update, delete)

```
Router - IController
         ↑
         Controller → IService
                      ↑
                      Service → IRepository
                                ↑
                                Repository → DB
```

References:
https://qiita.com/fghyuhi/items/8d5c0f7f8aec643e5907
https://zenn.dev/taiyou/articles/747ab00a61a2f2
https://qiita.com/koji0705/items/49172d713e13fa554ba7

# Gin の特徴

- パラメータ付き URL、グループ化、URL パターンマッチングなどの複雑なルーティングに対応

### 正規表現を使用したマッチング

```
router.GET("/user/:id([0-9]+)", func(c *gin.Context) {
    id := c.Param("id")
    // 数字のみのidを処理
})
```

:id([0-9]+)は 1 つ以上の数字のみにマッチする正規表現を使用し、数字以外の ID でアクセスした場合にはマッチしないため、そのルートは実行されない等の不正なリクエストを排除できる。

- ロギング、認証、CORS 処理などの共通処理を簡単に追加

### ロギング

リクエストとレスポンスに関する情報をログとして記録すること
アプリケーションの動作を監視し、問題発生時のデバッグに役立つ
Gin ではデフォルトでロギングミドルウェアが提供されており、次のようにして簡単に追加できる：

```
router := gin.Default()
```

### 認証

JWT（JSON Web Tokens）を使用した認証は、セキュアな API アクセスを提供するための一般的な方法

```
import (
    "github.com/appleboy/gin-jwt/v2"
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()
    authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
        Realm:       "test zone",
        Key:         []byte("secret key"),
        Timeout:     time.Hour,
        MaxRefresh:  time.Hour,
        Authenticator: func(c *gin.Context) (interface{}, error) {
            var loginVals login
            if err := c.ShouldBind(&loginVals); err != nil {
                return "", jwt.ErrMissingLoginValues
            }
            userID := loginVals.UserID
            password := loginVals.Password

            if (userID == "admin" && password == "admin") || (userID == "test" && password == "test") {
                return &User{
                    UserName:  userID,
                    LastName:  "Bo",
                    FirstName: "Li",
                }, nil
            }

            return nil, jwt.ErrFailedAuthentication
        },
    })

    if err != nil {
        log.Fatal("JWT Error:" + err.Error())
    }

    router.POST("/login", authMiddleware.LoginHandler)
    router.GET("/refresh_token", authMiddleware.RefreshHandler)

    router.GET("/secure", authMiddleware.MiddlewareFunc(), func(c *gin.Context) {
        claims := jwt.ExtractClaims(c)
        user, _ := c.Get(identityKey)
        c.JSON(200, gin.H{
            "userID":   claims["id"],
            "userName": user.(*User).UserName,
            "text":     "Hello World.",
        })
    })

    router.Run(":8080")
}
```

# CORS 処理

CORS（Cross-Origin Resource Sharing）は、異なるオリジン間でのリソース共有を許可するためのメカニズム
Gin では、CORS ミドルウェアを使用して、異なるオリジンからのリクエストを許可する設定ができる

```
import (
    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
    "time"
)

func main() {
    router := gin.Default()

    // CORS for all domains
    router.Use(cors.Default())
    // cors.Default()は、すべてのオリジンからのGET, POST, HEADのリクエストを許可。本番環境ではセキュリティリスクがあるため非推奨。

    // or more fine-grained setup
    router.Use(cors.New(cors.Config{
      AllowOrigins:     []string{"https://foo.com"},  // アクセスを許可するオリジン
      AllowMethods:     []string{"PUT", "PATCH"},    // 許可するHTTPメソッド
      AllowHeaders:     []string{"Origin"},          // 許可するHTTPヘッダ
      ExposeHeaders:    []string{"Content-Length"},  // JavaScriptからアクセスを許可するヘッダ
      AllowCredentials: true,                        // クレデンシャル付きのリクエストを許可するか
      AllowOriginFunc: func(origin string) bool {    // オリジンを動的に判断
          return origin == "https://github.com"
      },
      MaxAge: 12 * time.Hour,                        // プリフライトリクエストの結果をキャッシュする時間
   }))

    router.Run(":8080")
}
```

# 環境変数

https://github.com/joho/godotenv
go get github.com/joho/godotenv

# データベースの接続設定

- sqlite :テスト環境用 DB
- postgres :本番環境用 DB

```
go get -u gorm.io/gorm
```

```
go get -u gorm.io/driver/sqlite
```

```
go get -u gorm.io/driver/postgres
```

```
host := "localhost"
user := "user"
password := "password"
dbname := "mydatabase"
port := "5432"
```

# フィールドに指定可能なタグ

https://gorm.io/ja_JP/docs/models.html#%E3%83%95%E3%82%A3%E3%83%BC%E3%83%AB%E3%83%89%E3%81%AB%E6%8C%87%E5%AE%9A%E5%8F%AF%E8%83%BD%E3%81%AA%E3%82%BF%E3%82%B0

ロールバック機能がない

# Go の現在のバージョンを削除

sudo rm -rf /usr/local/go

# 新しいバージョンの Go をダウンロード

wget https://golang.org/dl/go1.20.linux-amd64.tar.gz

# アーカイブを解凍してインストール

sudo tar -C /usr/local -xzf go1.20.linux-amd64.tar.gz

# 環境変数の設定（必要に応じて）

export PATH=$PATH:/usr/local/go/bin

# JWT パッケージの利用（https://github.com/golang-jwt/jwt?tab=readme-ov-file#installation-guidelines）

```
go get -u github.com/golang-jwt/jwt/v5
```

JWT のシークレットキーの作成（32 バイトのランダムな 16 進数の文字列を生成）

```
openssl rand -hex 32
```

環境変数を以下のように設定
SECRET_KEY=ddc8510d20e53cf98797f8d1a938851a4966dd39f5fcc3409ea851d431db0011

デコードサイト
https://jwt.io/

# CORS の導入（https://github.com/gin-contrib/cors）

※ セキュリティー要件によって変更が必要

```
go get github.com/gin-contrib/cors
```

# 基礎知識の整理（ポインタの操作）

- 型の再定義を防ぐ（メモリ効率の向上）
- 値の格納されたアドレスであることを明確にする: \*int（データ共有）
- ポインタを通じてアドレスに格納された値を取得・変更

```
package main

import (
	"fmt"
)

func main() {
	var x int = 10
	var p *int = &x  // &xはxのアドレスを取得し、pに格納

	fmt.Println("xの値:", x)  // xの値を表示
	fmt.Println("xのアドレス:", &x)  // xのメモリアドレスを表示
	fmt.Println("pが指すアドレスの値:", *p)  // pをデリファレンスして、pが指すアドレスの値を表示

	*p = 20  // pをデリファレンスして、そのアドレスに新しい値(20)を格納
	fmt.Println("xの新しい値:", x)  // xの値に変更される
}
```
