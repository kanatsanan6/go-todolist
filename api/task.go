package api

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/kanatsanan6/go-todo-list/model"
	"github.com/kanatsanan6/go-todo-list/utils"
)

type TaskResponse struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func taskResponse(task model.Task) TaskResponse {
	return TaskResponse{
		ID:        task.ID,
		Title:     task.Title,
		Completed: task.Completed,
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
	}
}

func (server *Server) getTasks(c *fiber.Ctx) error {
	tasks, err := server.TaskRepo.FindAll()
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	taskResponses := []TaskResponse{}
	for _, task := range tasks {
		taskResponses = append(taskResponses, taskResponse(task))
	}
	return utils.JsonResponse(c, fiber.StatusOK, taskResponses)
}

type getTaskRequest struct {
	ID uint `json:"id"`
}

func (server *Server) getTask(c *fiber.Ctx) error {
	var req getTaskRequest
	if err := c.ParamsParser(&req); err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	task, err := server.TaskRepo.FindByID(req.ID)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusNotFound, err.Error())
	}
	return utils.JsonResponse(c, fiber.StatusOK, taskResponse(task))
}

type createTaskRequest struct {
	Title string `json:"title"`
}

func (server *Server) createTask(c *fiber.Ctx) error {
	var req createTaskRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	task, err := server.TaskRepo.Create(req.Title)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusUnprocessableEntity, err.Error())
	}

	return utils.JsonResponse(c, fiber.StatusOK, taskResponse(task))
}

type updateTaskParams struct {
	ID uint `json:"id"`
}

type updateTaskBody struct {
	Title     string `json:"title,omitempty"`
	Completed *bool  `json:"completed"`
}

func (server *Server) updateTask(c *fiber.Ctx) error {
	var params updateTaskParams
	if err := c.ParamsParser(&params); err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	var req updateTaskBody
	if err := c.BodyParser(&req); err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	task, err := server.TaskRepo.FindByID(params.ID)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusNotFound, err.Error())
	}

	b, err := utils.StructToMap(req)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}
	utils.RemoveNulls(b)

	task, err = server.TaskRepo.Update(&task, b)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusUnprocessableEntity, err.Error())
	}

	return utils.JsonResponse(c, fiber.StatusOK, taskResponse(task))
}

type deleteTaskParams struct {
	ID uint `json:"id"`
}

func (server *Server) deleteTask(c *fiber.Ctx) error {
	var params deleteTaskParams
	if err := c.ParamsParser(&params); err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	task, err := server.TaskRepo.FindByID(params.ID)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusNotFound, err.Error())
	}

	if err = server.TaskRepo.Delete(&task); err != nil {
		return utils.ErrorResponse(c, fiber.StatusUnprocessableEntity, err.Error())
	}

	return utils.JsonResponse(c, fiber.StatusNoContent, nil)
}
