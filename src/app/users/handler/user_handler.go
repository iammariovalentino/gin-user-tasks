package handler

import (
	"gin-user-tasks/src/app/users/schema"
	"gin-user-tasks/src/app/users/usecase"
	"gin-user-tasks/src/pkg/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	userHandler struct {
		userUsecase usecase.UserUsecase
	}
)

func NewUserHandler(usecase usecase.UserUsecase) *userHandler {
	return &userHandler{userUsecase: usecase}
}

func (h *userHandler) Mount(g *gin.RouterGroup) {
	g.POST("", h.RegisterUser)
	g.GET("", h.GetAllUsers)
	g.GET("/:id", h.GetUserByID)
	g.PUT("/:id", h.UpdateUserByID)
	g.DELETE("/:id", h.DeleteUserByID)
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	body := schema.RegisterUserRequest{}
	if err := c.ShouldBindJSON(&body); err != nil {
		resp := util.NewResponse(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), nil)
		resp.ErrorMessage = err.Error()
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	result, err := h.userUsecase.RegisterUser(c.Request.Context(), &body)
	if err != nil {
		resp := util.NewResponse(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), nil)
		resp.ErrorMessage = err.Error()
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	resp := util.NewResponse(http.StatusOK, `success to register`, result)
	c.JSON(http.StatusOK, resp)
}

func (h *userHandler) GetAllUsers(c *gin.Context) {
	result, err := h.userUsecase.GetAllUsers(c.Request.Context())
	if err != nil {
		resp := util.NewResponse(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), nil)
		resp.ErrorMessage = err.Error()
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	resp := util.NewResponse(http.StatusOK, ``, result)
	c.JSON(http.StatusOK, resp)
}

func (h *userHandler) GetUserByID(c *gin.Context) {
	uriParam := schema.GetUserByIDRequest{}

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

	result, err := h.userUsecase.GetUserByID(c.Request.Context(), uriParam.ID)
	if err != nil {
		resp := util.NewResponse(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), nil)
		resp.ErrorMessage = err.Error()
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	resp := util.NewResponse(http.StatusOK, ``, result)
	c.JSON(http.StatusOK, resp)
}

func (h *userHandler) UpdateUserByID(c *gin.Context) {
	uriParam := schema.EditUserURI{}

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

	body := schema.EditUserRequest{}
	if err := c.ShouldBindJSON(&body); err != nil {
		resp := util.NewResponse(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), nil)
		resp.ErrorMessage = err.Error()
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	result, err := h.userUsecase.EditUserByID(c.Request.Context(), uriParam.ID, &body)
	if err != nil {
		resp := util.NewResponse(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), nil)
		resp.ErrorMessage = err.Error()
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	resp := util.NewResponse(http.StatusOK, `success to update user`, result)
	c.JSON(http.StatusOK, resp)
}

func (h *userHandler) DeleteUserByID(c *gin.Context) {
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

	err := h.userUsecase.DeleteUserByID(c.Request.Context(), uriParam.ID)
	if err != nil {
		resp := util.NewResponse(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), nil)
		resp.ErrorMessage = err.Error()
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	resp := util.NewResponse(http.StatusOK, `success to delete user`, nil)
	c.JSON(http.StatusOK, resp)
}
