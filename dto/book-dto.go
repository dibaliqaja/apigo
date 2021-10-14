package dto

type BookUpdateDTO struct {
	Id uint64 `json:"id" form:"id" binding:"required"`
	Title string `json:"title" form:"title" binding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
	UserId uint64 `json:"user_id,omitempty" form:"user_id,omitempty"`
}

type BookCreateDTO struct {
	Title string `json:"title" form:"title" binding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
	UserId uint64 `json:"user_id,omitempty" form:"user_id,omitempty"`
}