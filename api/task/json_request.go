package task

type AddReq struct {
	UserId      string `json:"userId" binding:"required" extensions:"x-order=1"`
	Title       string `json:"title" binding:"required" extensions:"x-order=2"`
	Description string `json:"description" binding:"required" extensions:"x-order=3"`
}

type UpdateReq struct {
	Title       string `json:"title" binding:"required" extensions:"x-order=1"`
	Description string `json:"description" binding:"required" extensions:"x-order=2"`
}
