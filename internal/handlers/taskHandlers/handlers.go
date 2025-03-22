package taskHandlers

import (
	"context"
	"errors"
	"fmt"

	"P_ex.7/internal/service/taskService"
	"P_ex.7/internal/web/tasks"
	"gorm.io/gorm"
)

type Handler struct {
	Service *taskService.TaskService
}

func NewHandler(service *taskService.TaskService) *Handler {
	return &Handler{Service: service}
}

// GetTasks получает все задачи.
func (h *Handler) GetTasks(_ context.Context, _ tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error) {
	allTasks, err := h.Service.GetAllTasks()
	if err != nil {
		return nil, fmt.Errorf("failed to get tasks: %w", err)
	}

	response := make(tasks.GetTasks200JSONResponse, len(allTasks))
	for i, tsk := range allTasks {
		response[i] = tasks.Task{
			Id:     &tsk.ID,
			Task:   &tsk.Task,
			IsDone: &tsk.IsDone,
		}
	}
	return response, nil
}

// PostTasks создает новую задачу.
func (h *Handler) PostTasks(_ context.Context, request tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {
	taskRequest := request.Body
	if taskRequest.Task == nil || taskRequest.IsDone == nil {
		return nil, errors.New("invalid input: task and is_done are required")
	}

	// Создаем задачу с использованием указателя
	taskToCreate := &taskService.Task{
		Task:   *taskRequest.Task,
		IsDone: *taskRequest.IsDone,
	}

	// Передаем указатель на taskToCreate в CreateTask
	err := h.Service.CreateTask(taskToCreate)
	if err != nil {
		return nil, fmt.Errorf("failed to create task: %w", err)
	}

	response := tasks.PostTasks201JSONResponse{
		Id:     &taskToCreate.ID,
		Task:   &taskToCreate.Task,
		IsDone: &taskToCreate.IsDone,
	}
	return response, nil
}

// PatchTasksId обновляет задачу по ID.
func (h *Handler) PatchTasksId(_ context.Context, request tasks.PatchTasksIdRequestObject) (tasks.PatchTasksIdResponseObject, error) {
	taskID := request.Id
	taskRequest := request.Body

	// Получаем текущую задачу из базы данных
	existingTask, err := h.Service.GetTaskByID(taskID)
	if err != nil {
		return nil, fmt.Errorf("failed to get task with ID %d: %w", taskID, err)
	}

	// Обновляем поля задачи, если они переданы
	if taskRequest.Task != nil {
		existingTask.Task = *taskRequest.Task
	}
	if taskRequest.IsDone != nil {
		existingTask.IsDone = *taskRequest.IsDone
	}

	// Сохраняем обновленную задачу
	err = h.Service.UpdateTask(&existingTask)
	if err != nil {
		return nil, fmt.Errorf("failed to update task with ID %d: %w", taskID, err)
	}

	response := tasks.PatchTasksId200JSONResponse{
		Id:     &existingTask.ID,
		Task:   &existingTask.Task,
		IsDone: &existingTask.IsDone,
	}
	return response, nil
}

// DeleteTasksId удаляет задачу по ID.
func (h *Handler) DeleteTasksId(_ context.Context, request tasks.DeleteTasksIdRequestObject) (tasks.DeleteTasksIdResponseObject, error) {
	taskID := request.Id

	// Удаляем задачу
	err := h.Service.DeleteTask(taskID)
	if err != nil {
		return nil, fmt.Errorf("failed to delete task with ID %d: %w", taskID, err)
	}

	return tasks.DeleteTasksId204Response{}, nil
}

func (h *Handler) GetTasksId(_ context.Context, request tasks.GetTasksIdRequestObject) (tasks.GetTasksIdResponseObject, error) {
	taskID := request.Id
	task, err := h.Service.GetTaskByID(taskID)
	if err != nil {
		// Если задача не найдена, возвращаем ошибку 404
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return tasks.GetTasksId404Response{}, nil
		}
		return nil, fmt.Errorf("failed to get task with ID %d: %w", taskID, err)
	}
	response := tasks.GetTasksId200JSONResponse{
		Id:     &task.ID,
		Task:   &task.Task,
		IsDone: &task.IsDone,
	}
	return response, nil
}
