package dto

// リクエストからEmail、Passwordを受け取るためdtoを作成
// リクエストのバリデーションを行い、必要なデータをサーバーに渡す役割
type SignupInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}
