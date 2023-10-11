package handler

import (
	"gin-user-tasks/src/app/tasks/schema"
	"gin-user-tasks/src/app/tasks/usecase"
	"gin-user-tasks/src/pkg/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	taskHandler struct {
		taskUsecase usecase.TaskUsecase
	}
)

func NewTaskHandler(usecase usecase.TaskUsecase) *taskHandler {
	return &taskHandler{taskUsecase: usecase}
}

func (h *taskHandler) Mount(g *gin.RouterGroup) {
	g.POST("", h.InsertTask)
	g.GET("", h.GetAllTasks)
	g.GET("/:id", h.GetTaskByID)
	g.PUT("/:id", h.UpdateTaskByID)
	g.DELETE("/:id", h.DeleteTaskByID)
}

func (h *taskHandler) InsertTask(c *gin.Context) {
	body := schema.InsertTaskRequest{}
	if err := c.ShouldBindJSON(&body); err != nil {
		resp := util.NewResponse(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), nil)
		resp.ErrorMessage = err.Error()
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	body.UserID = 1
	result, err := h.taskUsecase.InsertTask(c.Request.Context(), &body)
	if err != nil {
		resp := util.NewResponse(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), nil)
		resp.ErrorMessage = err.Error()
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	resp := util.NewResponse(http.StatusOK, `success to add a task`, result)
	c.JSON(http.StatusOK, resp)
}

func (h *taskHandler) GetAllTasks(c *gin.Context) {
	result, err := h.taskUsecase.GetAllTasks(c.Request.Context())
	if err != nil {
		resp := util.NewResponse(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), nil)
		resp.ErrorMessage = err.Error()
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	resp := util.NewResponse(http.StatusOK, ``, result)
	c.JSON(http.StatusOK, resp)
}

func (h *taskHandler) GetTaskByID(c *gin.Context) {
	uriParam := schema.GetTaskByIDRequest{}

	if err := c.ShouldBindUri(&uriParam); err != nil {
		resp := util.NewResponse(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), nil)
		resp.ErrorMessage = err.Error()
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	if uriParam.ID == 0 {
		resp := util.NewResponse(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), nil)
		resp.ErrorMessage = `invalid task id`
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	result, err := h.taskUsecase.GetTaskByID(c.Request.Context(), uriParam.ID)
	if err != nil {
		resp := util.NewResponse(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), nil)
		resp.ErrorMessage = err.Error()
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	resp := util.NewResponse(http.StatusOK, ``, result)
	c.JSON(http.StatusOK, resp)

}

func (h *taskHandler) UpdateTaskByID(c *gin.Context) {
	uriParam := schema.EditTaskURI{}

	if err := c.ShouldBindUri(&uriParam); err != nil {
		resp := util.NewResponse(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), nil)
		resp.ErrorMessage = err.Error()
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	if uriParam.ID == 0 {
		resp := util.NewResponse(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), nil)
		resp.ErrorMessage = `invalid user id`
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	body := schema.EditTaskRequest{}
	if err := c.ShouldBindJSON(&body); err != nil {
		resp := util.NewResponse(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), nil)
		resp.ErrorMessage = err.Error()
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	result, err := h.taskUsecase.EditTaskByID(c.Request.Context(), uriParam.ID, &body)
	if err != nil {
		resp := util.NewResponse(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), nil)
		resp.ErrorMessage = err.Error()
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	resp := util.NewResponse(http.StatusOK, `success to update task`, result)
	c.JSON(http.StatusOK, resp)
}

func (h *taskHandler) DeleteTaskByID(c *gin.Context) {
	uriParam := schema.DeleteUserByIDRequest{}

	if err := c.ShouldBindUri(&uriParam); err != nil {
		resp := util.NewResponse(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), nil)
		resp.ErrorMessage = err.Error()
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	if uriParam.ID == 0 {
		resp := util.NewResponse(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), nil)
		resp.ErrorMessage = `invalid user id`
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	err := h.taskUsecase.DeleteTaskByID(c.Request.Context(), uriParam.ID)
	if err != nil {
		resp := util.NewResponse(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), nil)
		resp.ErrorMessage = err.Error()
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	resp := util.NewResponse(http.StatusOK, `success to delete task`, nil)
	c.JSON(http.StatusOK, resp)
}
