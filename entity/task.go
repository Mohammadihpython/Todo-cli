package entity

type Task struct {
	ID         int
	Title      string
	DueDate    string
	CategoryID int
	ISDone     bool
	UserID     int
}
