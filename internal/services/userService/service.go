package userService

import (
	"task_user_RestAPI/internal/models"
	"task_user_RestAPI/internal/repository"
	"task_user_RestAPI/internal/services/taskService"
)

type UserService struct {
	repo        *repository.UserRepository
	taskService *taskService.TaskService
}

func NewService(repo *repository.UserRepository, taskService *taskService.TaskService) *UserService {
	return &UserService{
		repo:        repo,
		taskService: taskService,
	}
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
	return s.repo.GetAllUsers()
}

func (s *UserService) GetUserByID(id uint) (models.User, error) {
	return s.repo.GetUserByID(id)
}

func (s *UserService) CreateUser(user *models.User) error {
	return s.repo.CreateUser(user)
}

func (s *UserService) UpdateUser(user *models.User) error {
	return s.repo.UpdateUserByID(user)
}

func (s *UserService) DeleteUser(id uint) error {
	return s.repo.DeleteUserByID(id)
}

// GetTasksForUser возвращает все задачи, принадлежащие пользователю
func (s *UserService) GetTasksForUser(userID uint) ([]models.Task, error) {
	// Проверяем существование пользователя
	_, err := s.GetUserByID(userID)
	if err != nil {
		return nil, err
	}

	// Получаем задачи пользователя через taskService
	return s.taskService.GetTasksByUserID(userID)
}
