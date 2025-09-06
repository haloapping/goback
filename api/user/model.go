package user

import "github.com/guregu/null/v6"

type UserRegister struct {
	Id        string    `json:"id" example:"DSFSGGRRG" binding:"required" extensions:"x-order=1"`
	Username  string    `json:"username" binding:"required" extensions:"x-order=2"`
	CreatedAt null.Time `json:"createdAt" binding:"required" extensions:"x-order=3"`
	UpdatedAt null.Time `json:"updatedAt" binding:"required" extensions:"x-order=4"`
}

type UserLogin struct {
	Token string `json:"token" extensions:"x-order=1"`
}

type UserBiodata struct {
	Id        string    `json:"id" binding:"required" extensions:"x-order=1"`
	Username  string    `json:"username" binding:"required" extensions:"x-order=2"`
	CreatedAt null.Time `json:"createdAt" binding:"required" extensions:"x-order=3"`
	UpdatedAt null.Time `json:"updatedAt" binding:"required" extensions:"x-order=4"`
}
