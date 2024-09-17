package database

import (
	"database/sql"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"gotask-management/internal/models"
)

var db *sql.DB

func InitDB(dataSourceName string) (*sql.DB, error) {
	var err error
	db, err = sql.Open("sqlite3", dataSourceName)
	if err != nil {
		return nil, err
	}

	// Create tasks table if it doesn't exist
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS tasks (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			title TEXT NOT NULL,
			description TEXT,
			status TEXT NOT NULL,
			assigned_to INTEGER,
			created_at DATETIME,
			updated_at DATETIME,
			deadline DATETIME
		)
	`)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func GetAllTasks() ([]*models.Task, error) {
	rows, err := db.Query("SELECT id, title, description, status, assigned_to, created_at, updated_at, deadline FROM tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []*models.Task
	for rows.Next() {
		task := &models.Task{}
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Status, &task.AssignedTo, &task.CreatedAt, &task.UpdatedAt, &task.Deadline)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func CreateTask(task *models.Task) (int64, error) {
	result, err := db.Exec("INSERT INTO tasks (title, description, status, created_at, updated_at) VALUES (?, ?, ?, ?, ?)",
		task.Title, task.Description, task.Status, time.Now(), time.Now())
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func GetTaskByID(id int64) (*models.Task, error) {
	task := &models.Task{}
	err := db.QueryRow("SELECT id, title, description, status, assigned_to, created_at, updated_at, deadline FROM tasks WHERE id = ?", id).
		Scan(&task.ID, &task.Title, &task.Description, &task.Status, &task.AssignedTo, &task.CreatedAt, &task.UpdatedAt, &task.Deadline)
	if err != nil {
		return nil, err
	}
	return task, nil
}

func UpdateTask(task *models.Task) error {
	_, err := db.Exec("UPDATE tasks SET title = ?, description = ?, status = ?, updated_at = ? WHERE id = ?",
		task.Title, task.Description, task.Status, time.Now(), task.ID)
	return err
}

func DeleteTask(id int64) error {
	_, err := db.Exec("DELETE FROM tasks WHERE id = ?", id)
	return err
}