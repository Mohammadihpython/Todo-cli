package deliveryparam

type Request struct {
	Command           string            `json:"command"`
	CreateTaskRequest CreateTaskRequest `json:"create_task_request"`
}

type CreateTaskRequest struct {
	Title      string `json:"title"`
	DuDate     string `json:"du_date"`
	CategoryId int    `json:"category_id"`
}
