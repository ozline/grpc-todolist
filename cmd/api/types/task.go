package types

type Task struct {
	ID        int64  `json:"id"`
	UserID    int64  `json:"user_id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Status    int64  `json:"status"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

type TaskCreateRequest struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

type TaskUpdateRequest struct {
	ID      int64  `json:"id" binding:"required"`
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

type TaskDeleteRequest struct {
	ID int64 `json:"id" binding:"required"`
}

type TaskListRequest struct {
	Status *int64 `json:"status" binding:"required" form:"status"`
}
