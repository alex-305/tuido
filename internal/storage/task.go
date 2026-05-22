package storage

import (
	"github.com/alex-305/tuido/internal/types"
)

func GetTasksWithProjectID(projectID string) ([]types.Task, error) {
	var tasks []types.Task
	err := db.Select(&tasks,
		getTasks+"WHERE project_id = ?",
		projectID,
	)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

const getTasks = "SELECT * FROM Tasks "
