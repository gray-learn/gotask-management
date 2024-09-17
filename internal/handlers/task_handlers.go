package handlers

import (
	// "encoding/json"
	"html/template"
	"net/http"
	"strconv"

	"gotask-management/internal/database"
	"gotask-management/internal/models"

	"github.com/gorilla/mux"
)

func ListTasksHandler(w http.ResponseWriter, r *http.Request) {
	tasks, err := database.GetAllTasks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl := template.Must(template.ParseFiles("web/templates/layout.html", "web/templates/tasks/list.html"))
	tmpl.Execute(w, tasks)
}

func CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		task := &models.Task{
			Title:       r.FormValue("title"),
			Description: r.FormValue("description"),
			Status:      r.FormValue("status"),
		}
		id, err := database.CreateTask(task)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		task.ID = id

		tmpl := template.Must(template.ParseFiles("web/templates/tasks/list.html"))
		tmpl.ExecuteTemplate(w, "task-row", task)
	} else {
		tmpl := template.Must(template.ParseFiles("web/templates/tasks/create.html"))
		tmpl.Execute(w, nil)
	}
}

func EditTaskHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.ParseInt(vars["id"], 10, 64)

	if r.Method == "PUT" {
		task := &models.Task{
			ID:          id,
			Title:       r.FormValue("title"),
			Description: r.FormValue("description"),
			Status:      r.FormValue("status"),
		}
		err := database.UpdateTask(task)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tmpl := template.Must(template.ParseFiles("web/templates/tasks/list.html"))
		tmpl.ExecuteTemplate(w, "task-row", task)
	} else {
		task, err := database.GetTaskByID(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tmpl := template.Must(template.ParseFiles("web/templates/tasks/edit.html"))
		tmpl.Execute(w, task)
	}
}

func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.ParseInt(vars["id"], 10, 64)

	err := database.DeleteTask(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
