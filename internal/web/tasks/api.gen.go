// Package tasks provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.16.3 DO NOT EDIT.
package tasks

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	strictecho "github.com/oapi-codegen/runtime/strictmiddleware/echo"
)

// Task defines model for Task.
type Task struct {
	Id     *uint   `json:"id,omitempty"`
	IsDone *bool   `json:"is_done,omitempty"`
	Title  *string `json:"title,omitempty"`
}

// PostTasksJSONRequestBody defines body for PostTasks for application/json ContentType.
type PostTasksJSONRequestBody = Task

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Get all tasks
	// (GET /tasks)
	GetTasks(ctx echo.Context) error
	// Create a new task
	// (POST /tasks)
	PostTasks(ctx echo.Context) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetTasks converts echo context to params.
func (w *ServerInterfaceWrapper) GetTasks(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetTasks(ctx)
	return err
}

// PostTasks converts echo context to params.
func (w *ServerInterfaceWrapper) PostTasks(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PostTasks(ctx)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/tasks", wrapper.GetTasks)
	router.POST(baseURL+"/tasks", wrapper.PostTasks)

}

type GetTasksRequestObject struct {
}

type GetTasksResponseObject interface {
	VisitGetTasksResponse(w http.ResponseWriter) error
}

type GetTasks200JSONResponse []Task

func (response GetTasks200JSONResponse) VisitGetTasksResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type PostTasksRequestObject struct {
	Body *PostTasksJSONRequestBody
}

type PostTasksResponseObject interface {
	VisitPostTasksResponse(w http.ResponseWriter) error
}

type PostTasks201JSONResponse Task

func (response PostTasks201JSONResponse) VisitPostTasksResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)

	return json.NewEncoder(w).Encode(response)
}

// StrictServerInterface represents all server handlers.
type StrictServerInterface interface {
	// Get all tasks
	// (GET /tasks)
	GetTasks(ctx context.Context, request GetTasksRequestObject) (GetTasksResponseObject, error)
	// Create a new task
	// (POST /tasks)
	PostTasks(ctx context.Context, request PostTasksRequestObject) (PostTasksResponseObject, error)
}

type StrictHandlerFunc = strictecho.StrictEchoHandlerFunc
type StrictMiddlewareFunc = strictecho.StrictEchoMiddlewareFunc

func NewStrictHandler(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares}
}

type strictHandler struct {
	ssi         StrictServerInterface
	middlewares []StrictMiddlewareFunc
}

// GetTasks operation middleware
func (sh *strictHandler) GetTasks(ctx echo.Context) error {
	var request GetTasksRequestObject

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetTasks(ctx.Request().Context(), request.(GetTasksRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetTasks")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(GetTasksResponseObject); ok {
		return validResponse.VisitGetTasksResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// PostTasks operation middleware
func (sh *strictHandler) PostTasks(ctx echo.Context) error {
	var request PostTasksRequestObject

	var body PostTasksJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PostTasks(ctx.Request().Context(), request.(PostTasksRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PostTasks")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(PostTasksResponseObject); ok {
		return validResponse.VisitPostTasksResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}
