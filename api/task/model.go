package task

import (
	"github.com/guregu/null/v6"
)

type Task struct {
	Id          string    `json:"id" binding:"required" extensions:"x-order=1"`
	UserId      string    `json:"userId" binding:"required" extensions:"x-order=2"`
	Title       string    `json:"title" binding:"required" extensions:"x-order=3"`
	Description string    `json:"description" binding:"required" extensions:"x-order=4"`
	CreatedAt   null.Time `json:"createdAt" binding:"required" extensions:"x-order=5"`
	UpdatedAt   null.Time `json:"updatedAt" binding:"required" extensions:"x-order=6"`
}

type UserTask struct {
	Id          string    `json:"id" binding:"required" extensions:"x-order=1"`
	Title       string    `json:"title" binding:"required" extensions:"x-order=2"`
	Description string    `json:"description" binding:"required" extensions:"x-order=3"`
	CreatedAt   null.Time `json:"createdAt" binding:"required" extensions:"x-order=4"`
	UpdatedAt   null.Time `json:"updatedAt" binding:"required" extensions:"x-order=5"`
}
