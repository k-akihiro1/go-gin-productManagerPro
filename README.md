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
