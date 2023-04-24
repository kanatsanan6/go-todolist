package repository

import (
	"github.com/kanatsanan6/go-todo-list/model"
	"gorm.io/gorm"
)

type TaskRepo struct {
	db *gorm.DB
}

func NewTaskRepo(db *gorm.DB) *TaskRepo {
	return &TaskRepo{db}
}

func (r *TaskRepo) FindAll() ([]model.Task, error) {
	var tasks []model.Task
	if err := r.db.Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r *TaskRepo) FindByID(id uint) (model.Task, error) {
	task := model.Task{ID: id}
	if err := r.db.First(&task).Error; err != nil {
		return model.Task{}, err
	}
	return task, nil
}

func (r *TaskRepo) Create(title string) (model.Task, error) {
	task := model.Task{Title: title}
	if err := r.db.Create(&task).Error; err != nil {
		return model.Task{}, err
	}
	return task, nil
}

func (r *TaskRepo) Update(Task *model.Task, params map[string]interface{}) (model.Task, error) {
	if err := r.db.Model(Task).Updates(params).Error; err != nil {
		return model.Task{}, err
	}

	task, err := r.FindByID(Task.ID)
	if err != nil {
		return model.Task{}, err
	}
	return task, nil
}

func (r *TaskRepo) Delete(Task *model.Task) error {
	if err := r.db.Delete(&Task).Error; err != nil {
		return err
	}
	return nil
}
