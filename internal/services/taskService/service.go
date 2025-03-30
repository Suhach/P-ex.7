package taskService

import (
	"fmt"
	"task_user_RestAPI/internal/models"
	"task_user_RestAPI/internal/repository"
)

type TaskService struct {
	repo     *repository.TaskRepository
	userRepo *repository.UserRepository
}

func NewService(repo *repository.TaskRepository, userRepo *repository.UserRepository) *TaskService {
	return &TaskService{
		repo:     repo,
		userRepo: userRepo,
	}
}

func (s *TaskService) GetAllTasks() ([]models.Task, error) {
	return s.repo.GetAllTasks()
}

func (s *TaskService) GetTasksByUserID(userID uint) ([]models.Task, error) {
	// Проверяем существование пользователя
	_, err := s.userRepo.GetUserByID(userID)
	if err != nil {
		return nil, fmt.Errorf("user with ID %d not found: %w", userID, err)
	}
	return s.repo.GetTasksByUserID(userID)
}

func (s *TaskService) GetTaskByID(id uint) (models.Task, error) {
	return s.repo.GetTaskByID(id)
}

func (s *TaskService) CreateTask(task *models.Task) error {
	// Проверяем существование пользователя перед созданием задачи
	_, err := s.userRepo.GetUserByID(task.UserID)
	if err != nil {
		return fmt.Errorf("cannot create task: user with ID %d not found: %w", task.UserID, err)
	}
	return s.repo.CreateTask(task)
}

func (s *TaskService) UpdateTask(task *models.Task) error {
	return s.repo.UpdateTaskByID(task)
}

func (s *TaskService) DeleteTask(id uint) error {
	return s.repo.DeleteTaskByID(id)
}
