package models

type ReqBody struct {
	Body  string `json:"body" binding:"required"`
	Title string `json:"title" binding:"required"`
}