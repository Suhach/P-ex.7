package repository

import (
	"task_user_RestAPI/internal/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetAllUsers() ([]models.User, error) {
	var users []models.User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *UserRepository) GetUserByID(id uint) (models.User, error) {
	var user models.User
	err := r.db.First(&user, id).Error
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (r *UserRepository) CreateUser(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepository) UpdateUserByID(user *models.User) error {
	return r.db.Save(user).Error
}

func (r *UserRepository) DeleteUserByID(id uint) error {
	return r.db.Delete(&models.User{}, id).Error
}

type TaskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

func (r *TaskRepository) GetAllTasks() ([]models.Task, error) {
	var tasks []models.Task
	err := r.db.Find(&tasks).Error
	return tasks, err
}

func (r *TaskRepository) GetTasksByUserID(userID uint) ([]models.Task, error) {
	var tasks []models.Task
	err := r.db.Where("user_id = ?", userID).Find(&tasks).Error
	return tasks, err
}

func (r *TaskRepository) GetTaskByID(id uint) (models.Task, error) {
	var task models.Task
	err := r.db.First(&task, id).Error
	if err != nil {
		return models.Task{}, err
	}
	return task, nil
}

func (r *TaskRepository) CreateTask(task *models.Task) error {
	return r.db.Create(task).Error
}

func (r *TaskRepository) UpdateTaskByID(task *models.Task) error {
	return r.db.Save(task).Error
}

func (r *TaskRepository) DeleteTaskByID(id uint) error {
	return r.db.Delete(&models.Task{}, id).Error
}
