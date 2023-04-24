package utils_test

import (
	"encoding/json"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/kanatsanan6/go-todo-list/model"
	"github.com/kanatsanan6/go-todo-list/utils"
	"github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp"
)

func TestErrorResponse(t *testing.T) {
	app := fiber.New()
	ctx := app.AcquireCtx(&fasthttp.RequestCtx{})

	err := utils.ErrorResponse(ctx, fiber.StatusUnprocessableEntity, "cannot update record")
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusUnprocessableEntity, ctx.Response().StatusCode())

	var data struct {
		Errors string `json:"errors"`
	}
	err = json.Unmarshal(ctx.Response().Body(), &data)
	assert.NoError(t, err)
	assert.Equal(t, "cannot update record", data.Errors)

}

func TestJsonResponse(t *testing.T) {
	app := fiber.New()
	ctx := app.AcquireCtx(&fasthttp.RequestCtx{})

	task := model.Task{Title: "title", Completed: true}
	err := utils.JsonResponse(ctx, fiber.StatusCreated, task)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusCreated, ctx.Response().StatusCode())

	var data struct {
		Data model.Task `json:"data"`
	}
	err = json.Unmarshal(ctx.Response().Body(), &data)
	assert.NoError(t, err)
	assert.Equal(t, task, data.Data)
}
