package userHandlers

import (
	"context"
	"errors"
	"fmt"

	"task_user_RestAPI/internal/models"
	"task_user_RestAPI/internal/services/userService"
	"task_user_RestAPI/internal/web/users"

	"gorm.io/gorm"
)

type Handler struct {
	Service *userService.UserService
}

func NewHandler(service *userService.UserService) *Handler {
	return &Handler{Service: service}
}

// GetUsers получает все user
func (h *Handler) GetUsers(_ context.Context, _ users.GetUsersRequestObject) (users.GetUsersResponseObject, error) {
	allUsers, err := h.Service.GetAllUsers()
	if err != nil {
		return nil, fmt.Errorf("failed to get users: %w", err)
	}
	response := make(users.GetUsers200JSONResponse, len(allUsers))
	for i, usr := range allUsers {
		response[i] = users.User{
			Id:    &usr.ID,
			Email: &usr.Email,
			Pass:  &usr.Pass,
		}
	}
	return response, nil
}

// PostUsers создаёт нового user
func (h *Handler) PostUsers(_ context.Context, request users.PostUsersRequestObject) (users.PostUsersResponseObject, error) {
	userRequest := request.Body
	if userRequest.Email == nil || userRequest.Pass == nil {
		return nil, errors.New("invalid input: email and pass are required")
	}
	userToCreate := &models.User{
		Email: *userRequest.Email,
		Pass:  *userRequest.Pass,
	}
	err := h.Service.CreateUser(userToCreate)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}
	response := users.PostUsers201JSONResponse{
		Id:    &userToCreate.ID,
		Email: &userToCreate.Email,
		Pass:  &userToCreate.Pass,
	}
	return response, nil
}
func (h *Handler) PatchUsersId(_ context.Context, request users.PatchUsersIdRequestObject) (users.PatchUsersIdResponseObject, error) {
	userID := request.Id
	userRequest := request.Body

	existingUser, err := h.Service.GetUserByID(userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user with ID %d: %w", userID, err)
	}
	if userRequest.Email != nil {
		existingUser.Email = *userRequest.Email
	}
	if userRequest.Pass != nil {
		existingUser.Pass = *userRequest.Pass
	}

	err = h.Service.UpdateUser(&existingUser)
	if err != nil {
		return nil, fmt.Errorf("failed to update user with ID %d: %w", userID, err)
	}
	response := users.PatchUsersId200JSONResponse{
		Id:    &existingUser.ID,
		Email: &existingUser.Email,
		Pass:  &existingUser.Pass,
	}
	return response, nil
}

func (h *Handler) DeleteUsersId(_ context.Context, request users.DeleteUsersIdRequestObject) (users.DeleteUsersIdResponseObject, error) {
	userID := request.Id

	err := h.Service.DeleteUser(userID)
	if err != nil {
		return nil, fmt.Errorf("failed to delete user with ID %d: %w", userID, err)
	}
	return users.DeleteUsersId204Response{}, nil
}

func (h *Handler) GetUsersId(_ context.Context, request users.GetUsersIdRequestObject) (users.GetUsersIdResponseObject, error) {
	userID := request.Id
	user, err := h.Service.GetUserByID(userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return users.GetUsersId404Response{}, nil
		}
		return nil, fmt.Errorf("failed to get user with ID %d: %w", userID, err)
	}
	response := users.GetUsersId200JSONResponse{
		Id:    &userID,
		Email: &user.Email,
		Pass:  &user.Pass,
	}
	return response, nil
}
