package api

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/kanatsanan6/go-todo-list/repository"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Server struct {
	DB       *gorm.DB
	Router   *fiber.App
	TaskRepo *repository.TaskRepo
}

func NewServer() (*Server, error) {
	server := &Server{}
	if err := server.connectDatabase(); err != nil {
		return nil, err
	}

	server.setupRepo()
	server.setupRouter()

	return server, nil
}

func (server *Server) connectDatabase() error {
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to open database connection: %w", err)
	}

	server.DB = db

	return nil
}

func (server *Server) setupRepo() {
	server.TaskRepo = repository.NewTaskRepo(server.DB)
}

func (server *Server) setupRouter() {
	app := fiber.New()

	app.Get("/tasks", server.getTasks)
	app.Get("/tasks/:id", server.getTask)
	app.Post("/tasks", server.createTask)
	app.Put("/tasks/:id", server.updateTask)
	app.Delete("/tasks/:id", server.deleteTask)

	server.Router = app
}

func (server *Server) Start(address string) error {
	port := fmt.Sprintf(":%s", address)

	return server.Router.Listen(port)
}
