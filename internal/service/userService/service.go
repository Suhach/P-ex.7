package userService

type UserService struct {
	repo *UserRepository
}

func NewService(repo *UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetAllUsers() ([]User, error) {
	return s.repo.GetAllUsers()
}

func (s *UserService) GetUserByID(id uint) (User, error) {
	return s.repo.GetUserByID(id)
}

func (s *UserService) CreateUser(user *User) error {
	return s.repo.CreateUser(user)
}

func (s *UserService) UpdateUser(user *User) error {
	return s.repo.UpdateUserByID(user)
}

func (s *UserService) DeleteUser(id uint) error {
	return s.repo.DeleteUserByID(id)
}
