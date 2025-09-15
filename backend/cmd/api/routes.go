package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.notFoundResponse)

	router.HandlerFunc(http.MethodGet, "/v1/task/:id", app.getTaskHandler)
	router.HandlerFunc(http.MethodPost, "/v1/task", app.CreateTaskHandler)
	router.HandlerFunc(http.MethodPatch, "/v1/task/:id", app.updateTaskHandler)
	router.HandlerFunc(http.MethodDelete, "/v1/task/:id", app.deleteTaskHandler)
	router.HandlerFunc(http.MethodGet, "/v1/tasks", app.listTasksHandler)

	return router
}
