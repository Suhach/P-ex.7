package taskService

type TaskService struct {
	repo *TaskRepository
}

func NewService(repo *TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) GetAllTasks() ([]Task, error) {
	return s.repo.GetAllTasks()
}

func (s *TaskService) GetTaskByID(id uint) (Task, error) {
	return s.repo.GetTaskByID(id)
}

func (s *TaskService) CreateTask(task *Task) error {
	return s.repo.CreateTask(task)
}

func (s *TaskService) UpdateTask(task *Task) error {
	return s.repo.UpdateTaskByID(task)
}

func (s *TaskService) DeleteTask(id uint) error {
	return s.repo.DeleteTaskByID(id)
}
