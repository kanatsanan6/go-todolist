package utils_test

import (
	"testing"

	"github.com/kanatsanan6/go-todo-list/utils"
	"github.com/stretchr/testify/assert"
)

func TestRemoveNulls(t *testing.T) {
	payload := map[string]interface{}{
		"data_1": "data_1",
		"data_2": map[string]interface{}{
			"data_2_1": "data_2_1",
			"data_2_2": nil,
		},
		"data_3": nil,
	}
	expected := map[string]interface{}{
		"data_1": "data_1",
		"data_2": map[string]interface{}{
			"data_2_1": "data_2_1",
		},
	}

	utils.RemoveNulls(payload)
	assert.Equal(t, expected, payload)
}

type Task struct {
	Title     string
	Completed bool
}

func TestStructToMap(t *testing.T) {
	t.Run("when there is no error", func(t *testing.T) {
		task := Task{
			Title:     "title",
			Completed: true,
		}
		expected := map[string]interface{}{
			"Title":     "title",
			"Completed": true,
		}

		result, err := utils.StructToMap(task)
		assert.NoError(t, err)
		assert.Equal(t, expected, result)
	})
	t.Run("when error happens", func(t *testing.T) {
		task := map[string]interface{}{
			"foo": make(chan int),
		}

		_, err := utils.StructToMap(task)
		assert.Error(t, err)
	})
}
