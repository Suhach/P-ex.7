package userService

import "gorm.io/gorm"

type UserRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetAllUsers() ([]User, error) {
	var users []User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *UserRepository) GetUserByID(id uint) (User, error) {
	var user User
	err := r.db.First(&user, id).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func (r *UserRepository) CreateUser(user *User) error {
	return r.db.Create(user).Error
}

func (r *UserRepository) UpdateUserByID(user *User) error {
	return r.db.Save(user).Error
}

func (r *UserRepository) DeleteUserByID(id uint) error {
	return r.db.Delete(&User{}, id).Error
}
