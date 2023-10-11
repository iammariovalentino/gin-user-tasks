package handler

import (
	"context"
	"gin-user-tasks/src/pkg/config"
	"gin-user-tasks/src/pkg/util"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2/clientcredentials"
)

type (
	oauthClientHandler struct {
		oauthConfig *clientcredentials.Config
	}
)

func NewOauthClientHandler(conf *config.OauthConfig) *oauthClientHandler {

	oa2 := &clientcredentials.Config{
		ClientID:     conf.ClientID,
		ClientSecret: conf.ClientSecret,
		TokenURL:     conf.TokenURL,
		Scopes:       []string{"tasks"},
	}

	return &oauthClientHandler{oauthConfig: oa2}
}

func (h *oauthClientHandler) Mount(g *gin.RouterGroup) {
	g.POST("/access-token", h.RequestToken)
}

func (h *oauthClientHandler) RequestToken(c *gin.Context) {
	token, err := h.oauthConfig.Token(context.Background())
	if err != nil {
		resp := util.NewResponse(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), nil)
		resp.ErrorMessage = err.Error()
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	resp := map[string]interface{}{}
	resp["token"] = token

	c.JSON(http.StatusOK, resp)
}
