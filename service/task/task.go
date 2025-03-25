package task

import (
	"fmt"
	"todocli/entity"
)

type ServiceRepository interface {
	CreateNewTask(t entity.Task) (entity.Task, error)
	ListUserTasks(userID int) ([]entity.Task, error)
}
type CategoryService interface {
	DoseThisUserHaveThisCategoryID(UserID, CategoryID int) bool
}
type Service struct {
	repository      ServiceRepository
	categoryService CategoryService
}

func NewService(repo ServiceRepository) Service {
	return Service{
		repository: repo,
	}
}

type CreateRequest struct {
	Title               string
	DuDate              string
	CategoryID          int
	AuthenticatedUserID int
}

type CreateResponse struct {
	Task entity.Task
}

func (t Service) Create(req CreateRequest) (CreateResponse, error) {
	createdTask, ok := t.repository.CreateNewTask(entity.Task{
		Title:      req.Title,
		DueDate:    req.DuDate,
		CategoryID: req.CategoryID,
		ISDone:     false,
		UserID:     req.AuthenticatedUserID,
	})
	if ok != nil {
		return CreateResponse{}, fmt.Errorf("User dose not have this category")
	}

	createdTask, cErr := t.repository.CreateNewTask(createdTask)
	if cErr != nil {
		return CreateResponse{}, fmt.Errorf("cant create new task:%w", cErr)
	}

	return CreateResponse{
		createdTask,
	}, nil

}

type ListRequest struct {
	UserID int
}

type ListResponse struct {
	Tasks []entity.Task
}

func (t Service) List(req ListRequest) (ListResponse, error) {
	tasks, err := t.repository.ListUserTasks(req.UserID)
	if err != nil {
		return ListResponse{}, fmt.Errorf("can't list user tasks: %v", err)
	}

	return ListResponse{Tasks: tasks}, nil
}
