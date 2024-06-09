package dto

type CreateProductInput struct {
	Name        string `json:"name" binding:"required,min=2"`
	Price       uint   `json:"price" binding:"required,min=1,max=999999"`
	Description string `json:"description"`
}

type UpdateProductInput struct {
	Name        *string `json:"name" binding:"omitnil,min=2"`
	Price       *uint   `json:"price" binding:"omitnil,min=1,max=999999"`
	Description *string `json:"description"`
	SoldOut     *bool   `json:"soldOut"`
}

// omitnilが設定されているフィールド
// nilの場合、そのフィールドは更新対象から除外される
