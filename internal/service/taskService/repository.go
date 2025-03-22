package taskService

import "gorm.io/gorm"

type TaskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

func (r *TaskRepository) GetAllTasks() ([]Task, error) {
	var tasks []Task
	err := r.db.Find(&tasks).Error
	return tasks, err
}

func (r *TaskRepository) GetTaskByID(id uint) (Task, error) {
	var task Task
	err := r.db.First(&task, id).Error
	if err != nil {
		return Task{}, err
	}
	return task, nil
}

func (r *TaskRepository) CreateTask(task *Task) error {
	return r.db.Create(task).Error
}

func (r *TaskRepository) UpdateTaskByID(task *Task) error {
	return r.db.Save(task).Error
}

func (r *TaskRepository) DeleteTaskByID(id uint) error {
	return r.db.Delete(&Task{}, id).Error
}
